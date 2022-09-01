#!/usr/bin/env bash
set -Eeuo pipefail

function run_solc() {
    # run_solc runs solc with various command line parameters and ouput the
    # results in conventional subfolders: abi,bin and combined-json.
    # 1: solc version
    # 2: contract file (.sol)
    # 3: output folder

    local version=$1
    solc_bin="${solc_folder}/${version}/solc"
    local bsc_name=$(basename "${2}")
    local out=$3

    if is_debug; then
        echo "${solc_bin} ${2} --abi --hashes --optimize -o ${out}/abi --overwrite"
    fi
    ${solc_bin} "${2}" --abi --hashes --optimize -o "${out}/abi" --overwrite
    [ $? = 0 ] || fail "failed generating ABI and function hashes of ${bsc_name} to ${out}/abi"

    if is_debug; then
        echo "${solc_bin} ${2} --bin --optimize -o ${out}/bin --overwrite"
    fi
    ${solc_bin} "${2}" --bin --optimize -o "${out}/bin" --overwrite
    [ $? = 0 ] || fail "failed generating BIN of ${bsc_name} to ${out}/bin"


    if is_debug; then
        echo "${solc_bin} ${2} --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ${out}/combined-json --overwrite"
    fi
    ${solc_bin} "${2}" --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o "${out}/combined-json" --overwrite
    [ $? = 0 ] || fail "failed producing combined json for $bsc_name"

}
