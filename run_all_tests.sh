#!/usr/bin/env bash

pushd contracts-go > /dev/null
go test -count=1 .
popd > /dev/null

for dir in contracts-go/v*; do
    pushd $dir > /dev/null
    go test -count=1 ./...
    popd > /dev/null
done
echo "done"
