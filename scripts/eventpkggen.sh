#!/usr/bin/env bash
set -Eeuo pipefail

curr_dir=$1
[ -z "$curr_dir" ] && fail "curr_dir not defined";
eventpkggen_dir=$(realpath "$curr_dir")

function build_eventpkggen() {
  # build eventpkggen
  echo "#### building eventpkggen command"
  pushd cmd > /dev/null
  go build -o ../eventpkggen ./eventpkggen
  popd > /dev/null
}

function run_eventpkggen() {

  if is_debug; then
      echo "$eventpkggen_dir/eventpkggen" "config.yaml"
  fi
  "$eventpkggen_dir/eventpkggen" "config.yaml"
}

function rm_eventpkggen_bin() {
  rm "$eventpkggen_dir/eventpkggen"
}


