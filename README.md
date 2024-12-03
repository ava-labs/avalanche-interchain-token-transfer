> [!IMPORTANT]
> This repository has been merged into Interchain Messaging Contracts repository at [ICM Contracts](https://github.com/ava-labs/icm-contracts). This repository is archived and will not be updated.

# Avalanche Interchain Token Transfer (ICTT)

## Overview

Avalanche Interchain Token Transfer (ICTT) is an application that allows users to transfer tokens between Subnets. The implementation is a set of smart contracts that are deployed across multiple Subnets, and leverages [Teleporter](https://github.com/ava-labs/teleporter) for cross-chain communication.

Each token transferrer instance consists of one "home" contract and at least one but possibly many "remote" contracts. Each home contract instance manages one asset to be transferred out to `TokenRemote` instances. The home contract lives on the Subnet where the asset to be transferred exists. A transfer consists of locking the asset as collateral on the home Subnet and minting a representation of the asset on the remote Subnet. The remote contracts, each of which has a single specified home contract, live on other Subnets that want to import the asset transferred by their specified home. The token transferrers are designed to be permissionless: anyone can register compatible `TokenRemote` instances to allow for transferring tokens from the `TokenHome` instance to that new `TokenRemote` instance. The home contract keeps track of token balances transferred to each `TokenRemote` instance, and handles returning the original tokens back to the user when assets are transferred back to the `TokenHome` instance. `TokenRemote` instances are registered with their home contract via a Teleporter message upon creation.

Home contract instances specify the asset to be transferred as either an ERC20 token or the native token, and they allow for transferring the token to any registered `TokenRemote` instances. The token representation on the remote chain can also either be an ERC20 or native token, allowing users to have any combination of ERC20 and native tokens between home and remote chains:

- `ERC20` -> `ERC20`
- `ERC20` -> `Native`
- `Native` -> `ERC20`
- `Native` -> `Native`

The remote tokens are designed to have compatibility with the token transferrer on the home chain by default, and they allow custom logic to be implemented in addition. For example, developers can inherit and extend the `ERC20TokenRemote` contract to add additional functionality, such as a custom minting, burning, or transfer logic.

The token transferrer also supports "multi-hop" transfers, where tokens can be transferred between remote chains. To illustrate, consider two remotes _R<sub>a</sub>_ and _R<sub>b</sub>_ that are both connected to the same home _H_. A multi-hop transfer from _R<sub>a</sub>_ to _R<sub>b</sub>_ first gets routed from _R<sub>a</sub>_ to _H_, where remote balances are updated, and then _H_ automatically routes the transfer on to _R<sub>b</sub>_.

In addition to supporting basic token transfers, the token transferrer contracts offer a `sendAndCall` interface for transferring tokens and using them in a smart contract interaction all within a single Teleporter message. If the call to the recipient smart contract fails, the transferred tokens are sent to a fallback recipient address on the destination chain of the transfer. The `sendAndCall` interface enables the direct use of transferred tokens in dApps on other chains, such as performing swaps, using the tokens to pay for fees when invoking services, etc.

A breakdown of the structure of the contracts that implement this function can be found under `./contracts` [here](./contracts/README.md).

## Audits

Some contracts in this repository have been audited. The `main` branch may contain unaudited code. Please check [here](./audits/README.md) for which versions of each contract have been audited.
DO NOT USE UN-AUDITED CODE IN PRODUCTION!

## Upgradeability

The avalanche-interchain-token-transfer contracts are non-upgradeable and cannot be changed once it is deployed. This provides immutability to the contracts, and ensures that the contract's behavior at each address is unchanging.

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests

## Structure

- `contracts/` is a Foundry project that includes the implementation of the token transferrer contracts and Solidity unit tests
- `scripts/` includes various bash utility scripts
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.

## Solidity Unit Tests

Unit tests are written under `contracts/test/` and can be run with `forge`:

```
cd contracts
forge test -vvv
```

Unit test coverage of the contracts can be viewed using `forge coverage`:

```
$ forge coverage
[⠢] Compiling...
[⠒] Compiling 78 files with 0.8.25
[⠆] Solc 0.8.25 finished in 3.92s
Compiler run successful!
Analysing contracts...
Running tests...
| File                                             | % Lines           | % Statements      | % Branches      | % Funcs           |
| ------------------------------------------------ | ----------------- | ----------------- | --------------- | ----------------- |
| src/TokenHome/ERC20TokenHome.sol                 | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/TokenHome/ERC20TokenHomeUpgradeable.sol      | 100.00% (27/27)   | 100.00% (33/33)   | 100.00% (6/6)   | 100.00% (11/11)   |
| src/TokenHome/NativeTokenHome.sol                | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/TokenHome/NativeTokenHomeUpgradeable.sol     | 100.00% (24/24)   | 100.00% (29/29)   | 100.00% (4/4)   | 100.00% (11/11)   |
| src/TokenHome/TokenHome.sol                      | 100.00% (158/158) | 100.00% (198/198) | 100.00% (26/26) | 100.00% (22/22)   |
| src/TokenRemote/ERC20TokenRemote.sol             | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/TokenRemote/ERC20TokenRemoteUpgradeable.sol  | 100.00% (36/36)   | 100.00% (40/40)   | 100.00% (10/10) | 100.00% (12/12)   |
| src/TokenRemote/NativeTokenRemote.sol            | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/TokenRemote/NativeTokenRemoteUpgradeable.sol | 100.00% (60/60)   | 100.00% (74/74)   | 100.00% (10/10) | 100.00% (19/19)   |
| src/TokenRemote/TokenRemote.sol                  | 100.00% (118/118) | 100.00% (155/155) | 100.00% (12/12) | 100.00% (24/24)   |
| src/WrappedNativeToken.sol                       | 100.00% (6/6)     | 100.00% (6/6)     | 100.00% (0/0)   | 100.00% (4/4)     |
| src/mocks/ExampleERC20Decimals.sol               | 100.00% (2/2)     | 100.00% (2/2)     | 100.00% (0/0)   | 100.00% (2/2)     |
| src/mocks/MockERC20SendAndCallReceiver.sol       | 100.00% (5/5)     | 100.00% (5/5)     | 100.00% (0/0)   | 100.00% (2/2)     |
| src/mocks/MockNativeSendAndCallReceiver.sol      | 100.00% (4/4)     | 100.00% (4/4)     | 100.00% (0/0)   | 100.00% (2/2)     |
| src/utils/CallUtils.sol                          | 100.00% (8/8)     | 100.00% (9/9)     | 100.00% (2/2)   | 100.00% (2/2)     |
| src/utils/SafeERC20TransferFrom.sol              | 100.00% (5/5)     | 100.00% (8/8)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/utils/SafeWrappedNativeTokenDeposit.sol      | 100.00% (5/5)     | 100.00% (8/8)     | 100.00% (0/0)   | 100.00% (1/1)     |
| src/utils/SendReentrancyGuard.sol                | 100.00% (8/8)     | 100.00% (10/10)   | 100.00% (0/0)   | 100.00% (4/4)     |
| src/utils/TokenScalingUtils.sol                  | 100.00% (8/8)     | 100.00% (14/14)   | 100.00% (2/2)   | 100.00% (4/4)     |
| Total                                            | 100.00% (478/478) | 100.00% (599/599) | 100.00% (72/72) | 100.00% (125/125) |
```

## E2E tests

End-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Transfer an ERC20 token between two Subnets` test:

```bash
GINKGO_FOCUS="Transfer an ERC20 token between two Subnets" ./scripts/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="Transfer an ERC20 token" ./scripts/e2e_test.sh
```

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all `ERC20TokenHome` E2E tests:

```go
	ginkgo.It("Transfer an ERC20 token between two Subnets",
		ginkgo.Label(erc20TokenHomeLabel, erc20TokenRemoteLabel),
		func() {
			flows.ERC20TokenHomeERC20TokenRemote(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="ERC20TokenHome" ./scripts/e2e_test.sh
```
