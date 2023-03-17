#!/usr/bin/env bash
set -Eeuo pipefail


run_latest=true # when "true" run latest only
curr_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
dist_dir=$curr_dir/dist
out_dir_set=()

function usage() {
    echo "usage: $0 -all|-latest [--debug]"
    echo
    echo "Build contracts versions defined in 'config.yaml'"
    echo "-all:    build all of contracts versions."
    echo "-latest: build the last one"
    echo ""
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

#
# verify_versions verifies that versions are tags, i.e. a string starting with
# 'v' except the last version which can be a commit hash
# It also makes sure the required version of solc are available.
function verify_versions() {
    local length
    local last
    local tag=""
    local solc=""
    local use=""

    length=$( config ".versions | length" )
    last=$((length - 1))
    last_tag=$(configString ".versions[$last].tag")

    if is_debug; then
        echo "verifying config with $length versions"
    fi

    for ((idx=0; idx<length; ++idx)); do
        tag=$(configString ".versions[$idx].tag")
        solc=$(configString ".versions[$idx].solc")
        use=$(configString ".versions[$idx].use")
        if is_debug; then
            echo "$idx : ${tag} - ${solc} - ${use}"
        fi

        if ! is_tag "${tag}"; then
            fail "invalid version '$idx : ${tag} - ${solc}' - tag must start with v"
        fi

        if [[ "$idx" != "$last" && "${tag}" != "${last_tag}" && "${use}" != "null" ]]; then
            fail "invalid version '$idx : ${tag} - ${solc} - ${use}' - only the last version may use a commit hash"
        fi

        if [[ "$idx" = "$last" || "${run_latest}" != "true" ]]; then
            require_solc "${solc}"
        fi

    done
}

# checkout 'known' versions of contracts in folder _build_contracts_go
# and produce go-bindings in ../../contracts-go
function build_version() {
  local tag=$1
  local solc_version=$2
  local commit_hash=$3
  local solidity_file=$4
  local output_folder=$5
  local dependencies=$6

  local real_tag=$tag
  echo ""
  echo "== $tag =="
  if ! [[ "${commit_hash}" = "null" ]] ; then
      real_tag=$(eval echo "$commit_hash")
  fi

  mkdir -p _build_contracts_go
  (
    cd _build_contracts_go

    if [ ! -d contracts-evm ]; then
      echo "#### cloning contracts repo"
      git clone https://github.com/eluv-io/contracts-evm.git
    fi
    cd contracts-evm/
    git fetch -f --tags --all
    git checkout "$real_tag"

    # to install dependencies
    git submodule init
    git submodule update

    #
    # build 'contracts'
    #
    build_package "${solidity_file}" "${output_folder}" "${tag}" "${dependencies}"
 )
}

# build all tags
function build_all() {
    local length
    local tag=""
    local solc=""
    local use=""
    local solidity_file=""
    local output_folder=""
    local dependencies=""

    length=$( config ".versions | length" )

    for ((idx=0; idx<length; ++idx)); do
        tag=$( configString ".versions[$idx].tag" )
        solc=$( configString ".versions[$idx].solc" )
        use=$( configString ".versions[$idx].use" )
        solidity_file=$( configString ".versions[$idx].solidity_file" )
        output_folder=$( configString ".versions[$idx].output_folder" )
        dependencies=$( configString ".versions[$idx].dependencies" )

        if is_debug; then
            echo ""
            echo "building: ${idx} ${tag} ${solc} ${use} ${solidity_file} ${output_folder} ${dependencies}"
        fi
        build_version "${tag}" "${solc}" "${use}" "${solidity_file}" "${output_folder}" "${dependencies}"
    done
}

# build latest tag
function build_latest() {
    local length
    local last_index
    local tag=""
    local solc=""
    local use=""
    local solidity_file=""
    local output_folder=""
    local dependencies=""

    length=$( config ".versions | length" )
    last_index=$(( length-1 ))
    last_index_tag=$( configString ".versions[$last_index].tag" )

    for ((idx=0; idx<length; ++idx)); do
      tag=$( configString ".versions[$idx].tag" )
      if [[ $tag == "$last_index_tag" ]]; then
        solc=$( configString ".versions[$idx].solc" )
        use=$( configString ".versions[$idx].use" )
        solidity_file=$( configString ".versions[$idx].solidity_file" )
        output_folder=$( configString ".versions[$idx].output_folder" )
        dependencies=$( configString ".versions[$idx].dependencies" )

        if is_debug; then
            echo ""
            echo "building: ${idx} ${tag} ${solc} ${use} ${solidity_file} ${output_folder} ${dependencies}"
        fi
        build_version "${tag}" "${solc}" "${use}" "${solidity_file}" "${output_folder}" "${dependencies}"
      fi
    done
}

function build_package() {
  local solidity_file=$1
  local output_folder=$2
  local tag=$3
  local dependencies=$4

  local pkg="${output_folder}_${tag}"
  pkg="${pkg//[.]/_}" # replace '.' with '_' in package name
  pkg=$(eval echo "$pkg")
  event_pkg="${output_folder}/${output_folder}_go"

  local go_binding_dir="../../${output_folder}/${output_folder}_go"
  mkdir -p "${go_binding_dir}/${tag}"
  mkdir -p "${dist_dir}/${output_folder}/${tag}"

  if [[ -z "${dependencies}" ]]; then
    run_solc "$solc_version" "$solidity_file" "${dist_dir}/${output_folder}/${tag}"
  else
    run_solc "$solc_version" "$solidity_file" "${dist_dir}/${output_folder}/${tag}" "${dependencies}"
  fi

  run_abigen "${dist_dir}/${output_folder}/${tag}/combined-json/combined.json" "$pkg" "$event_pkg" "${go_binding_dir}/${tag}/contracts.go"
}

function run_gomod_tidy(){
  length=$( config ".versions | length" )

  for ((idx=0; idx<length; ++idx)); do
    output_folder=$( configString ".versions[$idx].output_folder" )
    output_folder="${output_folder}/${output_folder}_go"
    if [ ${#out_dir_set[@]} -eq 0 ]; then
        out_dir_set+=( "${output_folder}" )
        ( cd "${output_folder}" && go mod tidy )
    else
       for f in "${out_dir_set[@]}"; do
         if [[ $f == "${output_folder}" ]]; then
           break
         fi
         out_dir_set+=( "${output_folder}" )
         ( cd "${output_folder}" && go mod tidy )
       done
    fi
  done
}

if [[ $# -eq 0 ]]; then
    usage
    exit 1
fi

case "$1" in
    -all | --all)
        run_latest=false
        shift
        ;;
    -latest | --latest)
        run_latest=true
        shift
        ;;
    *)
        echo "unknown option \"$1\""
        usage
        exit 1
        ;;
esac

source scripts/build-common.sh "$curr_dir" "$@"
source scripts/solc.sh
source scripts/abigen.sh "$curr_dir"
source scripts/eventpkggen.sh "$curr_dir"
source scripts/config.sh

read_config config.yaml

verify_versions

echo ""
if is_debug; then
    echo "build tool version : $( configString '.build.version' )"
    echo "tags               : $( configString '.versions | length' )"
    echo ""
fi

# build 'our' abigen
build_abigen
# build eventpkggen
build_eventpkggen

dist_dir="../../dist"
mkdir -p "$dist_dir"

if [ "${run_latest}" = "true" ]; then
    build_latest
else
    build_all
fi

run_eventpkggen

#run go mod tidy
run_gomod_tidy

echo "done"