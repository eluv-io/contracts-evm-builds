#!/usr/bin/env bash
set -Eeuo pipefail

yaml_config=""
gojq=""

function read_config() {
   local yaml_file=$1
   yaml_config=$( cat "$yaml_file" )
}

function config {
  echo "$yaml_config" | ${gojq} "--yaml-input" "$@"
}

function configString {
  config --raw-output "$@"
}

gojq=$(install_gojq)