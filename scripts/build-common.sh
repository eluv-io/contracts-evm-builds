#!/usr/bin/env bash
set -Eeuo pipefail

DEBUG="" # DEBUG is set via --debug

function fail() {
  echo "$1"
  exit 1
}

function is_debug() {
    if [ "$DEBUG" = "--debug" ]; then
        true
    else
        false
    fi
}

function download_solc() {
  local version=$1
  local jq_cmd=".releases[\"${version}\"]"
  local binary

  # cat solc_macosx-amd64_list.json | jq '.releases["0.3.6"]'
  binary=$( curl -q -s "https://raw.githubusercontent.com/ethereum/solc-bin/gh-pages/${platform}/list.json" | jq "${jq_cmd}" )
  binary=${binary//'+'/'%2B'}
  binary=${binary%\"}
  binary=${binary#\"}

  mkdir -p "${solc_folder}"/"${version}"
  curl -q -s -o "${solc_folder}"/"${version}"/solc "https://binaries.soliditylang.org/${platform}/${binary}"
  chmod ugo+x "${solc_folder}"/"${version}"/solc
}

# require_solc makes sure the required version of solc in $1 is available.
# Unless a solc binary is found in folder '_bin/${platform}/solc/${version},
# the script downloads the solc binary from 'binaries.soliditylang.org'
function require_solc() {
  solc_version="$1"

  if is_debug; then
    echo "verifying solc ${solc_version}"
  fi

  if ! [ -f "${solc_folder}"/"${solc_version}"/solc ]; then
    echo "downloading solc ${solc_version}"
    download_solc "${solc_version}"
  fi

  # if the download file contains 'xml version' it's probably an error message
  if head -n 1 "${solc_folder}/${solc_version}/solc" | grep -q "xml version" ; then
    # show it entirely
    cat "${solc_folder}/${solc_version}/solc"
    echo ""
    echo "WARN: '${solc_version}' is probably an invalid version; see above message."
    echo ""
  fi

  # run solc --version to check the binary
  if is_debug; then
      "${solc_folder}/${solc_version}/solc" --version | grep "Version: ${solc_version}+"
  else
      "${solc_folder}/${solc_version}/solc" --version | grep "Version: ${solc_version}+" > /dev/null
  fi
}

function install_gojq() {
  if ! [[ -f "${gojq_folder}/gojq" ]]; then
      export GOBIN="${gojq_folder}"
      go install github.com/itchyny/gojq/cmd/gojq@v0.12.8
  fi
  echo "${gojq_folder}/gojq"
}

function solc_dir() {
  echo "${solc_folder}"
}

curr_dir=$1
[ -z "$curr_dir" ] && fail "curr_dir not defined";

if [[ $# -gt 1 ]]; then
    DEBUG="$2"
fi


platform=""
case $(uname) in
    Darwin*)
        platform=macosx-amd64
        ;;
    Linux*)
        platform=linux-amd64
        ;;
    *)
        echo "Not sure what environment this is.  Cannot setup solc"
        exit 1
        ;;
esac

bin_folder="_bin"
solc_folder=${curr_dir}/${bin_folder}/${platform}/solc
mkdir -p "${solc_folder}"

gojq_folder=${curr_dir}/${bin_folder}/${platform}/gojq
mkdir -p "${gojq_folder}"

# invoked from config.sh
#install_gojq
