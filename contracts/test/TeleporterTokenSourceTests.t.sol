// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {TeleporterTokenSource, IWarpMessenger} from "../src/TeleporterTokenSource.sol";

contract ExampleSourceApp is TeleporterTokenSource {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeTokenAddress
    ) TeleporterTokenSource(teleporterRegistryAddress, teleporterManager, feeTokenAddress) {}

    function _withdraw(address recipient, uint256 amount) internal virtual override {}
}

contract TeleporterTokenSourceTest is Test {
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    ExampleSourceApp public app;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_SOURCE_BLOCKCHAIN_ID)
        );

        vm.deployContract(
            "ExampleSourceApp",
            "ExampleSourceApp(address,address,address)",
            address(teleporterRegistry),
            address(this),
            address(feeToken)
        );
    }
}
