#!/usr/bin/env bash

#export GO111MODULE="off"
#go get github.com/gogo/protobuf/protoc-gen-gogo


# check the grpc plugin is installed.
if ! type protoc > /dev/null 2>&1; then
  echo "Error: the plugin <protoc> not install, please install it before running"
  exit 1
fi

if ! type protoc-gen-go > /dev/null 2>&1; then
  echo "Error: the plugin <protoc-gen-go> not install, please install it before running"
  exit 1
fi

# check the plugin version.
if [[ $(protoc --version | cut -f2 -d' ') != "3.19.3" ]]; then
  echo "Error: could not find protoc 3.19.3, is it installed in you PATH?"
  exit 1
fi

if [[ $(protoc-gen-go --version 2>&1 | cut -f2 -d' ') != "v1.27.1" ]]; then
  echo "Error: could not find protoc-gen-go v1.27.1, is it installed in you PATH?"
  exit 1
fi
