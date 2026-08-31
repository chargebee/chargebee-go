// The fixer is its own module so that its toolchain and analysis dependencies
// stay out of the SDK's dependency graph, which servers import.
module github.com/chargebee/chargebee-go/sdk-upgrade

go 1.27.0

require golang.org/x/tools v0.38.0

require (
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
)
