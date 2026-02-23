build:
	go build -gcflags="-e" ./...

test: build
	go test -v -race ./...

format:
	@command -v goimports-reviser >/dev/null 2>&1 || { \
		echo "Installing goimports-reviser..."; \
		go install github.com/rinchsan/go-action-lint/cmd/goimports-reviser@latest; \
	}
	goimports-reviser -rm-unused -use-cache -format -apply-to-generated-files ./...

.PHONY: build test format
