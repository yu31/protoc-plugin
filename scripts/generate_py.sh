#!/usr/bin/env bash
# Generate protobuf code for python

if ! [[ "$0" =~ scripts/generate_py.sh ]]; then
	echo "must be run from repository root"
	exit 255
fi

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)

# load project env.
if [ -f "./project_env.sh" ]; then
  . ./project_env.sh
fi

# install dep package.
sh "${current_path}/ensure_dep.sh"

cd "${current_path}"/.. || exit 1

GOPATH=$(go env GOPATH)

output_dir="./xpy/src/pb"

# To avoids invalid code residue.
/bin/rm -fr "$output_dir"
mkdir -p "$output_dir"

for file in proto/*.proto; do
    # generate python code.
  protoc -I. -I./proto -I"${GOPATH}"/src --python_out="$output_dir" "$file"
done

