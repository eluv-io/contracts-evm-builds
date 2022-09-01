#!/usr/bin/env bash
set -Eeuo pipefail



run_latest="true"    # run latest per default

function usage() {
    echo "usage: $0 -all|-latest [--debug]"
    echo
    echo "Build contracts versions defined in 'config.yaml'"
    echo "-all:    build all of contracts versions."
    echo "-latest: build the last one"
    echo ""
}

if [[ $# -eq 0 ]]; then
    usage
    exit 1
fi

case "$1" in
    -all | --all)
        run_latest="false"
        shift
        ;;
    -latest | --latest)
        run_latest="true"
        shift
        ;;
    *)
        echo "unknown option \"$1\""
        usage
        exit 1
        ;;
esac
# Debug
DEBUG=""
if [[ $# -gt 0 ]]; then
    DEBUG="$1"
fi


curr_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
gojq=""

source scripts/build-common.sh $curr_dir
source scripts/solc.sh
source scripts/abigen.sh
source scripts/config.sh

yaml_config=""
read_config config.yaml


function is_tag(){
    v=$(eval echo "$1") # unquote the value
    firstCharacter=${v:0:1}
    if [ $firstCharacter = "v" ]; then
        true
    else
        false
    fi
}

#
# verify_versions verifies that versions are tags, i.e. a string starting with
# 'v' except the last version which can be a commit hash
#
function verify_versions() {
    local length=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions | length" )
    local last=$(( $length-1 ))
    local version=""
    local solc=""
    local use=""

    for ((idx=0; idx<${length}; ++idx)); do
        version=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].tag" )
        solc=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].solc" )
        use=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].use" )
        if is_debug; then
            echo "$idx" "${version}" "${solc}" "${use}"
        fi

        if ! is_tag "${version}"; then
            fail "invalid version '$idx ${version} ${solc}' - tag must start with v"
        fi
        if [[ $idx != $last ]] && ! [[ "${use}" = "null" ]]; then
            fail "invalid version '$idx ${version} ${solc} ${use}' - only the last version may use a commit hash"
        fi

    done
}
verify_versions

echo "build version : "$( echo "$yaml_config" | ${gojq} "--yaml-input" ".build.version" )
echo "versions count: "$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions | length" )
if is_debug; then
    echo "tags          : "$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[].tag" )
fi
echo ""

# build 'our' agigen
build_abigen


# checkout 'known' versions of contracts in folder _build_contracts_go
# and produce go-bindings in ../../contracts-go
function build_version() {
  local tag=$(eval echo $1)
  local solc_version=$(eval echo $2)
  local commit_hash=$(eval echo "$3")

  local pkg="contracts_$tag"
  local real_tag=$tag

  echo ""
  echo "== $tag =="
  if ! [[ "${commit_hash}" = "null" ]] ; then
      real_tag=$(eval echo "$commit_hash")
  fi
  pkg=$(echo "${pkg//[.]/_}") # replace '.' with '_' in package name
  pkg=$(eval echo "$pkg")

  mkdir -p _build_contracts_go
  pushd _build_contracts_go > /dev/null

  if [ ! -d contracts-evm ]; then
    echo "#### cloning contracts repo"
    git clone https://github.com/eluv-io/contracts-evm.git
  fi
  cd contracts-evm/
  git fetch -f --tags --all
  git checkout "$real_tag"

  #
  # build 'contracts'
  #
  local solidity_file=src/tenant/Tenant.sol

  mkdir -p "$dist_dir/$tag/contracts"
  run_solc "$solc_version" "$solidity_file" "$dist_dir/$tag/contracts"

  mkdir -p "../../contracts-go/$tag/contracts"
  run_abigen "$dist_dir/$tag/contracts/combined-json/combined.json" "$pkg" "../../contracts-go/$tag/contracts/contracts.go"

  #
  # build 'legacy'
  #
  solidity_file=src/legacy/main.sol
  mkdir -p "$dist_dir/$tag/legacy"
  run_solc "$solc_version" "$solidity_file" "$dist_dir/$tag/legacy"

  mkdir -p "../../contracts-go/$tag/legacy"
  run_abigen "$dist_dir/$tag/legacy/combined-json/combined.json" "$pkg" "../../contracts-go/$tag/legacy/contracts.go"


  popd > /dev/null
}


function build_all() {
    local length=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions | length" )
    local last_index=$(( $length-1 ))
    local tag=""
    local solc=""
    local use=""

    for ((idx=0; idx<${length}; ++idx)); do
        tag=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].tag" )
        solc=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].solc" )
        use=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$idx].use" )

        if is_debug; then
            echo ""
            echo "building: ${i} ${version} ${solc} ${use}"
        fi
        build_version "${tag}" "${solc}" "${use}"
    done
}

function build_latest() {
    local length=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions | length" )
    local last_index=$(( $length-1 ))
    local tag=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$last_index].tag" )
    local solc=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$last_index].solc" )
    local use=$( echo "$yaml_config" | ${gojq} "--yaml-input" ".versions[$last_index].use" )

    echo "building latest"
    build_version "${tag}" "${solc}" "${use}"
}


if [[ "${run_latest}" == "true" ]]; then
    build_latest
else
    build_all
fi

echo "done"
