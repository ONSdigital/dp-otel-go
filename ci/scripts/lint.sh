#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/otel.go
# Install golangci-lint
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
  make lint
popd
