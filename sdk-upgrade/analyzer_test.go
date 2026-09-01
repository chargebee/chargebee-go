package main

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/txtar"
)

func TestAnalyzer(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, unpackTestdata(t), Analyzer, "callsites", "noimport")
}

// unpackTestdata writes the fixture archive to a temporary directory and
// returns it, for use as the GOPATH that analysistest runs against.
func unpackTestdata(t *testing.T) string {
	t.Helper()

	archive, err := txtar.ParseFile(filepath.Join("testdata", "analyzer.txtar"))
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	for _, file := range archive.Files {
		path := filepath.Join(dir, filepath.FromSlash(file.Name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, file.Data, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}
