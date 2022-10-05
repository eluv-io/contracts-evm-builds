#!/usr/bin/env bash
set -Eeuo pipefail

function run_solc() {
    # run_solc runs solc with various command line parameters and writes the
    # results in conventional sub-folders: abi,bin and combined-json.
    # 1: solc version
    # 2: contract file (.sol)
    # 3: output folder

    local version=$1
    solc_bin="$(solc_dir)/${version}/solc"
    local bsc_name
    local out=$3

    bsc_name=$(basename "${2}")

    if is_debug; then
        echo "${solc_bin} ${2} --abi --hashes --optimize -o ${out}/abi --overwrite"
    fi
    if ! ${solc_bin} "${2}" --abi --hashes --optimize -o "${out}/abi" --overwrite; then
      fail "failed generating ABI and function hashes of ${bsc_name} to ${out}/abi"
    fi

    if is_debug; then
        echo "${solc_bin} ${2} --bin --optimize -o ${out}/bin --overwrite"
    fi
    if ! ${solc_bin} "${2}" --bin --optimize -o "${out}/bin" --overwrite; then
      fail "failed generating BIN of ${bsc_name} to ${out}/bin"
    fi

    if is_debug; then
        echo "${solc_bin} ${2} --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ${out}/combined-json --overwrite"
    fi
    if ! ${solc_bin} "${2}" --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o "${out}/combined-json" --overwrite; then
      fail "failed producing combined json for $bsc_name"
    fi

}
