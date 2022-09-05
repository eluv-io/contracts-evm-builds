#!/usr/bin/env bash
set -Eeuo pipefail


run_latest="true" # when "true" run latest only
curr_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
dist_dir=$curr_dir/dist

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
        if [[ $idx != "$last" ]] && ! [[ "${use}" = "null" ]]; then
            fail "invalid version '$idx : ${tag} - ${solc} - ${use}' - only the last version may use a commit hash"
        fi

        if [[ $idx = "$last" ]] || ! [[ "${run_latest}" = "true" ]]; then
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

  local pkg="contracts_$tag"
  local real_tag=$tag

  echo ""
  echo "== $tag =="
  if ! [[ "${commit_hash}" = "null" ]] ; then
      real_tag=$(eval echo "$commit_hash")
  fi
  pkg="${pkg//[.]/_}" # replace '.' with '_' in package name
  pkg=$(eval echo "$pkg")

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
  )
}

function build_all() {
    local length
    local tag=""
    local solc=""
    local use=""

    length=$( config ".versions | length" )

    for ((idx=0; idx<length; ++idx)); do
        tag=$( configString ".versions[$idx].tag" )
        solc=$( configString ".versions[$idx].solc" )
        use=$( configString ".versions[$idx].use" )

        if is_debug; then
            echo ""
            echo "building: ${idx} ${tag} ${solc} ${use}"
        fi
        build_version "${tag}" "${solc}" "${use}"
    done
}

function build_latest() {
    local length
    local last_index
    local tag=""
    local solc=""
    local use=""

    length=$( config ".versions | length" )
    last_index=$(( length-1 ))
    tag=$( configString ".versions[$last_index].tag" )
    solc=$( configString ".versions[$last_index].solc" )
    use=$( configString ".versions[$last_index].use" )

    echo "building latest"
    build_version "${tag}" "${solc}" "${use}"
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

source scripts/build-common.sh "$curr_dir" "$@"
source scripts/solc.sh
source scripts/abigen.sh "$curr_dir"
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

mkdir -p "$dist_dir"

if [[ "${run_latest}" == "true" ]]; then
    build_latest
else
    build_all
fi

echo "done"
