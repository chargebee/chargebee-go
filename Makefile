build:
	go build -gcflags="-e" ./...

test: build
	go test -v -race ./...

format:
	goimports-reviser -rm-unused -use-cache -format -apply-to-generated-files ./...

.PHONY: build test format
