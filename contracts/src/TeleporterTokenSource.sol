// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITeleporterTokenBridge, SendTokensInput} from "./interfaces/ITeleporterTokenBridge.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title TeleporterTokenSource
 * @dev Abstract contract for a Teleporter token bridge that sends tokens to {TeleporterTokenDestination} instances.
 *
 * This contract also handles multihop transfers, where tokens sent from a {TeleporterTokenDestination}
 * instance are forwarded to another {TeleporterTokenDestination} instance.
 */
abstract contract TeleporterTokenSource is ITeleporterTokenBridge, TeleporterOwnerUpgradeable {
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The ERC20 token this contract uses to pay for Teleporter fees.
    IERC20 public immutable feeToken;

    /**
     * @notice Tracks the balances of tokens sent to other bridge instances.
     * Bridges are not allowed to unwrap more than has been sent to them.
     * @dev (destinationBlockchainID, destinationBridgeAddress) -> balance
     */
    mapping(
        bytes32 destinationBlockchainID
            => mapping(address destinationBridgeAddress => uint256 balance)
    ) public bridgedBalances;

    /**
     * @notice Initializes this source token bridge instance to send
     * tokens to the specified destination chain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeToken_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        feeToken = IERC20(feeToken_);
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
    }

    /**
     * @notice Sends tokens to the specified destination token bridge instance.
     *
     * @dev Increases the bridge balance sent to each destination token bridge instance,
     * and uses Teleporter to send a cross chain message.
     * TODO: Determine if this can be abstracted to a common function with {TeleporterTokenDestination}
     * Requirements:
     *
     * - `input.destinationBlockchainID` cannot be the same as the current blockchainID
     * - `input.destinationBridgeAddress` cannot be the zero address
     * - `input.recipient` cannot be the zero address
     * - `amount` must be greater than 0
     * - `amount` must be greater than `input.primaryFee`
     */
    function _send(SendTokensInput memory input, uint256 amount) internal virtual {
        require(
            input.destinationBlockchainID != blockchainID,
            "TeleporterTokenSource: cannot bridge to same chain"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TeleporterTokenSource: zero destination bridge address"
        );
        require(input.recipient != address(0), "TeleporterTokenSource: zero recipient address");
        require(amount > 0, "TeleporterTokenSource: zero send amount");
        require(
            amount > input.primaryFee, "TeleporterTokenSource: insufficient amount to cover fees"
        );

        // Subtract fee amount from amount and increase bridge balance
        amount -= input.primaryFee;
        bridgedBalances[input.destinationBlockchainID][input.destinationBridgeAddress] += amount;

        // send message to destinationBridgeAddress
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(feeToken),
                    amount: input.primaryFee
                }),
                // TODO: Set requiredGasLimit
                requiredGasLimit: 0,
                allowedRelayerAddresses: input.allowedRelayerAddresses,
                message: abi.encode(input.recipient, amount)
            })
        );

        emit SendTokens(messageID, msg.sender, amount);
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual override {
        (SendTokensInput memory input, uint256 amount) =
            abi.decode(message, (SendTokensInput, uint256));

        // Check that bridge instance returning has sufficient amount in balance
        uint256 senderBalance = bridgedBalances[sourceBlockchainID][originSenderAddress];
        require(senderBalance >= amount, "TeleporterTokenSource: insufficient bridge balance");

        // Decrement the bridge balance by the unwrap amount
        bridgedBalances[sourceBlockchainID][originSenderAddress] = senderBalance - amount;

        // decrement totalAmount from bridge balance
        if (input.destinationBlockchainID == blockchainID) {
            require(
                input.destinationBridgeAddress == address(this),
                "TeleporterTokenSource: invalid bridge address"
            );
            _withdraw(input.recipient, amount);
        }

        _send(input, amount);
    }

    // solhint-disable-next-line no-empty-blocks
    function _withdraw(address recipient, uint256 amount) internal virtual;
}
