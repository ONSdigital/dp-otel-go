#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-otel-go
# Install golangci-lint
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
  make lint
popd
