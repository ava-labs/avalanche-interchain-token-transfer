// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokendestination

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MultiHopInput is an auto generated low-level Go binding around an user-defined struct.
type MultiHopInput struct {
	SecondaryFee    *big.Int
	FallbackAddress common.Address
}

// NativeTokenDestinationSettings is an auto generated low-level Go binding around an user-defined struct.
type NativeTokenDestinationSettings struct {
	NativeAssetSymbol                   string
	TeleporterRegistryAddress           common.Address
	TeleporterManager                   common.Address
	SourceBlockchainID                  [32]byte
	TokenSourceAddress                  common.Address
	InitialReserveImbalance             *big.Int
	DecimalsShift                       uint8
	MultiplyOnDestination               bool
	BurnedFeesReportingRewardPercentage *big.Int
}

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	Input             SendTokensInput
	RecipientPayload  []byte
	RecipientGasLimit *big.Int
	FallbackRecipient common.Address
}

// SendAndCallMultiHopInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallMultiHopInput struct {
	SendAndCallInput SendAndCallInput
	MultiHopInput    MultiHopInput
}

// SendMultiHopInput is an auto generated low-level Go binding around an user-defined struct.
type SendMultiHopInput struct {
	Input         SendTokensInput
	MultiHopInput MultiHopInput
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFeeTokenAddress   common.Address
	PrimaryFee               *big.Int
	RequiredGasLimit         *big.Int
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nativeAssetSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenSourceAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimalsShift\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnDestination\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"burnedFeesReportingRewardPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structNativeTokenDestinationSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesBurned\",\"type\":\"uint256\"}],\"name\":\"ReportBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BURNED_FOR_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_MINTER\",\"outputs\":[{\"internalType\":\"contractINativeMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_DESTINATION_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE_CHAIN_BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnedFeesReportingRewardPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastestBurnedFeesReported\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnDestination\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"name\":\"reportBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"sendAndCallInput\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackAddress\",\"type\":\"address\"}],\"internalType\":\"structMultiHopInput\",\"name\":\"multiHopInput\",\"type\":\"tuple\"}],\"internalType\":\"structSendAndCallMultiHopInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCallMultiHop\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fallbackAddress\",\"type\":\"address\"}],\"internalType\":\"structMultiHopInput\",\"name\":\"multiHopInput\",\"type\":\"tuple\"}],\"internalType\":\"structSendMultiHopInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendMulithop\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNativeAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620055ef380380620055ef8339810160408190526200003591620007a6565b80602001518160400151826060015183608001518460a001518560c001518660e001518686818a600001516040516020016200007291906200089b565b60408051601f198184030181529190528b5160036200009283826200095c565b506004620000a182826200095c565b50506001600555506001600160a01b0381166200012b5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000176573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200019c919062000a28565b60075550620001ab336200053e565b620001b68162000590565b505060016009819055507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000213573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000239919062000a28565b60a05284620002a05760405162461bcd60e51b81526020600482015260356024820152600080516020620055cf83398151915260448201527f20736f7572636520626c6f636b636861696e2049440000000000000000000000606482015260840162000122565b60a0518503620003285760405162461bcd60e51b815260206004820152604660248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a2063616e6e60448201527f6f74206465706c6f7920746f2073616d6520626c6f636b636861696e20617320606482015265736f7572636560d01b608482015260a40162000122565b6001600160a01b038416620003955760405162461bcd60e51b81526020600482015260356024820152600080516020620055cf83398151915260448201527f20746f6b656e20736f7572636520616464726573730000000000000000000000606482015260840162000122565b60128260ff161115620004055760405162461bcd60e51b815260206004820152603160248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e76616044820152701b1a5908191958da5b585b1cd4da1a599d607a1b606482015260840162000122565b60c08590526001600160a01b03841660e052610140839052600a805460ff191684151781556200043790839062000b57565b6101005215156101205250505060a08401516000039250620004c59150505760405162461bcd60e51b815260206004820152603660248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20696e6960448201527f7469616c207265736572766520696d62616c616e636500000000000000000000606482015260840162000122565b6064816101000151106200052f5760405162461bcd60e51b815260206004820152602a60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20696e76616c69642060448201526970657263656e7461676560b01b606482015260840162000122565b61010001516101605262000b6f565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6200059a6200060f565b6001600160a01b038116620006015760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000122565b6200060c816200053e565b50565b6008546001600160a01b031633146200066b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000122565b565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b0381118282101715620006a957620006a96200066d565b60405290565b60005b83811015620006cc578181015183820152602001620006b2565b50506000910152565b600082601f830112620006e757600080fd5b81516001600160401b03808211156200070457620007046200066d565b604051601f8301601f19908116603f011681019082821181831017156200072f576200072f6200066d565b816040528381528660208588010111156200074957600080fd5b6200075c846020830160208901620006af565b9695505050505050565b80516001600160a01b03811681146200077e57600080fd5b919050565b805160ff811681146200077e57600080fd5b805180151581146200077e57600080fd5b600060208284031215620007b957600080fd5b81516001600160401b0380821115620007d157600080fd5b908301906101208286031215620007e757600080fd5b620007f162000683565b8251828111156200080157600080fd5b6200080f87828601620006d5565b825250620008206020840162000766565b6020820152620008336040840162000766565b604082015260608301516060820152620008506080840162000766565b608082015260a083015160a08201526200086d60c0840162000783565b60c08201526200088060e0840162000795565b60e08201526101009283015192810192909252509392505050565b6702bb930b83832b2160c51b815260008251620008c0816008850160208701620006af565b9190910160080192915050565b600181811c90821680620008e257607f821691505b6020821081036200090357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200095757600081815260208120601f850160051c81016020861015620009325750805b601f850160051c820191505b8181101562000953578281556001016200093e565b5050505b505050565b81516001600160401b038111156200097857620009786200066d565b6200099081620009898454620008cd565b8462000909565b602080601f831160018114620009c85760008415620009af5750858301515b600019600386901b1c1916600185901b17855562000953565b600085815260208120601f198616915b82811015620009f957888601518255948401946001909101908401620009d8565b508582101562000a185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020828403121562000a3b57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600181815b8085111562000a9957816000190482111562000a7d5762000a7d62000a42565b8085161562000a8b57918102915b93841c939080029062000a5d565b509250929050565b60008262000ab25750600162000b51565b8162000ac15750600062000b51565b816001811462000ada576002811462000ae55762000b05565b600191505062000b51565b60ff84111562000af95762000af962000a42565b50506001821b62000b51565b5060208310610133831016604e8410600b841016171562000b2a575081810a62000b51565b62000b36838362000a58565b806000190482111562000b4d5762000b4d62000a42565b0290505b92915050565b600062000b6860ff84168362000aa1565b9392505050565b60805160a05160c05160e0516101005161012051610140516101605161491362000cbc6000396000818161052e0152610fb701526000818161067201528181610a6f0152610c1e0152600081816103f001528181610abb0152818161103e015281816131a301526131ee01526000818161077b01528181610a920152818161101d0152818161318201526131cd0152600081816108df01528181610b500152818161115f015281816122c7015281816125e50152818161284e015281816129e601528181612c600152612d7e01526000818161046301528181610b2501528181611139015281816122a1015281816125bf015281816128280152818161297001528181612bd60152612d580152600081816108360152818161243e0152818161272601526130da0152600081816103a4015281816113ad01528181611e600152612e4901526149136000f3fe6080604052600436106102b25760003560e01c806370a0823111610175578063c1f823c7116100dc578063d127dc9b11610095578063ecd4ed1b1161006f578063ecd4ed1b14610570578063f2fde38b1461088d578063f3f981d8146108ad578063f5ea0603146108cd576102c1565b8063d127dc9b14610824578063d2cc7a7014610858578063dd62ed3e1461086d576102c1565b8063c1f823c7146107b0578063c452165e146107c3578063c868efaa146107db578063cb6a6417146107fb578063d0e30db0146102c1578063d10a5b8c1461080e576102c1565b8063973142971161012e57806397314297146106da578063a2309ff814610713578063a457c2d714610729578063a9059cbb14610749578063ba3f5a1214610769578063be0c6b1c1461079d576102c1565b806370a0823114610615578063715018a61461064b5780638ac7dd20146106605780638be887c9146106945780638da5cb5b146106a757806395d89b41146106c5576102c1565b80632e1a7d4d1161021957806347a9a22c116101d257806347a9a22c1461057057806349e3284e1461058d578063525975e6146105a757806355538c8b146105be57806355c26725146105de5780635eb99514146105f5576102c1565b80632e1a7d4d146104a5578063313ce567146104c5578063329c3e12146104e157806339509351146104fc5780633a23dfe21461051c5780634511243e14610550576102c1565b80631a7f5bec1161026b5780631a7f5bec146103925780632001eff6146103de578063223668441461041257806323b872dd1461043157806329b7b3fd146104515780632b0d8f1814610485576102c1565b806306fdde03146102c9578063095ea7b3146102f4578063151637f61461032457806315beb59f1461034457806318160ddd146103685780631906529c1461037d576102c1565b366102c1576102bf610901565b005b6102bf610901565b3480156102d557600080fd5b506102de610942565b6040516102eb9190613c36565b60405180910390f35b34801561030057600080fd5b5061031461030f366004613c69565b6109d4565b60405190151581526020016102eb565b34801561033057600080fd5b506102bf61033f366004613c95565b6109ee565b34801561035057600080fd5b5061035a61213481565b6040519081526020016102eb565b34801561037457600080fd5b5060025461035a565b34801561038957600080fd5b5061035a610bfc565b34801561039e57600080fd5b506103c67f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102eb565b3480156103ea57600080fd5b506103147f000000000000000000000000000000000000000000000000000000000000000081565b34801561041e57600080fd5b50600a5461031490610100900460ff1681565b34801561043d57600080fd5b5061031461044c366004613cad565b610c5d565b34801561045d57600080fd5b5061035a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561049157600080fd5b506102bf6104a0366004613cee565b610c81565b3480156104b157600080fd5b506102bf6104c0366004613d0b565b610d7a565b3480156104d157600080fd5b50604051601281526020016102eb565b3480156104ed57600080fd5b506103c66001600160991b0181565b34801561050857600080fd5b50610314610517366004613c69565b610dc6565b34801561052857600080fd5b5061035a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561055c57600080fd5b506102bf61056b366004613cee565b610de8565b34801561057c57600080fd5b506103c662010203600160981b0181565b34801561059957600080fd5b50600a546103149060ff1681565b3480156105b357600080fd5b5061035a6203a98081565b3480156105ca57600080fd5b506102bf6105d9366004613d0b565b610ee5565b3480156105ea57600080fd5b5061035a620249f081565b34801561060157600080fd5b506102bf610610366004613d0b565b611259565b34801561062157600080fd5b5061035a610630366004613cee565b6001600160a01b031660009081526020819052604090205490565b34801561065757600080fd5b506102bf61126a565b34801561066c57600080fd5b5061035a7f000000000000000000000000000000000000000000000000000000000000000081565b6102bf6106a2366004613d24565b61127c565b3480156106b357600080fd5b506008546001600160a01b03166103c6565b3480156106d157600080fd5b506102de6112a8565b3480156106e657600080fd5b506103146106f5366004613cee565b6001600160a01b031660009081526006602052604090205460ff1690565b34801561071f57600080fd5b5061035a600b5481565b34801561073557600080fd5b50610314610744366004613c69565b6112b7565b34801561075557600080fd5b50610314610764366004613c69565b611332565b34801561077557600080fd5b5061035a7f000000000000000000000000000000000000000000000000000000000000000081565b6102bf6107ab366004613d5f565b611340565b6102bf6107be366004613d72565b61136c565b3480156107cf57600080fd5b506103c6600160981b81565b3480156107e757600080fd5b506102bf6107f6366004613dae565b611398565b6102bf610809366004613e37565b61155c565b34801561081a57600080fd5b5061035a600c5481565b34801561083057600080fd5b5061035a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561086457600080fd5b5060075461035a565b34801561087957600080fd5b5061035a610888366004613e49565b611588565b34801561089957600080fd5b506102bf6108a8366004613cee565b6115b3565b3480156108b957600080fd5b5061035a6108c8366004613d0b565b611629565b3480156108d957600080fd5b506103c67f000000000000000000000000000000000000000000000000000000000000000081565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a26109403334611640565b565b60606003805461095190613e82565b80601f016020809104026020016040519081016040528092919081815260200182805461097d90613e82565b80156109ca5780601f1061099f576101008083540402835291602001916109ca565b820191906000526020600020905b8154815290600101906020018083116109ad57829003601f168201915b5050505050905090565b6000336109e2818585611700565b60019150505b92915050565b600a54610100900460ff1615610a625760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20616c726560448201526d18591e481c9959da5cdd195c995960921b60648201526084015b60405180910390fd5b60408051606080820183527f000000000000000000000000000000000000000000000000000000000000000082527f000000000000000000000000000000000000000000000000000000000000000060208084019182527f00000000000000000000000000000000000000000000000000000000000000001515848601908152855180870187526000815286518651818501529351848801529051151583850152855180840390940184526080909201855281810192909252835160c0810185527f000000000000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928101929092529192610bf691908101610b8f36879003870187613f2e565b8152620249f060208201526040016000604051908082528060200260200182016040528015610bc8578160200160208202803683370190505b50815260200183604051602001610bdf9190613f68565b604051602081830303815290604052815250611825565b50505050565b600080610c1862010203600160981b0131600160981b31613fc3565b905060007f0000000000000000000000000000000000000000000000000000000000000000600b54610c4a9190613fc3565b9050610c568282613fd6565b9250505090565b600033610c6b85828561194b565b610c768585856119bf565b506001949350505050565b610c89611b63565b6001600160a01b038116610caf5760405162461bcd60e51b8152600401610a5990613fe9565b6001600160a01b03811660009081526006602052604090205460ff1615610d2e5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610a59565b6001600160a01b038116600081815260066020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610db93382611b6b565b610dc33382611c9a565b50565b6000336109e2818585610dd98383611588565b610de39190613fc3565b611700565b610df0611b63565b6001600160a01b038116610e165760405162461bcd60e51b8152600401610a5990613fe9565b6001600160a01b03811660009081526006602052604090205460ff16610e905760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610a59565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600660205260409020805460ff19169055565b600160095414610f075760405162461bcd60e51b8152600401610a5990614037565b6002600955600c54600160981b31908111610f9c5760405162461bcd60e51b815260206004820152604960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a206275726e2061646460448201527f726573732062616c616e6365206e6f742067726561746572207468616e206c616064820152681cdd081c995c1bdc9d60ba1b608482015260a401610a59565b6000600c5482610fac9190613fd6565b905060006064610fdc7f00000000000000000000000000000000000000000000000000000000000000008461407b565b610fe69190614092565b90506000610ff48284613fd6565b600c859055905081156110165761100b3083611db3565b61101482611e35565b505b60006110637f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000084611e45565b116110d65760405162461bcd60e51b815260206004820152603960248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2073636160448201527f6c656420616d6f756e7420746f207265706f7274206275726e000000000000006064820152608401610a59565b604080518082018252600181528151808301835262010203600160981b0181526020808201859052925160009380840192611113929091016140b4565b6040516020818303038152906040528152509050600061120f6040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280306001600160a01b03168152602001888152508152602001898152602001600067ffffffffffffffff8111156111cf576111cf613ecc565b6040519080825280602002602001820160405280156111f8578160200160208202803683370190505b50815260200184604051602001610bdf9190613f68565b9050807f0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a58460405161124391815260200190565b60405180910390a2505060016009555050505050565b611261611b63565b610dc381611e5c565b611272611ffc565b6109406000612056565b600a5460ff1661129e5760405162461bcd60e51b8152600401610a59906140d4565b610dc381346120a8565b60606004805461095190613e82565b600033816112c58286611588565b9050838110156113255760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610a59565b610c768286868403611700565b6000336109e28185856119bf565b600a5460ff166113625760405162461bcd60e51b8152600401610a59906140d4565b610dc381346123a8565b600a5460ff1661138e5760405162461bcd60e51b8152600401610a59906140d4565b610dc381346126ac565b6113a0612915565b6007546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611417573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143b9190614128565b10156114a25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610a59565b6114ab336106f5565b156115115760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610a59565b611552848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061296e92505050565b610bf66001600555565b600a5460ff1661157e5760405162461bcd60e51b8152600401610a59906140d4565b610dc38134612ba2565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6115bb611ffc565b6001600160a01b0381166116205760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a59565b610dc381612056565b6000600561163883601f613fc3565b901c92915050565b6001600160a01b0382166116965760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610a59565b80600260008282546116a89190613fc3565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6001600160a01b0383166117625760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610a59565b6001600160a01b0382166117c35760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610a59565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b600080611830612e44565b604084015160200151909150156118d5576040830151516001600160a01b03166118b25760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610a59565b6040830151602081015190516118d5916001600160a01b03909116908390612f58565b604051630624488560e41b81526001600160a01b03821690636244885090611901908690600401614141565b6020604051808303816000875af1158015611920573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119449190614128565b9392505050565b60006119578484611588565b90506000198114610bf657818110156119b25760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610a59565b610bf68484848403611700565b6001600160a01b038316611a235760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610a59565b6001600160a01b038216611a855760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610a59565b6001600160a01b03831660009081526020819052604090205481811015611afd5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610a59565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610bf6565b610940611ffc565b6001600160a01b038216611bcb5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610a59565b6001600160a01b03821660009081526020819052604090205481811015611c3f5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610a59565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101611818565b505050565b80471015611cea5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610a59565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611d37576040519150601f19603f3d011682016040523d82523d6000602084013e611d3c565b606091505b5050905080611c955760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610a59565b80600b6000828254611dc59190613fc3565b90915550506040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba90604401600060405180830381600087803b158015611e1957600080fd5b505af1158015611e2d573d6000803e3d6000fd5b505050505050565b6000611e413083611640565b5090565b6000611e54848484600061303d565b949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ebc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ee09190614128565b60075490915081831115611f505760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610a59565b808311611fc55760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610a59565b6007839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6008546001600160a01b031633146109405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a59565b600880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600954146120ca5760405162461bcd60e51b8152600401610a5990614037565b60026009556120d88261306e565b366120e383806141fa565b90506020830181600061210c856121006080850160608601613cee565b60808501358635613134565b60408051808201909152919650915060009080600481526020016040518061016001604052806121393390565b6001600160a01b03168152602001866000013581526020018660200160208101906121649190613cee565b6001600160a01b031681526020016121826060880160408901613cee565b6001600160a01b03168152602081018a90526040016121a460c08a018a61421b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060e089013560208201526040016121fa6101208a016101008b01613cee565b6001600160a01b031681526020018660a0013581526020018760200160208101906122259190613cee565b6001600160a01b031681528735602091820152604051612246929101614269565b60408051601f1981840301815291905290529050600061213461227661226f60c089018961421b565b9050611629565b612280919061407b565b61228d906203a980613fc3565b905060006123526040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808960600160208101906123129190613cee565b6001600160a01b031681526020908101899052908252818101869052604080516000815280830182528184015251606090920191610bdf91889101613f68565b9050336001600160a01b0316817fb3209fd9e86173715f7081e24de11ecce66057a53b36d0733f6574c55e390c46898b6040516123909291906143aa565b60405180910390a35050600160095550505050505050565b6001600954146123ca5760405162461bcd60e51b8152600401610a5990614037565b60026009558160c081016123dd82613292565b60006123ef6040830160208401613cee565b6001600160a01b0316036124155760405162461bcd60e51b8152600401610a5990614464565b60006124378461242b6080860160608701613cee565b60808601358535613134565b90945090507f000000000000000000000000000000000000000000000000000000000000000083350361249b57306124756040850160208601613cee565b6001600160a01b03160361249b5760405162461bcd60e51b8152600401610a59906144a5565b60408051808201825260038152815160e0810183528535815260009260208084019291828201916124d0918a01908a01613cee565b6001600160a01b031681526020016124ee6060890160408a01613cee565b6001600160a01b03168152602001888152602001866000013581526020018760a0013581526020018660200160208101906125299190613cee565b6001600160a01b031690526040516125999190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b604051602081830303815290604052815250905060006126586040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808960600160208101906126309190613cee565b6001600160a01b0316815260209081018890529082526203a9809082015260400160006111cf565b9050336001600160a01b0316817f8a4d0e8ef2f385ab7df8c78e89773da8ae3679f1258fd71823e8c39e552a956487896040516126969291906144f0565b60405180910390a3505060016009555050505050565b6001600954146126ce5760405162461bcd60e51b8152600401610a5990614037565b60026009556126dc8261343f565b8160006126ff836126f36080850160608601613cee565b84608001356000613134565b60408051808201909152919450915060009080600281526020016040518060e001604052807f000000000000000000000000000000000000000000000000000000000000000081526020016127513390565b6001600160a01b0316815260200161276f6060880160408901613cee565b6001600160a01b031681526020810188905260400161279160c08a018a61421b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060e089013560208201526040016127e76101208a016101008b01613cee565b6001600160a01b03169052604051612802919060200161450b565b604051602081830303815290604052815250905060006128c26040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052808860600160208101906128999190613cee565b6001600160a01b03168152602090810188905290825260a08801359082015260400160006111cf565b9050336001600160a01b0316817fb3209fd9e86173715f7081e24de11ecce66057a53b36d0733f6574c55e390c4688886040516129009291906143aa565b60405180910390a35050600160095550505050565b6002600554036129675760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610a59565b6002600555565b7f000000000000000000000000000000000000000000000000000000000000000083146129e45760405162461bcd60e51b815260206004820152603060248201526000805160206148be83398151915260448201526f3634b21039b7bab931b29031b430b4b760811b6064820152608401610a59565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614612a795760405162461bcd60e51b815260206004820152603860248201526000805160206148be83398151915260448201527f6c696420746f6b656e20736f75726365206164647265737300000000000000006064820152608401610a59565b600081806020019051810190612a8f9190614609565b600a54909150610100900460ff161580612aac5750600a5460ff16155b15612ac157600a805461ffff19166101011790555b600181516004811115612ad657612ad6613eb6565b03612b0f5760008160200151806020019051810190612af59190614698565b9050612b0981600001518260200151613590565b50610bf6565b600281516004811115612b2457612b24613eb6565b03612b535760008160200151806020019051810190612b4391906146d2565b9050612b098182606001516135dd565b60405162461bcd60e51b815260206004820152603060248201526000805160206148be83398151915260448201526f6c6964206d657373616765207479706560801b6064820152608401610a59565b600160095414612bc45760405162461bcd60e51b8152600401610a5990614037565b6002600955612bd282613292565b81357f000000000000000000000000000000000000000000000000000000000000000014612c565760405162461bcd60e51b815260206004820152603d60248201526000805160206148be83398151915260448201527f6c69642064657374696e6174696f6e20626c6f636b636861696e2049440000006064820152608401610a59565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016612c906040840160208501613cee565b6001600160a01b031614612cb65760405162461bcd60e51b8152600401610a59906144a5565b6000612cd882612ccc6080860160608701613cee565b85608001356000613134565b60408051808201909152919350915060009080600181526020016040518060400160405280876040016020810190612d109190613cee565b6001600160a01b0316815260200186815250604051602001612d3291906140b4565b60405160208183030381529060405281525090506000612df26040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016040518060400160405280896060016020810190612dc99190613cee565b6001600160a01b03168152602090810188905290825260a08901359082015260400160006111cf565b9050336001600160a01b0316817f8a4d0e8ef2f385ab7df8c78e89773da8ae3679f1258fd71823e8c39e552a95648787604051612e309291906144f0565b60405180910390a350506001600955505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ea5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ec99190614790565b9050612eed816001600160a01b031660009081526006602052604090205460ff1690565b15612f535760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610a59565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015612fa9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fcd9190614128565b612fd79190613fc3565b6040516001600160a01b038516602482015260448101829052909150610bf690859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613703565b60008115158415150361305b57613054858461407b565b9050611e54565b6130658584614092565b95945050505050565b61308061307b82806141fa565b61343f565b600061308c82806141fa565b61309c90368190038101906147ad565b905060006130b06060840160408501613cee565b6001600160a01b0316036130d65760405162461bcd60e51b8152600401610a5990614464565b80517f000000000000000000000000000000000000000000000000000000000000000090036116fc57306001600160a01b031681602001516001600160a01b0316036116fc5760405162461bcd60e51b8152600401610a59906144a5565b60008061314086611e35565b9550831561317457306001600160a01b038616036131675761316184611e35565b50613174565b61317185856137d5565b93505b61317d86613934565b6131c87f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000085611e45565b6132137f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089611e45565b116132865760405162461bcd60e51b815260206004820152603b60248201527f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e737560448201527f6666696369656e7420746f6b656e7320746f207472616e7366657200000000006064820152608401610a59565b50939491935090915050565b60006132a46060830160408401613cee565b6001600160a01b0316036133035760405162461bcd60e51b8152602060048201526032602482015260008051602061489e83398151915260448201527120726563697069656e74206164647265737360701b6064820152608401610a59565b60008160a00135116133615760405162461bcd60e51b8152602060048201526033602482015260008051602061489e833981519152604482015272081c995c5d5a5c99590819d85cc81b1a5b5a5d606a1b6064820152608401610a59565b80356133c35760405162461bcd60e51b815260206004820152603a602482015260008051602061489e83398151915260448201527f2064657374696e6174696f6e20626c6f636b636861696e2049440000000000006064820152608401610a59565b60006133d56040830160208401613cee565b6001600160a01b031603610dc35760405162461bcd60e51b815260206004820152603b602482015260008051602061489e83398151915260448201527f2064657374696e6174696f6e20627269646765206164647265737300000000006064820152608401610a59565b61344881613292565b60008160e00135116134a75760405162461bcd60e51b8152602060048201526034602482015260008051602061489e833981519152604482015273081c9958da5c1a595b9d0819d85cc81b1a5b5a5d60621b6064820152608401610a59565b60a081013560e0820135106135125760405162461bcd60e51b815260206004820152603760248201526000805160206148be83398151915260448201527f6c696420726563697069656e7420676173206c696d69740000000000000000006064820152608401610a59565b600061352661012083016101008401613cee565b6001600160a01b031603610dc35760405162461bcd60e51b815260206004820152603b602482015260008051602061489e83398151915260448201527f2066616c6c6261636b20726563697069656e74206164647265737300000000006064820152608401610a59565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b826040516135cb91815260200190565b60405180910390a26116fc8282611db3565b6135e73082611db3565b600082600001518360200151846080015160405160240161360a9392919061483f565b60408051601f198184030181529181526020820180516001600160e01b0316630d57eddf60e21b17905260a0850151908501519192506000916136509190859085613951565b905080156136a45783604001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff48460405161369791815260200190565b60405180910390a2610bf6565b83604001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0846040516136e391815260200190565b60405180910390a260c0840151610bf6906001600160a01b031684611c9a565b6000613758826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613a269092919063ffffffff16565b805190915015611c9557808060200190518101906137769190614869565b611c955760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610a59565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa15801561381e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138429190614128565b90506138596001600160a01b038516333086613a35565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156138a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138c49190614128565b905081811161392a5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610a59565b6130658282613fd6565b61393e3082611b6b565b610dc362010203600160981b0182611c9a565b6000845a10156139a35760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e742067617300000000006044820152606401610a59565b834710156139f35760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c75650000006044820152606401610a59565b826001600160a01b03163b600003613a0d57506000611e54565b600080600084516020860188888bf19695505050505050565b6060611e548484600085613a6d565b6040516001600160a01b0380851660248301528316604482015260648101829052610bf69085906323b872dd60e01b90608401613006565b606082471015613ace5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610a59565b600080866001600160a01b03168587604051613aea919061488b565b60006040518083038185875af1925050503d8060008114613b27576040519150601f19603f3d011682016040523d82523d6000602084013e613b2c565b606091505b5091509150613b3d87838387613b48565b979650505050505050565b60608315613bb7578251600003613bb0576001600160a01b0385163b613bb05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a59565b5081611e54565b611e548383815115613bcc5781518083602001fd5b8060405162461bcd60e51b8152600401610a599190613c36565b60005b83811015613c01578181015183820152602001613be9565b50506000910152565b60008151808452613c22816020860160208601613be6565b601f01601f19169290920160200192915050565b6020815260006119446020830184613c0a565b6001600160a01b0381168114610dc357600080fd5b8035612f5381613c49565b60008060408385031215613c7c57600080fd5b8235613c8781613c49565b946020939093013593505050565b600060408284031215613ca757600080fd5b50919050565b600080600060608486031215613cc257600080fd5b8335613ccd81613c49565b92506020840135613cdd81613c49565b929592945050506040919091013590565b600060208284031215613d0057600080fd5b813561194481613c49565b600060208284031215613d1d57600080fd5b5035919050565b600060208284031215613d3657600080fd5b813567ffffffffffffffff811115613d4d57600080fd5b82016060818503121561194457600080fd5b60006101008284031215613ca757600080fd5b600060208284031215613d8457600080fd5b813567ffffffffffffffff811115613d9b57600080fd5b8201610120818503121561194457600080fd5b60008060008060608587031215613dc457600080fd5b843593506020850135613dd681613c49565b9250604085013567ffffffffffffffff80821115613df357600080fd5b818701915087601f830112613e0757600080fd5b813581811115613e1657600080fd5b886020828501011115613e2857600080fd5b95989497505060200194505050565b600060c08284031215613ca757600080fd5b60008060408385031215613e5c57600080fd5b8235613e6781613c49565b91506020830135613e7781613c49565b809150509250929050565b600181811c90821680613e9657607f821691505b602082108103613ca757634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613f0557613f05613ecc565b60405290565b60405160e0810167ffffffffffffffff81118282101715613f0557613f05613ecc565b600060408284031215613f4057600080fd5b613f48613ee2565b8235613f5381613c49565b81526020928301359281019290925250919050565b602081526000825160058110613f8e57634e487b7160e01b600052602160045260246000fd5b806020840152506020830151604080840152611e546060840182613c0a565b634e487b7160e01b600052601160045260246000fd5b808201808211156109e8576109e8613fad565b818103818111156109e8576109e8613fad565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b80820281158282048414176109e8576109e8613fad565b6000826140af57634e487b7160e01b600052601260045260246000fd5b500490565b81516001600160a01b0316815260208083015190820152604081016109e8565b60208082526034908201527f4e6174697665546f6b656e44657374696e6174696f6e3a20636f6e7472616374604082015273081d5b99195c98dbdb1b185d195c985b1a5e995960621b606082015260800190565b60006020828403121561413a57600080fd5b5051919050565b6020808252825182820152828101516001600160a01b039081166040808501919091528401518051821660608501528083015160808501526000929161010085019190606087015160a0870152608087015160e060c0880152805193849052840192600092506101208701905b808410156141d0578451831682529385019360019390930192908501906141ae565b5060a0880151878203601f190160e089015294506141ee8186613c0a565b98975050505050505050565b6000823561011e1983360301811261421157600080fd5b9190910192915050565b6000808335601e1984360301811261423257600080fd5b83018035915067ffffffffffffffff82111561424d57600080fd5b60200191503681900382131561426257600080fd5b9250929050565b602081526142836020820183516001600160a01b03169052565b60208201516040820152600060408301516142a960608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a08301516101608060c08501526142e5610180850183613c0a565b915060c085015160e085015260e085015161010061430d818701836001600160a01b03169052565b860151610120868101919091528601519050610140614336818701836001600160a01b03169052565b959095015193019290925250919050565b80358252602081013561435981613c49565b6001600160a01b03908116602084015260408201359061437882613c49565b908116604084015260608201359061438f82613c49565b1660608301526080818101359083015260a090810135910152565b604081526143bb6040820184614347565b600060c0840135601e198536030181126143d457600080fd5b840160208101903567ffffffffffffffff8111156143f157600080fd5b80360382131561440057600080fd5b610120610100818187015282610160870152610180838582890137600081858901015260e089013583880152614437828a01613c5e565b6001600160a01b031661014088015260208701979097525050601f01601f19169092019092019392505050565b602080825260339082015260008051602061489e833981519152604082015272206d756c74692d686f702066616c6c6261636b60681b606082015260800190565b6020808252603e908201526000805160206148be83398151915260408201527f6c69642064657374696e6174696f6e2062726964676520616464726573730000606082015260800190565b60e081016144fe8285614347565b8260c08301529392505050565b60208152815160208201526000602083015160018060a01b038082166040850152806040860151166060850152606085015160808501526080850151915060e060a085015261455e610100850183613c0a565b915060a085015160c08501528060c08601511660e0850152508091505092915050565b600082601f83011261459257600080fd5b815167ffffffffffffffff808211156145ad576145ad613ecc565b604051601f8301601f19908116603f011681019082821181831017156145d5576145d5613ecc565b816040528381528660208588010111156145ee57600080fd5b6145ff846020830160208901613be6565b9695505050505050565b60006020828403121561461b57600080fd5b815167ffffffffffffffff8082111561463357600080fd5b908301906040828603121561464757600080fd5b61464f613ee2565b82516005811061465e57600080fd5b815260208301518281111561467257600080fd5b61467e87828601614581565b60208301525095945050505050565b8051612f5381613c49565b6000604082840312156146aa57600080fd5b6146b2613ee2565b82516146bd81613c49565b81526020928301519281019290925250919050565b6000602082840312156146e457600080fd5b815167ffffffffffffffff808211156146fc57600080fd5b9083019060e0828603121561471057600080fd5b614718613f0b565b825181526147286020840161468d565b60208201526147396040840161468d565b60408201526060830151606082015260808301518281111561475a57600080fd5b61476687828601614581565b60808301525060a083015160a082015261478260c0840161468d565b60c082015295945050505050565b6000602082840312156147a257600080fd5b815161194481613c49565b600060c082840312156147bf57600080fd5b60405160c0810181811067ffffffffffffffff821117156147e2576147e2613ecc565b6040528235815260208301356147f781613c49565b6020820152604083013561480a81613c49565b6040820152606083013561481d81613c49565b60608201526080838101359082015260a0928301359281019290925250919050565b8381526001600160a01b038316602082015260606040820181905260009061306590830184613c0a565b60006020828403121561487b57600080fd5b8151801515811461194457600080fd5b60008251614211818460208701613be656fe54656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f54656c65706f72746572546f6b656e44657374696e6174696f6e3a20696e7661a2646970667358221220ca87d02f4314bedb8db8ca47536ccfd5eb10516eaafdde21c7a784e2fe434c9b64736f6c6343000812003354656c65706f72746572546f6b656e44657374696e6174696f6e3a207a65726f",
}

// NativeTokenDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenDestinationMetaData.ABI instead.
var NativeTokenDestinationABI = NativeTokenDestinationMetaData.ABI

// NativeTokenDestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenDestinationMetaData.Bin instead.
var NativeTokenDestinationBin = NativeTokenDestinationMetaData.Bin

// DeployNativeTokenDestination deploys a new Ethereum contract, binding an instance of NativeTokenDestination to it.
func DeployNativeTokenDestination(auth *bind.TransactOpts, backend bind.ContractBackend, settings NativeTokenDestinationSettings) (common.Address, *types.Transaction, *NativeTokenDestination, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenDestinationBin), backend, settings)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NativeTokenDestination is an auto generated Go binding around an Ethereum contract.
type NativeTokenDestination struct {
	NativeTokenDestinationCaller     // Read-only binding to the contract
	NativeTokenDestinationTransactor // Write-only binding to the contract
	NativeTokenDestinationFilterer   // Log filterer for contract events
}

// NativeTokenDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenDestinationSession struct {
	Contract     *NativeTokenDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeTokenDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenDestinationCallerSession struct {
	Contract *NativeTokenDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativeTokenDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenDestinationTransactorSession struct {
	Contract     *NativeTokenDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativeTokenDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenDestinationRaw struct {
	Contract *NativeTokenDestination // Generic contract binding to access the raw methods on
}

// NativeTokenDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCallerRaw struct {
	Contract *NativeTokenDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactorRaw struct {
	Contract *NativeTokenDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenDestination creates a new instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestination(address common.Address, backend bind.ContractBackend) (*NativeTokenDestination, error) {
	contract, err := bindNativeTokenDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NewNativeTokenDestinationCaller creates a new read-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenDestinationCaller, error) {
	contract, err := bindNativeTokenDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCaller{contract: contract}, nil
}

// NewNativeTokenDestinationTransactor creates a new write-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenDestinationTransactor, error) {
	contract, err := bindNativeTokenDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransactor{contract: contract}, nil
}

// NewNativeTokenDestinationFilterer creates a new log filterer instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenDestinationFilterer, error) {
	contract, err := bindNativeTokenDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationFilterer{contract: contract}, nil
}

// bindNativeTokenDestination binds a generic wrapper to an already deployed contract.
func bindNativeTokenDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.NativeTokenDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transact(opts, method, params...)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDFORBRIDGEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_FOR_BRIDGE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDFORBRIDGEADDRESS is a free data retrieval call binding the contract method 0x47a9a22c.
//
// Solidity: function BURNED_FOR_BRIDGE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDFORBRIDGEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDFORBRIDGEADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPCALLGASPERWORD(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MULTIHOPREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "MULTI_HOP_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// MULTIHOPREQUIREDGAS is a free data retrieval call binding the contract method 0x525975e6.
//
// Solidity: function MULTI_HOP_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MULTIHOPREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.MULTIHOPREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenDestination.Contract.NATIVEMINTER(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) REGISTERDESTINATIONREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "REGISTER_DESTINATION_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// REGISTERDESTINATIONREQUIREDGAS is a free data retrieval call binding the contract method 0x55c26725.
//
// Solidity: function REGISTER_DESTINATION_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) REGISTERDESTINATIONREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REGISTERDESTINATIONREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SOURCECHAINBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "SOURCE_CHAIN_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// SOURCECHAINBURNADDRESS is a free data retrieval call binding the contract method 0xecd4ed1b.
//
// Solidity: function SOURCE_CHAIN_BURN_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SOURCECHAINBURNADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.SOURCECHAINBURNADDRESS(&_NativeTokenDestination.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.Allowance(&_NativeTokenDestination.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NativeTokenDestination.Contract.BalanceOf(&_NativeTokenDestination.CallOpts, account)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.BlockchainID(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BurnedFeesReportingRewardPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "burnedFeesReportingRewardPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// BurnedFeesReportingRewardPercentage is a free data retrieval call binding the contract method 0x3a23dfe2.
//
// Solidity: function burnedFeesReportingRewardPercentage() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BurnedFeesReportingRewardPercentage() (*big.Int, error) {
	return _NativeTokenDestination.Contract.BurnedFeesReportingRewardPercentage(&_NativeTokenDestination.CallOpts)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _NativeTokenDestination.Contract.CalculateNumWords(&_NativeTokenDestination.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Decimals() (uint8, error) {
	return _NativeTokenDestination.Contract.Decimals(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenDestination.Contract.GetMinTeleporterVersion(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsRegistered() (bool, error) {
	return _NativeTokenDestination.Contract.IsRegistered(&_NativeTokenDestination.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsRegistered() (bool, error) {
	return _NativeTokenDestination.Contract.IsRegistered(&_NativeTokenDestination.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenDestination.Contract.IsTeleporterAddressPaused(&_NativeTokenDestination.CallOpts, teleporterAddress)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) LastestBurnedFeesReported(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "lastestBurnedFeesReported")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// LastestBurnedFeesReported is a free data retrieval call binding the contract method 0xd10a5b8c.
//
// Solidity: function lastestBurnedFeesReported() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) LastestBurnedFeesReported() (*big.Int, error) {
	return _NativeTokenDestination.Contract.LastestBurnedFeesReported(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) MultiplyOnDestination(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "multiplyOnDestination")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) MultiplyOnDestination() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnDestination(&_NativeTokenDestination.CallOpts)
}

// MultiplyOnDestination is a free data retrieval call binding the contract method 0x2001eff6.
//
// Solidity: function multiplyOnDestination() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) MultiplyOnDestination() (bool, error) {
	return _NativeTokenDestination.Contract.MultiplyOnDestination(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Name() (string, error) {
	return _NativeTokenDestination.Contract.Name(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Owner() (common.Address, error) {
	return _NativeTokenDestination.Contract.Owner(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) Symbol() (string, error) {
	return _NativeTokenDestination.Contract.Symbol(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterRegistry(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenMultiplier() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TokenMultiplier(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TokenSourceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "tokenSourceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TokenSourceAddress is a free data retrieval call binding the contract method 0xf5ea0603.
//
// Solidity: function tokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.TokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalNativeAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalNativeAssetSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalNativeAssetSupply is a free data retrieval call binding the contract method 0x1906529c.
//
// Solidity: function totalNativeAssetSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalNativeAssetSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalNativeAssetSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Approve(&_NativeTokenDestination.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.DecreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Deposit() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Deposit(&_NativeTokenDestination.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.IncreaseAllowance(&_NativeTokenDestination.TransactOpts, spender, addedValue)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.PauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) RegisterWithSource(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "registerWithSource", feeInfo)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) RegisterWithSource(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RegisterWithSource(&_NativeTokenDestination.TransactOpts, feeInfo)
}

// RegisterWithSource is a paid mutator transaction binding the contract method 0x151637f6.
//
// Solidity: function registerWithSource((address,uint256) feeInfo) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) RegisterWithSource(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RegisterWithSource(&_NativeTokenDestination.TransactOpts, feeInfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.RenounceOwnership(&_NativeTokenDestination.TransactOpts)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReportBurnedTxFees(opts *bind.TransactOpts, requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "reportBurnedTxFees", requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// ReportBurnedTxFees is a paid mutator transaction binding the contract method 0x55538c8b.
//
// Solidity: function reportBurnedTxFees(uint256 requiredGasLimit) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReportBurnedTxFees(requiredGasLimit *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportBurnedTxFees(&_NativeTokenDestination.TransactOpts, requiredGasLimit)
}

// Send is a paid mutator transaction binding the contract method 0xcb6a6417.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xcb6a6417.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xcb6a6417.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Send(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0xc1f823c7.
//
// Solidity: function sendAndCall(((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0xc1f823c7.
//
// Solidity: function sendAndCall(((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0xc1f823c7.
//
// Solidity: function sendAndCall(((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCall(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCallMultiHop is a paid mutator transaction binding the contract method 0x8be887c9.
//
// Solidity: function sendAndCallMultiHop((((bytes32,address,address,address,uint256,uint256),bytes,uint256,address),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendAndCallMultiHop(opts *bind.TransactOpts, input SendAndCallMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendAndCallMultiHop", input)
}

// SendAndCallMultiHop is a paid mutator transaction binding the contract method 0x8be887c9.
//
// Solidity: function sendAndCallMultiHop((((bytes32,address,address,address,uint256,uint256),bytes,uint256,address),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendAndCallMultiHop(input SendAndCallMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCallMultiHop(&_NativeTokenDestination.TransactOpts, input)
}

// SendAndCallMultiHop is a paid mutator transaction binding the contract method 0x8be887c9.
//
// Solidity: function sendAndCallMultiHop((((bytes32,address,address,address,uint256,uint256),bytes,uint256,address),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) SendAndCallMultiHop(input SendAndCallMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendAndCallMultiHop(&_NativeTokenDestination.TransactOpts, input)
}

// SendMulithop is a paid mutator transaction binding the contract method 0xbe0c6b1c.
//
// Solidity: function sendMulithop(((bytes32,address,address,address,uint256,uint256),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) SendMulithop(opts *bind.TransactOpts, input SendMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "sendMulithop", input)
}

// SendMulithop is a paid mutator transaction binding the contract method 0xbe0c6b1c.
//
// Solidity: function sendMulithop(((bytes32,address,address,address,uint256,uint256),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) SendMulithop(input SendMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendMulithop(&_NativeTokenDestination.TransactOpts, input)
}

// SendMulithop is a paid mutator transaction binding the contract method 0xbe0c6b1c.
//
// Solidity: function sendMulithop(((bytes32,address,address,address,uint256,uint256),(uint256,address)) input) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) SendMulithop(input SendMultiHopInput) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.SendMulithop(&_NativeTokenDestination.TransactOpts, input)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Transfer(&_NativeTokenDestination.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferFrom(&_NativeTokenDestination.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferOwnership(&_NativeTokenDestination.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UnpauseTeleporterAddress(&_NativeTokenDestination.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.UpdateMinTeleporterVersion(&_NativeTokenDestination.TransactOpts, version)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Withdraw(&_NativeTokenDestination.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Fallback(&_NativeTokenDestination.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.Receive(&_NativeTokenDestination.TransactOpts)
}

// NativeTokenDestinationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NativeTokenDestination contract.
type NativeTokenDestinationApprovalIterator struct {
	Event *NativeTokenDestinationApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationApproval represents a Approval event raised by the NativeTokenDestination contract.
type NativeTokenDestinationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NativeTokenDestinationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationApprovalIterator{contract: _NativeTokenDestination.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationApproval)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseApproval(log types.Log) (*NativeTokenDestinationApproval, error) {
	event := new(NativeTokenDestinationApproval)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailedIterator struct {
	Event *NativeTokenDestinationCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallFailed represents a CallFailed event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallFailedIterator{contract: _NativeTokenDestination.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallFailed)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallFailed(log types.Log) (*NativeTokenDestinationCallFailed, error) {
	event := new(NativeTokenDestinationCallFailed)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceededIterator struct {
	Event *NativeTokenDestinationCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCallSucceeded represents a CallSucceeded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenDestinationCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCallSucceededIterator{contract: _NativeTokenDestination.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCallSucceeded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenDestinationCallSucceeded, error) {
	event := new(NativeTokenDestinationCallSucceeded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the NativeTokenDestination contract.
type NativeTokenDestinationDepositIterator struct {
	Event *NativeTokenDestinationDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationDeposit represents a Deposit event raised by the NativeTokenDestination contract.
type NativeTokenDestinationDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationDepositIterator{contract: _NativeTokenDestination.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationDeposit)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseDeposit(log types.Log) (*NativeTokenDestinationDeposit, error) {
	event := new(NativeTokenDestinationDeposit)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenDestinationMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationMinTeleporterVersionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenDestination contract.
type NativeTokenDestinationMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenDestinationMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationMinTeleporterVersionUpdatedIterator{contract: _NativeTokenDestination.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenDestinationMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenDestinationMinTeleporterVersionUpdated)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferredIterator struct {
	Event *NativeTokenDestinationOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenDestination contract.
type NativeTokenDestinationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenDestinationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationOwnershipTransferredIterator{contract: _NativeTokenDestination.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationOwnershipTransferred)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenDestinationOwnershipTransferred, error) {
	event := new(NativeTokenDestinationOwnershipTransferred)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationReportBurnedTxFeesIterator is returned from FilterReportBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportBurnedTxFees events raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFeesIterator struct {
	Event *NativeTokenDestinationReportBurnedTxFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationReportBurnedTxFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationReportBurnedTxFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationReportBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationReportBurnedTxFees represents a ReportBurnedTxFees event raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportBurnedTxFees struct {
	TeleporterMessageID [32]byte
	FeesBurned          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportBurnedTxFees is a free log retrieval operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterReportBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenDestinationReportBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationReportBurnedTxFeesIterator{contract: _NativeTokenDestination.contract, event: "ReportBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportBurnedTxFees is a free log subscription operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchReportBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationReportBurnedTxFees, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "ReportBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationReportBurnedTxFees)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportBurnedTxFees is a log parse operation binding the contract event 0x0832c643b65d6d3724ed14ac3a655fbc7cae54fb010918b2c2f70ef6b1bb94a5.
//
// Solidity: event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseReportBurnedTxFees(log types.Log) (*NativeTokenDestinationReportBurnedTxFees, error) {
	event := new(NativeTokenDestinationReportBurnedTxFees)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTeleporterAddressPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressPausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressPaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenDestinationTeleporterAddressPaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressPaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenDestinationTeleporterAddressUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTeleporterAddressUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenDestinationTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTeleporterAddressUnpausedIterator{contract: _NativeTokenDestination.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTeleporterAddressUnpaused)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenDestinationTeleporterAddressUnpaused, error) {
	event := new(NativeTokenDestinationTeleporterAddressUnpaused)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSentIterator struct {
	Event *NativeTokenDestinationTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0xb3209fd9e86173715f7081e24de11ecce66057a53b36d0733f6574c55e390c46.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, ((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensAndCallSentIterator{contract: _NativeTokenDestination.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0xb3209fd9e86173715f7081e24de11ecce66057a53b36d0733f6574c55e390c46.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, ((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensAndCallSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0xb3209fd9e86173715f7081e24de11ecce66057a53b36d0733f6574c55e390c46.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, ((bytes32,address,address,address,uint256,uint256),bytes,uint256,address) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenDestinationTokensAndCallSent, error) {
	event := new(NativeTokenDestinationTokensAndCallSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSentIterator struct {
	Event *NativeTokenDestinationTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensSent represents a TokensSent event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x8a4d0e8ef2f385ab7df8c78e89773da8ae3679f1258fd71823e8c39e552a9564.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenDestinationTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensSentIterator{contract: _NativeTokenDestination.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x8a4d0e8ef2f385ab7df8c78e89773da8ae3679f1258fd71823e8c39e552a9564.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensSent)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0x8a4d0e8ef2f385ab7df8c78e89773da8ae3679f1258fd71823e8c39e552a9564.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensSent(log types.Log) (*NativeTokenDestinationTokensSent, error) {
	event := new(NativeTokenDestinationTokensSent)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawnIterator struct {
	Event *NativeTokenDestinationTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenDestinationTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTokensWithdrawnIterator{contract: _NativeTokenDestination.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTokensWithdrawn)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenDestinationTokensWithdrawn, error) {
	event := new(NativeTokenDestinationTokensWithdrawn)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferIterator struct {
	Event *NativeTokenDestinationTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTransfer represents a Transfer event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NativeTokenDestinationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferIterator{contract: _NativeTokenDestination.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTransfer)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransfer(log types.Log) (*NativeTokenDestinationTransfer, error) {
	event := new(NativeTokenDestinationTransfer)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawalIterator struct {
	Event *NativeTokenDestinationWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenDestinationWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenDestinationWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenDestinationWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationWithdrawal represents a Withdrawal event raised by the NativeTokenDestination contract.
type NativeTokenDestinationWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*NativeTokenDestinationWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationWithdrawalIterator{contract: _NativeTokenDestination.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationWithdrawal)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseWithdrawal(log types.Log) (*NativeTokenDestinationWithdrawal, error) {
	event := new(NativeTokenDestinationWithdrawal)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
