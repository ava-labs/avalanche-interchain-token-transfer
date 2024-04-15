// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {MockERC20SendAndCallReceiver} from "../src/mocks/MockERC20SendAndCallReceiver.sol";
import {MockNativeSendAndCallReceiver} from "../src/mocks/MockNativeSendAndCallReceiver.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";

contract MockERC20SendAndCallReceiverTest is Test {
    IERC20 public erc20;
    MockERC20SendAndCallReceiver public erc20SendAndCallReceiver;

    event TokensReceived(address token, uint256 amount, bytes payload);

    function setUp() public virtual {
        erc20 = new ExampleERC20();
        erc20SendAndCallReceiver = new MockERC20SendAndCallReceiver();
    }

    function testRevert() public {
        vm.expectRevert("MockERC20SendAndCallReceiver: empty payload");
        erc20SendAndCallReceiver.receiveTokens(address(erc20), 10, new bytes(0));
    }

    function testSuccess() public {
        uint256 amount = 100;
        bytes memory payload = hex"9876543210";
        erc20.approve(address(erc20SendAndCallReceiver), amount);
        vm.expectEmit(true, true, true, true, address(erc20SendAndCallReceiver));
        emit TokensReceived(address(erc20), amount, payload);
        erc20SendAndCallReceiver.receiveTokens(address(erc20), amount, payload);
        assertEq(erc20.balanceOf(address(erc20SendAndCallReceiver)), amount);
    }
}

contract MockNativeSendAndCallReceiverTest is Test {
    MockNativeSendAndCallReceiver public nativeSendAndCallReceiver;

    event TokensReceived(uint256 amount, bytes payload);

    function setUp() public virtual {
        nativeSendAndCallReceiver = new MockNativeSendAndCallReceiver();
    }

    function testRevert() public {
        vm.expectRevert("MockNativeSendAndCallReceiver: empty payload");
        nativeSendAndCallReceiver.receiveTokens(new bytes(0));
    }

    function testSuccess() public {
        uint256 amount = 10;
        bytes memory payload = hex"1234567890";
        vm.expectEmit(true, true, true, true, address(nativeSendAndCallReceiver));
        emit TokensReceived(amount, payload);
        nativeSendAndCallReceiver.receiveTokens{value: 10}(payload);
        assertEq(address(nativeSendAndCallReceiver).balance, amount);
    }
}