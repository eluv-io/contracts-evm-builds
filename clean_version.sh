#!/usr/bin/env bash

function usage() {
    echo "usage: $0 <vx.y.z>"
    echo
    echo "Remove output of version vx.y.z"
    echo ""
}

function fail() {
  echo "$1"
  exit 1
}

function is_tag(){
    v=$1
    firstCharacter=${v:0:1}
    if [ "$firstCharacter" = "v" ]; then
        true
    else
        false
    fi
}

if [[ $# -eq 0 ]]; then
    usage
    exit 1
fi

version=$1

if ! is_tag "${version}" ; then
  fail "not a valid version: $version"
fi

echo "removing dist/${version}"
rm -rf "dist/${version}"

echo "removing files in contracts-go/${version}"
rm "contracts-go/${version}/contracts/contracts.go"
rm "contracts-go/${version}/legacy/contracts.go"



echo "done"
