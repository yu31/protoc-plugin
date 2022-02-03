#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

cp -frp ./build/bin/* "$GOPATH/bin/"

#go install github.com/yu31/proto-go-plugin/cmd/protoc-gen-gosql
