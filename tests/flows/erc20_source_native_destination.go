package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	nativetokendestination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenDestination"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token source on the primary network
 * Deploys NativeDestination to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A as Subnet A's native token
 * Bridge tokens from Subnet A to Subnet B through multihop
 * Bridge back tokens from Subnet B to Subnet A through multihop
 */
func ERC20SourceNativeDestination(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the source token to be bridged
	sourceTokenAddress, sourceToken := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20Source for bridging the source token
	erc20SourceAddress, erc20Source := utils.DeployERC20Source(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		sourceTokenAddress,
	)

	// Deploy an example WAVAX on Subnet A
	wavaxAddressA, _ := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Deploy a NativeTokenDestination to Subnet A
	nativeTokenDestinationAddressA, nativeTokenDestinationA := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		wavaxAddressA,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to Subnet A
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressA,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
	}
	// Bridge the initial imbalance, which is scaled up on the destination by 1 decimal place
	// This will mint 9/10 the initial imbalance amount on the destination (after fees)
	receipt, bridgedAmount := utils.SendERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
		input,
		initialReserveImbalance,
		fundedKey,
	)

	// Amount received by the destination is bridgedAmount * 10 - initialReserveImbalance
	receivedAmount := big.NewInt(0).Sub(big.NewInt(0).Mul(bridgedAmount, big.NewInt(10)), initialReserveImbalance)

	// Relay the message to subnet A and check for a native token mint withdrawal
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckNativeTokenDestinationMint(
		ctx,
		nativeTokenDestinationA,
		recipientAddress,
		receipt,
		receivedAmount,
	)
	teleporterUtils.CheckBalance(
		ctx,
		recipientAddress,
		receivedAmount,
		subnetAInfo.RPCClient,
	)

	// Verify the recipient received the tokens
	teleporterUtils.CheckBalance(ctx, recipientAddress, receivedAmount, subnetAInfo.RPCClient)

	// Send back to the source chain and check that ERC20Source received the tokens
	input_A := nativetokendestination.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: erc20SourceAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
	}
	// Send half of the received amount to account for gas expenses
	amountToSend := big.NewInt(0).Div(receivedAmount, big.NewInt(2))
	receipt, bridgedAmount = utils.SendNativeTokenDestination(
		ctx,
		subnetAInfo,
		nativeTokenDestinationA,
		input_A,
		amountToSend,
		recipientKey,
		tokenMultiplier,
		multiplyOnReceive,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	// Check that the recipient received the tokens
	utils.CheckERC20SourceWithdrawal(
		ctx,
		erc20SourceAddress,
		sourceToken,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := sourceToken.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
