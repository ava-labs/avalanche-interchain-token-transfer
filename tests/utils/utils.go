// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	nativetokendestination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	examplewavax "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/ExampleWAVAX"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

// Deployer keys set in the genesis file in order to determine the deployed address in advance.
// The deployed address is set as an admin for the Native Minter precompile.

var nativeTokenDestinationDeployerKeys = []string{
	// Deployer address: 			   0x1337cfd2dCff6270615B90938aCB1efE79801704
	// NativeTokenDestination address: 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B
	"aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985",

	// Deployer address: 			   0xFcec6c0674037f99fa473de09609B4b6D8158863
	// NativeTokenDestination address: 0x962c62B01529ecc0561D85d3fe395921ddC3665B
	"81e5e98c89023dabbe43e1081314eaae174330aae6b44c9d1371b6c0bb7ae74a",

	// Deployer address:			   0x2e1533d976A675bCD6306deC3B05e9f73e6722Fb
	// NativeTokenDestination address: 0x1549B96D9D97F435CA9b25000FEDE3A7e54C0bb9
	"5ded9cacaca7b88d6a3dc24641cfe41ef00186f98e7fa65135eac50fd5977f7a",

	// Deployer address:			   0xA638b0a597dc0520e2f20E83cFbeBBCd45a79990
	// NativeTokenDestination address: 0x190110D1228EB2cDd36559b2215A572Dc8592C3d
	"a6c530cb407778d10e1f70be6624aa57d0c724f6f9cb585e9744052d7f48ba19",

	// Deployer address:			   0x787C079cB0d5A7AA1Cae95d991F76Dce771A432D
	// NativeTokenDestination address: 0xf9EF017A764F265A1fD0975bfc200725E41d860E
	"e95fa6fd1d2a6b02890b75062bed583ce6256c5b473b3323b93ac4cbf20dbe7a",

	// Deployer address:			   0x741D536f5B07bcD43727CD8435389CA36aE5A4Ae
	// NativeTokenDestination address: 0x4f3663be6d22B0F19F8617f1A9E9485aB0144Bff
	"8a92f3f468ce5b0d99f9aaa55695f93e03dbbb6d5e3faba80f92a7876be740d6",
}
var nativeTokenDestinationDeployerKeyIndex = 0

func DeployERC20Source(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenSourceAddress common.Address,
) (common.Address, *erc20source.ERC20Source) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Source, err := erc20source.DeployERC20Source(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenSourceAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Source
}

func DeployERC20Destination(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	sourceBlockchainID ids.ID,
	tokenSourceAddress common.Address,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,
) (common.Address, *erc20destination.ERC20Destination) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Destination, err := erc20destination.DeployERC20Destination(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		sourceBlockchainID,
		tokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Destination
}

func DeployNativeTokenDestination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	sourceBlockchainID ids.ID,
	tokenSourceAddress common.Address,
	feeTokenAddress common.Address,
	initialReserveImbalance *big.Int,
	decimalsShift *big.Int,
	multiplyOnReceive bool,
	burnedFeesReportingRewardPercentage *big.Int,
) (common.Address, *nativetokendestination.NativeTokenDestination) {
	// The Native Token Destination needs a unique deployer key, whose nonce 0 is used to deploy the contract.
	// The resulting contract address has been added to the genesis file as an admin for the Native Minter precompile.
	Expect(nativeTokenDestinationDeployerKeyIndex).Should(BeNumerically("<", len(nativeTokenDestinationDeployerKeys)))
	deployerKeyStr := nativeTokenDestinationDeployerKeys[nativeTokenDestinationDeployerKeyIndex]
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	fmt.Println("Deployer Address: ", crypto.PubkeyToAddress(deployerPK.PublicKey))

	opts, err := bind.NewKeyedTransactorWithChainID(
		deployerPK,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())

	address, tx, nativeTokenDestination, err := nativetokendestination.DeployNativeTokenDestination(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		sourceBlockchainID,
		tokenSourceAddress,
		feeTokenAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
		burnedFeesReportingRewardPercentage,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	fmt.Println("Deployed NativeTokenDestination contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Increment to the next deployer key so that the next contract deployment succeeds
	nativeTokenDestinationDeployerKeyIndex++

	return address, nativeTokenDestination
}

func DeployNativeTokenSource(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	feeTokenAddress common.Address,
) (common.Address, *nativetokensource.NativeTokenSource) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, nativeTokenSource, err := nativetokensource.DeployNativeTokenSource(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		feeTokenAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, nativeTokenSource
}

func DeployExampleWAVAX(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *examplewavax.ExampleWAVAX) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy mock WAVAX contract
	address, tx, token, err := examplewavax.DeployExampleWAVAX(opts, subnet.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed ExampleWAVAX contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, token
}

func SendERC20Source(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	input erc20source.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20Source to spend the tokens
	teleporterUtils.ERC20Approve(
		ctx,
		sourceToken,
		erc20SourceAddress,
		amount,
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())
	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	Expect(event.Amount).Should(Equal(bridgedAmount))

	return receipt, event.Amount
}

func SendNativeTokenSource(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	input nativetokensource.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenSource.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())
	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSource.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	Expect(event.Amount).Should(Equal(bridgedAmount))

	return receipt, event.Amount
}

func SendNativeTokenDestination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	input nativetokendestination.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
	tokenMultiplier *big.Int,
	multiplyOnReceive bool,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenDestination.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())
	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)
	var scaledBridgedAmount *big.Int
	if multiplyOnReceive {
		scaledBridgedAmount = big.NewInt(0).Div(bridgedAmount, tokenMultiplier)
	} else {
		scaledBridgedAmount = teleporterUtils.BigIntMul(bridgedAmount, tokenMultiplier)
	}

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenDestination.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, scaledBridgedAmount)

	return receipt, event.Amount
}

