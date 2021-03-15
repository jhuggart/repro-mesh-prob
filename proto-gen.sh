#!/bin/bash

FILES="proto/*.proto"

if [ ! -d gen/go/customers ]; then
  mkdir -p gen/go/customers
fi

if [ ! -d public/docs ]; then
  mkdir -p public/docs
fi

for file in $FILES; do
  protoc \
    -I ${GOPATH}/src \
    -I proto \
    --go_opt=module=github.com/jhuggart/repro-mesh-prob \
    --go_out=gen/go \
    --go-grpc_opt paths=source_relative \
    --go-grpc_out=gen/go/customers \
    $file
done