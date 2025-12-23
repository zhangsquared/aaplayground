# Variables
BINARY_NAME=aaplayground
GO_FILES=$(shell find . -type f -name '*.go')

## help: Help for the Makefile
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

## fmt: Format all Go source files
fmt:
	go fmt ./...

## test: Run all tests with coverage
test:
	go test -v -cover ./...

## run: Format, build, and run the application
run: fmt
	go run main.go

## build: Build the binary
build:
	go build -o bin/$(BINARY_NAME) main.go

## clean: Remove build artifacts
clean:
	rm -rf bin/
	go clean

.PHONY: help fmt test run build clean
