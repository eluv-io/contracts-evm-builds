#!/usr/bin/env bash
set -Eeuo pipefail

function read_config() {
   local yaml_file=$1
   yaml_config=$( cat $yaml_file )
}
