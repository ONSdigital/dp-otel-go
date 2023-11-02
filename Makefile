SHELL=bash

BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

.PHONY: all
all: audit test build

.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: build
build: 
	go build -tags 'production' -o $(BINPATH)/dp-otel-go -ldflags="-X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.Version=$(VERSION)'"

.PHONY: debug
debug:
	go build -tags 'debug' -race -o $(BINPATH)/dp-otel-go -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: test
test: 
	go test -race -cover -tags 'production' ./...




