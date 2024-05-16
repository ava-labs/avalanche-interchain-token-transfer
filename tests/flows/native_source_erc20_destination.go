package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TeleporterTokenDestination/ERC20Destination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TeleporterTokenSource/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a native token source on the primary network
 * Deploys ERC20Destination to Subnet A
 * Bridges C-Chain native tokens to Subnet A
 * Bridge back tokens from Subnet A to C-Chain
 */
func NativeSourceERC20Destination(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create a NativeTokenSource for bridging the native token
	nativeTokenSourceAddress, nativeTokenSource := utils.DeployNativeTokenSource(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := wavax.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavax.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavax.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination to Subnet A
	erc20DestinationAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20DestinationOnSource(
		ctx,
		network,
		cChainInfo,
		nativeTokenSourceAddress,
		subnetAInfo,
		erc20DestinationAddress,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A
	input := nativetokensource.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultERC20RequiredGas,
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, bridgedAmount := utils.SendNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		wavax,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to Subnet A and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		erc20Destination,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20Destination.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	input_A := erc20destination.SendTokensInput{
		DestinationBlockchainID:  cChainInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSourceAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   erc20DestinationAddress,
		PrimaryFee:               big.NewInt(1e10),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	// Send tokens on Subnet A back for native tokens on C-Chain
	receipt, bridgedAmount = utils.SendERC20Destination(
		ctx,
		subnetAInfo,
		erc20Destination,
		erc20DestinationAddress,
		input_A,
		teleporterUtils.BigIntSub(bridgedAmount, input_A.PrimaryFee),
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	// Check that the recipient received the tokens
	utils.CheckNativeTokenSourceWithdrawal(
		ctx,
		nativeTokenSourceAddress,
		wavax,
		receipt,
		bridgedAmount,
	)

	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, cChainInfo.RPCClient)
}
