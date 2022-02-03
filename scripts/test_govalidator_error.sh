#!/usr/bin/env bash

current_path=$(cd "$(dirname "${0}")" || exit 1; pwd)
cd "${current_path}"/.. || exit 1

export PATH="./build/bin:$PATH"

for path1 in xgo/tests/govalidatorexternal/test_invalid*proto; do
#  echo ${path1}
  protoc -I=. -I=./xgo --go_opt=paths=source_relative --govalidator_opt=paths=source_relative --go_out=. --govalidator_out=. "${path1}"
done

make check

#for path1 in xgo/tests/govalidatorexternal/test_error*proto; do
##  echo ${path1}
#  protoc -I=. -I=./xgo --go_opt=paths=source_relative --govalidator_opt=paths=source_relative --go_out=. --govalidator_out=. "${path1}"
##  echo $?  >/dev/null 2>&1
#  if [ $? != 1 ]; then
#    echo "Unexpected result with test file ${path1}"
#    exit 1
#  fi
#done
