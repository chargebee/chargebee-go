build:
	go build ./...

test: build
	go test -v -race -coverprofile=coverage.out .
	go tool cover -html=coverage.out -o coverage.html

format:
	goimports-reviser -rm-unused -use-cache -format -apply-to-generated-files ./...

.PHONY: build test format
