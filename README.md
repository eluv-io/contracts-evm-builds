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
    solc: "0.5.4" # version of the solidity compiler to user
  -
    tag: 'v1.0.0'
    solc: '0.8.13'
  -
    tag: "v1.2.0"  # only on last version tag may reflect the future tag when 'use' is a commit hash 
    solc: "0.8.13"
    use: "5f7a39e46d57232da5db9df5104de8c877544504" # only last version may use a commit hash
```

## build script

The build script makes builds for two contracts projects - 'new contracts' and 'legacy' - in `github.com/eluv-io/contracts-evm`.

The corresponding main files are:

* new-contracts: `src/tenant/Tenant.sol`
* legacy: `src/legacy/main.sol`

```
usage: ./build.sh -all|-latest [--debug]

Build contracts versions defined in 'config.yaml'
-all:    build all of contracts versions.
-latest: build the last one
```

**note**: using `--debug` shows the `solc` & `abigen` command lines used during execution.

## build output

The build output goes into 2 sub-folders:

* `dist`: for each version contains sub folders with dedicated outputs of `solc`
    * `abi`: abi and signatures of contracts
    * `bin`: binaries of contracts
    * `combined-json`: contains one single file `combined.json` used as input of `abigen`
* `contracts-go`: for each version contains bindings for the `go` language.

Shortened view of the `dist` folder:

```
dist/
├── v0.0.1
│   ├── contracts
│   │   ├── abi
...
│   │   │   ├── Tenant.abi
│   │   │   └── Tenant.signatures
│   │   ├── bin
...
│   │   │   └── Tenant.bin
│   │   └── combined-json
│   │       └── combined.json
│   └── legacy
│       ├── abi
...
│       │   ├── UserSpace.abi
│       │   └── UserSpace.signatures
│       ├── bin
...
│       │   └── UserSpace.bin
│       └── combined-json
│           └── combined.json
├── v1.0.0
│   ├── contracts
│   └── legacy
└── v1.2.0
    ├── contracts
    └── legacy
```

View of `contracts-go`:

```
contracts-go/
├── doc.go
├── events
│   ├── event_info.go
│   └── event_info_test.gox
├── events.go
├── go.mod
├── go.sum
├── unique_events_test.go
├── v0.0.1
│   ├── contracts
│   │   └── contracts.go
│   └── legacy
│       └── contracts.go
├── v1.0.0
│   ├── contracts
│   │   └── contracts.go
│   ├── contracts_test
│   │   └── contracts_test.go
│   └── legacy
│       └── contracts.go
└── v1.2.0
    ├── contracts
    │   └── contracts.go
    └── legacy
        └── contracts.go
```

## adding a new version

To add a new version, follow these step:

* add a new `version` in the configuration
* run `./build.sh -latest`
* edit file `contracts-go/events.go` and add the new bindings to the `allEventInfos` variable, as follows:

```
	allEventInfos = []packageEvent{
		{v100.UniqueEvents, "100"},
	}
```

## tagging

Once done with all changes, the project needs to be tagged in order to be used from external projects.

Tagging must be done at 2 levels:

* project level
* sub-module level for go-bindings

Note: this can also be achieved through the github web interface.

```
  $ git tag v1.0.0
  $ git tag contracts-go/v1.0.0
  $ git push origin v1.0.0
  $ git push origin contracts-go/v1.0.0
```

The version can now be imported into an external project:

```
go get -d github.com/eluv-io/contracts-evm-builds/contracts-go@v1.0.0
```
