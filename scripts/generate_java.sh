#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)

# install dep package.
sh "${current_path}/ensure_dep.sh"

# load project env.
if [ -f "./project_env.sh" ]; then
  . ./project_env.sh
fi

cd "${current_path}"/.. || exit 1

output_dir="./xjava/src/main/java"

# To avoids invalid code residue.
/bin/rm -fr "$output_dir/protoc/pb"
mkdir -p "$output_dir"

for f in proto/*.proto; do
  # code for java
  protoc -I./proto --java_out="$output_dir" "$f"
done

#if git status |grep 'src/main/java' >/dev/null; then
#  echo "mvn clean package deploy"
#  mvn clean package deploy >/dev/null 2>&1 || exit $?
#else
#  echo "no java code generated, skip mvn deploy"
#fi
