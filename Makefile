VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

GOPATH := $(shell go env GOPATH)
GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint

.PHONY: build release test testv testc testf clean help lint

build:
	go build -o fog .

release:
	go build $(LDFLAGS) -trimpath -o fog .

test:
	go test ./...

testv:
	go test -v ./...

testc:
	go test -cover ./...

testf:
	go test -v -run TestFilterDebug ./filters

clean:
	rm -f fog

lint:
	@if [ ! -x "$(GOLANGCI_LINT)" ]; then \
		echo "golangci-lint not found. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi
	$(GOLANGCI_LINT) run

help:
	@echo "Available targets:"
	@echo "  build    - Build binary"
	@echo "  release  - Build release binary (stripped, with version)"
	@echo "  test     - Run all tests"
	@echo "  testv    - Run all tests verbose"
	@echo "  testc    - Run tests with coverage"
	@echo "  testf    - Run filter debug test"
	@echo "  clean    - Remove binary"
	@echo "  lint     - Run golangci-lint"

%:
	@echo "Unknown target: $@"
	@echo "Run 'make help' for available targets"
	@exit 1