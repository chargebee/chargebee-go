// Command sdk-upgrade rewrites Chargebee Go SDK call sites that pass
// a bare string id to an operation which now takes a request object.
//
// Upgrade the SDK first, then run:
//
//	go run github.com/chargebee/chargebee-go/sdk-upgrade@latest -fix ./...
package main

import "golang.org/x/tools/go/analysis/singlechecker"

func main() { singlechecker.Main(Analyzer) }
