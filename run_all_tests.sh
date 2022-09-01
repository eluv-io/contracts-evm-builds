#!/usr/bin/env bash

for dir in contracts-go/v*; do
    pushd $dir > /dev/null
    go test ./...
    popd > /dev/null
done
echo "done"
