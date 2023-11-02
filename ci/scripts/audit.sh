#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/otel.go
  make audit
popd