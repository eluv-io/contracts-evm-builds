#!/usr/bin/env bash


function fail() {
  echo "$1"
  exit 1
}

(
  cd contracts-go || fail "could not cd into contracts-go"
  go test -count=1 .
)

for dir in contracts-go/v*; do
  (
    [[ -d "$dir" ]] || fail "not a directory: $dir"
    cd "$dir" || exit
    go test -count=1 ./...
  )
done
echo "done"
