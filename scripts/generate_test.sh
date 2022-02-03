#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

export PATH="./build/bin:$PATH"

for cmd in xgo/cmd/protoc-gen*; do
  plugin="${cmd/xgo\/cmd\/protoc-gen-}"

  for path2 in xgo/tests/*/"${plugin}"*proto; do
    protoc -I=. -I=./xgo --go_opt=paths=source_relative --"${plugin}"_opt=paths=source_relative --go_out=. --"${plugin}"_out=. "${path2}"
  done
done
