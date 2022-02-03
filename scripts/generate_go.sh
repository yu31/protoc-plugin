#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)

# install dep package.
sh "${current_path}/ensure_dep.sh"

# load project env.
if [ -f "./project_env.sh" ]; then
  . ./project_env.sh
fi

cd "${current_path}"/.. || exit 1

output_dir="./xgo/pb"
# To avoids invalid code residue.
/bin/rm -fr "$output_dir"
mkdir -p "$output_dir"

MODULE="$(go list)"

for f in proto/*.proto; do
  protoc -I./proto --go_opt=module="${MODULE}" --go_out=. "$f"
done
