#!/usr/bin/env bash


function usage() {
    echo "usage: $0 [--debug]"
    echo
    echo "Run all go unit-tests"
    echo ""
}

DEBUG="" # DEBUG is set via --debug

if [[ $# -gt 0 ]]; then
  case "$1" in
      -h | --help)
          usage
          exit 1
          ;;
      -d | --debug)
          DEBUG="--debug"
          shift
          ;;
      *)
          echo "unknown option \"$1\""
          usage
          exit 1
          ;;
  esac
fi

function is_debug() {
    if [ "$DEBUG" = "--debug" ]; then
        true
    else
        false
    fi
}

if [[ $# -gt 0 ]]; then
    DEBUG="$1"
fi

function fail() {
  echo "$1"
  exit 1
}

(
  cd contracts-go || fail "could not cd into contracts-go"
  go test -count=1 .
)

for dir in commerce/commerce_go elv_token/elv_token_go; do
  printf '\n%s\n' $dir
  (
    [[ -d "$dir" ]] || fail "not a directory: $dir!"
    cd "$dir" || exit
    flags=""
    if is_debug; then
      flags="-v"
    fi
    go test $flags -count=1 ./...
  )
done
echo "done"