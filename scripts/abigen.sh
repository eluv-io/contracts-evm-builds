#!/usr/bin/env bash
set -Eeuo pipefail

function build_abigen() {
  # build abigen
  echo "#### building abigen command"
  pushd cmd > /dev/null
  go build -o ../abigen ./abigen
  popd > /dev/null
}

function run_abigen() {
  local combined_json=$1
  local pkg=$2
  local out_file=$3
  if is_debug; then
      echo "$abigen_dir/abigen" --pkg=${pkg} --out "${out_file}" --combined-json ${combined_json}
  fi
  "$abigen_dir/abigen" --pkg=${pkg} --out "${out_file}" --combined-json ${combined_json}
}
