#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

for cmd in cmd/protoc-gen*; do
  name="${cmd/cmd\/}"
  go build -o ./build/bin/"$name" ./"$cmd"
done
