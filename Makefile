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
	go build ./...

.PHONY: test
test: 
	go test -race -cover 

.PHONY: lint
lint:
	golangci-lint --deadline=10m --fast --enable=gosec --enable=gocritic --enable=gofmt --enable=gocyclo --enable=bodyclose --enable=gocognit run
