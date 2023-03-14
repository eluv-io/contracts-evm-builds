# Building evm contracts bindings

Eluvio Content Fabric Ethereum/EVM contract builds for all released versions

This project provides a mechanism for building ABI and `go` bindings for specific - and well identified - versions
of contracts.

## configuration

Builds are driven by a configuration file in yaml: see the example below.

The configuration is mainly a list of `versions`. A version is made of:

* `solc`: the version of the solidity compiler to use
* `tag` : tag or **future** tag
    * tags *must* start with a 'v'
    * on the last version **only** tag may designate a _future_ tag when field `use` points to a commit hash
* `use` : a commit hash. Must not be set except within the last version, when `tag` is intended to be a _future_ tag.
* `solidity_file` : source file in the `contracts-evm` repo.
* `dependencies` : comma separated list of dependencies. Example : `dependencies: "openzeppelin/=lib/openzeppelin-contracts/"`
* `output_folder` : folder name to put the raw output under dist and to generate go_bindings. The go_bindings will be generated in `${output_folder}/${output_folder}_go` folder.

Within the last version, `tag` might be used to specify the value of the _future_ tag that will be used for the version.
In that case the field `use` must specify the commit hash to use.

When done changing the code of the contracts and the new version is tagged, remove the `use` field.

```
---
build:
  version: v1.0.0  # version of the build is informative

# list of versions
versions:
  -
    tag: "v0.0.1" # tag to retrieve the contract version Tags must start with v
    solc: "0.5.4" # version of the solidity compiler to use
    solidity_file: "src/legacy/main.sol"
    output_folder: "legacy"
  -
    tag: 'v1.0.0'
    solc: '0.8.13'
    solidity_file: "<PATH/TO/FILE>"
    output_folder: "<PATH/TO/FOLDER>"
  -
    tag: "v1.2.0"  # only on last version tag may reflect the future tag when 'use' is a commit hash 
    solc: "0.8.13"
    use: "591bce9bce2ae22a4498100d046a07805f63a277" # only last version may use a commit hash
    solidity_file: "src/commerce/Payment.sol"
    output_folder: "commerce"
    dependencies: "openzeppelin/=lib/openzeppelin-contracts/"
```

## build script

The build script generates go-bindings, abi, bin and combined-json for all the contract versions if `-all` flag is provided or only the latest tag if `-latest` flag is provided.

```
usage: ./build.sh -all|-latest [--debug]

Build contracts versions defined in 'config.yaml'
-all:    build all of contracts versions.
-latest: build the last one
```

**note**: using `--debug` shows the `solc` & `abigen` command lines used during execution.

## build output

The build output goes into 2 sub-folders:

* `dist`: contains `ouptut_folder`(provided config file) with sub folders for each version with dedicated outputs of `solc`
    * `abi`: abi and signatures of contracts
    * `bin`: binaries of contracts
    * `combined-json`: contains one single file `combined.json` used as input of `abigen`
* `${ouput_folder}/${output_folder}_go`: for each version contains bindings for the `go` language.

Shortened view of the `dist` folder:

contains `commerce` and `token_permit` abi, bin and combined-json.
```
dist
├── commerce
│   └── v0.0.1
│       ├── abi
│       │   ├── Address.abi
│       │   ├── Address.signatures
...
│       │   ├── Payment.signatures
│       │   ├── PaymentSplitter.abi
│       │   ├── PaymentSplitter.signatures
│       │   └── SafeERC20.signatures
│       ├── bin
│       │   ├── Address.bin
...
│       │   ├── PaymentSplitter.bin
│       │   └── SafeERC20.bin
│       └── combined-json
│           └── combined.json
└── token
    └── v0.0.1
        ├── abi
        │   ├── Context.abi
        │   ├── Context.signatures
        ...
        │   ├── Ownable.signatures
        │   ├── SignedMath.abi
        │   ├── SignedMath.signatures
        │   ├── Strings.abi
        │   └── Strings.signatures
        ├── bin
        │   ├── Context.bin
        │   ├── Counters.bin
        ...
        │   ├── SignedMath.bin
        │   └── Strings.bin
        └── combined-json
            └── combined.json

```



View of `commerce/commerce_go` go binding:

```
commerce
└── commerce_go
    ├── doc.go
    ├── events
    │   └── event_info.go
    ├── events.go
    ├── go.mod
    ├── go.sum
    ├── unique_events_test.go
    └── v0.0.1
        └── contracts.go

```

## adding a new version

To add a new version, follow these step:

* add a new `version` in the configuration
* run `./build.sh -latest`

**Note:** The config file needs have all the versions of contracts, in order to keep the package handling events handle both old and new versions.

## tagging

Once done with all changes, the project needs to be tagged in order to be used from external projects.

Tagging must be done at 2 levels:

* project level
* sub-module level for go-bindings

Note: this can also be achieved through the github web interface.

For instance, sub-module level go-binding for `commerce` contracts :

```
  $ git tag v1.0.0
  $ git tag commerce/commerce_go/v1.0.0
  $ git push origin v1.0.0
  $ git push origin commerce/commerce_go/v1.0.0
```

The version can now be imported into an external project:

```
go get -d github.com/eluv-io/contracts-evm-builds/contracts-go@v1.0.0
```
