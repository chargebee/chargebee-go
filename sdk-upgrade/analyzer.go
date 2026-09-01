package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `chargebeerequest: rewrite string-id calls to use request objects

Operations such as Retrieve and Delete used to take the resource id as a bare
string:

	res, err := client.Subscription.Retrieve("sub_1")

They now take a request object, so that a context, custom headers, an
idempotency key and other per-request settings can be attached:

	res, err := client.Subscription.Retrieve(&chargebee.SubscriptionRetrieveRequest{Id: "sub_1"})

This analyzer finds the old call form and offers the new one as a fix. Upgrade
the SDK first, then run the analyzer with -fix over the packages to upgrade. It
reads the new signatures from the upgraded SDK rather than from a baked-in list
of operations, so it keeps working as operations are added.`

// sdkPackagePrefix identifies the packages whose service types are rewritten.
const sdkPackagePrefix = "github.com/chargebee/chargebee-go/v4"

// Analyzer rewrites string-id calls into request-object calls.
var Analyzer = &analysis.Analyzer{
	Name:     "chargebeerequest",
	Doc:      doc,
	URL:      "https://github.com/chargebee/chargebee-go/tree/v4/sdk-upgrade",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	// The whole point of the analyzer is to run over code that no longer
	// compiles against the upgraded SDK.
	RunDespiteErrors: true,
	// Declaring a fact type makes the driver type-check every package from
	// source rather than from export data. Export data is unavailable for a
	// build that fails, which is exactly the state this analyzer runs in.
	FactTypes: []analysis.Fact{(*sourceMode)(nil)},
	Run:       run,
}

// sourceMode is never reported; it only exists to select source-mode loading.
// See Analyzer.FactTypes.
type sourceMode struct{}

func (*sourceMode) AFact() {}

func run(pass *analysis.Pass) (any, error) {
	byFilename := make(map[string]*ast.File, len(pass.Files))
	for _, file := range pass.Files {
		if tokenFile := pass.Fset.File(file.FileStart); tokenFile != nil {
			byFilename[tokenFile.Name()] = file
		}
	}
	// An import is added at most once per file, so that applying every fix in
	// a file cannot produce duplicate import lines.
	importAdded := make(map[*ast.File]bool)

	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(node ast.Node) {
		call := node.(*ast.CallExpr)
		if len(call.Args) != 1 || call.Ellipsis.IsValid() {
			return
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		fn, ok := pass.TypesInfo.Uses[sel.Sel].(*types.Func)
		if !ok {
			return
		}
		service, ok := sdkServiceReceiver(fn)
		if !ok {
			return
		}
		request, idField, ok := requestParam(fn, service.Pkg())
		if !ok {
			return
		}

		arg := call.Args[0]
		argType := pass.TypesInfo.TypeOf(arg)
		if argType == nil {
			return
		}
		if types.AssignableTo(argType, types.NewPointer(request)) {
			return // already upgraded
		}
		if !types.AssignableTo(argType, idField.Type()) {
			return // not a plain string id; leave it for the compiler to report
		}

		file := byFilename[pass.Fset.File(call.Pos()).Name()]
		if file == nil {
			return
		}
		qualifier, edits, ok := importEdits(file, request.Obj().Pkg(), importAdded)
		if !ok {
			pass.Reportf(call.Pos(), "%s takes a *%s; import %q to build one",
				fn.Name(), request.Obj().Name(), request.Obj().Pkg().Path())
			return
		}

		var literal bytes.Buffer
		literal.WriteString("&")
		if qualifier != "" {
			literal.WriteString(qualifier)
			literal.WriteString(".")
		}
		literal.WriteString(request.Obj().Name())
		literal.WriteString("{Id: ")
		if err := printer.Fprint(&literal, pass.Fset, arg); err != nil {
			return
		}
		literal.WriteString("}")

		edits = append(edits, analysis.TextEdit{
			Pos:     arg.Pos(),
			End:     arg.End(),
			NewText: literal.Bytes(),
		})
		fix := analysis.SuggestedFix{
			Message:   fmt.Sprintf("Use %s", request.Obj().Name()),
			TextEdits: edits,
		}
		pass.Report(analysis.Diagnostic{
			Pos: call.Pos(),
			End: call.End(),
			Message: fmt.Sprintf(
				"%s.%s no longer takes a string id; pass a *%s so the request can carry a context and headers",
				service.Name(), fn.Name(), request.Obj().Name()),
			SuggestedFixes: []analysis.SuggestedFix{fix},
		})
	})
	return nil, nil
}