func SendERC20Destination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Destination *erc20destination.ERC20Destination,
	erc20DestinationAddress common.Address,
	input erc20destination.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Destination.Approve(
		opts,
		erc20DestinationAddress,
		amount,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20Destination.Send(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	Expect(event.Amount).Should(Equal(bridgedAmount))

	return receipt, event.Amount
}

// Send a native token from fromBridge to toBridge via multihop through the C-Chain
// Requires that both fromBridge and toBridge are fully collateralized
// Requires that both fromBridge and toBridge have the same tokenMultiplier and multiplyOnReceive
// with respect to the original asset on the C-Chain
func SendNativeMultihopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *nativetokendestination.NativeTokenDestination,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *nativetokendestination.NativeTokenDestination,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	bridgedAmount *big.Int,
	tokenMultiplier *big.Int,
	multiplyOnReceive bool,
) {
	input := nativetokendestination.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultNativeTokenRequiredGasLimit,
	}
	// Find the amount sent by fromBridge. This is before any scaling/unscaling is applied.
	bridgedAmount = big.NewInt(0).Sub(bridgedAmount, input.PrimaryFee)

	// Send tokens through a multihop transfer
	originReceipt, _ := SendNativeTokenDestination(
		ctx,
		fromSubnet,
		fromBridge,
		input,
		bridgedAmount,
		sendingKey,
		tokenMultiplier,
		multiplyOnReceive,
	)

	// Relay the first message back to the home-chain, in this case C-Chain,
	// which then performs the multihop transfer to the destination chain
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	// When we relay the above message to the home-chain, a multihop transfer
	// is performed to the destination chain. Parse for the send tokens event
	// and relay to final destination.
	destinationReceipt := network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)

	CheckNativeTokenDestinationMint(
		ctx,
		toBridge,
		recipientAddress,
		destinationReceipt,
		bridgedAmount,
	)
}

func SendERC20MultihopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *erc20destination.ERC20Destination,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *erc20destination.ERC20Destination,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	bridgedAmount *big.Int,
) {
	teleporterUtils.SendNativeTransfer(
		ctx,
		fromSubnet,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	input := erc20destination.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultERC20RequiredGasLimit,
	}

	// Send tokens through a multihop transfer
	originReceipt, bridgedAmount := SendERC20Destination(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		bridgedAmount,
		sendingKey,
	)

	// Relay the first message back to the home-chain, in this case C-Chain,
	// which then performs the multihop transfer to the destination chain
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	// When we relay the above message to the home-chain, a multihop transfer
	// is performed to the destination chain. Parse for the send tokens event
	// and relay to final destination.
	destinationReceipt := network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)

	CheckERC20DestinationWithdrawal(
		ctx,
		toBridge,
		destinationReceipt,
		recipientAddress,
		bridgedAmount,
	)

	balance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}

func CheckERC20SourceWithdrawal(
	ctx context.Context,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	sourceTransferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceToken.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(sourceTransferEvent.From).Should(Equal(erc20SourceAddress))
	Expect(sourceTransferEvent.To).Should(Equal(expectedRecipientAddress))
	Expect(sourceTransferEvent.Value).Should(Equal(expectedAmount))
}

func CheckERC20DestinationWithdrawal(
	ctx context.Context,
	erc20Destination *erc20destination.ERC20Destination,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	transferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(transferEvent.From).Should(Equal(common.Address{}))
	Expect(transferEvent.To).Should(Equal(expectedRecipientAddress))
	Expect(transferEvent.Value).Should(Equal(expectedAmount))
}

func CheckNativeTokenSourceWithdrawal(
	ctx context.Context,
	nativeTokenSourceAddress common.Address,
	sourceToken *examplewavax.ExampleWAVAX,
	receipt *types.Receipt,
	expectedAmount *big.Int,
) {
	withdrawalEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceToken.ParseWithdrawal)
	Expect(err).Should(BeNil())
	Expect(withdrawalEvent.Sender).Should(Equal(nativeTokenSourceAddress))
	Expect(withdrawalEvent.Amount).Should(Equal(expectedAmount))
}

func CheckNativeTokenDestinationMint(
	ctx context.Context,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	recipient common.Address,
	receipt *types.Receipt,
	expectedMint *big.Int,
) {
	mintEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenDestination.ParseNativeTokensMinted)
	Expect(err).Should(BeNil())
	Expect(mintEvent.Recipient).Should(Equal(recipient))
	teleporterUtils.ExpectBigEqual(mintEvent.Amount, expectedMint)
}

func CheckNativeTokenDestinationCollateralize(
	ctx context.Context,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	receipt *types.Receipt,
	expectedAmount *big.Int,
	expectedRemaining *big.Int,
) {
	collateralEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenDestination.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	fmt.Println("log ", collateralEvent)
	teleporterUtils.ExpectBigEqual(collateralEvent.Amount, expectedAmount)
	teleporterUtils.ExpectBigEqual(collateralEvent.Remaining, expectedRemaining)
}
