#!/usr/bin/env bash
set -Eeuo pipefail

function run_solc() {
    # run_solc runs solc with various command line parameters and writes the
    # results in conventional sub-folders: abi,bin and combined-json.
    # 1: solc version
    # 2: contract file (.sol)
    # 3: output folder
    # 4: external library path like openzeppelin

    local version=$1
    solc_bin="$(solc_dir)/${version}/solc"
    local bsc_name
    local out=$3
    local ext_lib

    local ext_lib_enabled
    if [ $# -gt 3 ]; then
      ext_lib_enabled=true
      ext_lib=$4
    fi

    bsc_name=$(basename "${2}")

    if is_debug; then
       if [ ${ext_lib_enabled} ]; then
         echo "${solc_bin} ${ext_lib} ${2} --abi --hashes --optimize -o ${out}/abi --overwrite"
       else
           echo "${solc_bin} ${2} --abi --hashes --optimize -o ${out}/abi --overwrite"
       fi
    fi

    if [ ${ext_lib_enabled} ]; then
        if ! ${solc_bin} "${ext_lib}" "${2}" --abi --hashes --optimize -o "${out}/abi" --overwrite; then
              fail "failed generating ABI and function hashes of ${bsc_name} to ${out}/abi"
        fi
    else
        if ! ${solc_bin} "${2}" --abi --hashes --optimize -o "${out}/abi" --overwrite; then
              fail "failed generating ABI and function hashes of ${bsc_name} to ${out}/abi"
        fi
    fi


    if is_debug; then
      if [ ${ext_lib_enabled} ]; then
        echo "${solc_bin} ${ext_lib} ${2} --bin --optimize -o ${out}/bin --overwrite"
      else
        echo "${solc_bin} ${2} --bin --optimize -o ${out}/bin --overwrite"
      fi
    fi

    if [ ${ext_lib_enabled} ]; then
        if ! ${solc_bin} "${ext_lib}" "${2}" --bin --optimize -o "${out}/bin" --overwrite; then
          fail "failed generating BIN of ${bsc_name} to ${out}/bin"
        fi
    else
        if ! ${solc_bin} "${2}" --bin --optimize -o "${out}/bin" --overwrite; then
          fail "failed generating BIN of ${bsc_name} to ${out}/bin"
        fi
    fi


    if is_debug; then
      if [ ${ext_lib_enabled} ]; then
            echo "${solc_bin} ${ext_lib} ${2} --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ${out}/combined-json --overwrite"
      else
            echo "${solc_bin} ${2} --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o ${out}/combined-json --overwrite"
      fi
    fi

    if [ ${ext_lib_enabled} ]; then
          if ! ${solc_bin} ${ext_lib} "${2}" --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o "${out}/combined-json" --overwrite; then
                fail "failed producing combined json for $bsc_name"
          fi
    else
          if ! ${solc_bin} "${2}" --combined-json bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes --optimize -o "${out}/combined-json" --overwrite; then
                fail "failed producing combined json for $bsc_name"
          fi
    fi

}