// sdkServiceReceiver reports whether fn is a method on a Chargebee SDK service
// type, and returns that type.
func sdkServiceReceiver(fn *types.Func) (*types.TypeName, bool) {
	sig, ok := fn.Type().(*types.Signature)
	if !ok || sig.Recv() == nil {
		return nil, false
	}
	recv := sig.Recv().Type()
	if ptr, ok := recv.(*types.Pointer); ok {
		recv = ptr.Elem()
	}
	named, ok := recv.(*types.Named)
	if !ok {
		return nil, false
	}
	obj := named.Obj()
	if obj.Pkg() == nil || !strings.HasPrefix(obj.Pkg().Path(), sdkPackagePrefix) {
		return nil, false
	}
	if !strings.HasSuffix(obj.Name(), "Service") {
		return nil, false
	}
	return obj, true
}

// requestParam returns the request type of a single-parameter SDK method along
// with its Id field, provided both look like generated request plumbing.
func requestParam(fn *types.Func, sdkPkg *types.Package) (*types.Named, *types.Var, bool) {
	sig := fn.Type().(*types.Signature)
	if sig.Params().Len() != 1 || sig.Variadic() {
		return nil, nil, false
	}
	ptr, ok := sig.Params().At(0).Type().(*types.Pointer)
	if !ok {
		return nil, nil, false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok || named.Obj().Pkg() != sdkPkg || !strings.HasSuffix(named.Obj().Name(), "Request") {
		return nil, nil, false
	}
	strct, ok := named.Underlying().(*types.Struct)
	if !ok {
		return nil, nil, false
	}
	for i := 0; i < strct.NumFields(); i++ {
		field := strct.Field(i)
		if field.Name() != "Id" || !field.Exported() {
			continue
		}
		if basic, ok := field.Type().Underlying().(*types.Basic); ok && basic.Kind() == types.String {
			return named, field, true
		}
	}
	return nil, nil, false
}

// importEdits returns the name under which pkg can be referenced in file,
// together with any edit needed to import it.
func importEdits(file *ast.File, pkg *types.Package, added map[*ast.File]bool) (string, []analysis.TextEdit, bool) {
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil || path != pkg.Path() {
			continue
		}
		if spec.Name == nil {
			return pkg.Name(), nil, true
		}
		switch spec.Name.Name {
		case ".":
			return "", nil, true
		case "_":
			continue
		default:
			return spec.Name.Name, nil, true
		}
	}
	if added[file] {
		return pkg.Name(), nil, true // another fix in this file already adds it
	}
	edit, ok := addImport(file, pkg.Path())
	if !ok {
		return "", nil, false
	}
	added[file] = true
	return pkg.Name(), []analysis.TextEdit{edit}, true
}

// addImport builds the edit that adds path to the file's imports.
func addImport(file *ast.File, path string) (analysis.TextEdit, bool) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			continue
		}
		if gen.Lparen.IsValid() {
			return analysis.TextEdit{
				Pos:     gen.Rparen,
				End:     gen.Rparen,
				NewText: []byte("\t" + strconv.Quote(path) + "\n"),
			}, true
		}
		return analysis.TextEdit{
			Pos:     gen.End(),
			End:     gen.End(),
			NewText: []byte("\nimport " + strconv.Quote(path)),
		}, true
	}
	if file.Name == nil {
		return analysis.TextEdit{}, false
	}
	return analysis.TextEdit{
		Pos:     file.Name.End(),
		End:     file.Name.End(),
		NewText: []byte("\n\nimport " + strconv.Quote(path)),
	}, true
}
