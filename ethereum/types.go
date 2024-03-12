// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	// NodeVersion is the version of geth we are using.
	NodeVersion = "1.9.24"

	// Blockchain is Ethereum.
	Blockchain string = "Corechain"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// RopstenNetwork is the value of the network
	// in RopstenNetworkIdentifier.
	RopstenNetwork string = "Ropsten"

	// RinkebyNetwork is the value of the network
	// in RinkebyNetworkNetworkIdentifier.
	RinkebyNetwork string = "Rinkeby"

	// GoerliNetwork is the value of the network
	// in GoerliNetworkNetworkIdentifier.
	GoerliNetwork string = "Goerli"

	// DevNetwork is the value of the network
	// in DevNetworkNetworkIdentifier.
	DevNetwork string = "Dev"

	// CoreNetwork is the value of the network
	// in CoreNetworkNetworkIdentifier.
	CoreNetwork string = "Core"

	// BuffaloNetwork is the value of the network
	// in BuffaloNetworkNetworkIdentifier.
	BuffaloNetwork string = "Buffalo"

	// Symbol is the symbol value
	// used in Currency.
	Symbol = "CORE"

	// Decimals is the decimals value
	// used in Currency.
	Decimals = 18

	// MinerRewardOpType is used to describe
	// a miner block reward.
	MinerRewardOpType = "MINER_REWARD"

	// UncleRewardOpType is used to describe
	// an uncle block reward.
	UncleRewardOpType = "UNCLE_REWARD"

	// FeeOpType is used to represent fee operations.
	FeeOpType = "FEE"

	// CallOpType is used to represent CALL trace operations.
	CallOpType = "CALL"

	// CreateOpType is used to represent CREATE trace operations.
	CreateOpType = "CREATE"

	// Create2OpType is used to represent CREATE2 trace operations.
	Create2OpType = "CREATE2"

	// SelfDestructOpType is used to represent SELFDESTRUCT trace operations.
	SelfDestructOpType = "SELFDESTRUCT"

	// CallCodeOpType is used to represent CALLCODE trace operations.
	CallCodeOpType = "CALLCODE"

	// DelegateCallOpType is used to represent DELEGATECALL trace operations.
	DelegateCallOpType = "DELEGATECALL"

	// StaticCallOpType is used to represent STATICCALL trace operations.
	StaticCallOpType = "STATICCALL"

	// DestructOpType is a synthetic operation used to represent the
	// deletion of suicided accounts that still have funds at the end
	// of a transaction.
	DestructOpType = "DESTRUCT"

	// SuccessStatus is the status of any
	// Ethereum operation considered successful.
	SuccessStatus = "SUCCESS"

	// FailureStatus is the status of any
	// Ethereum operation considered unsuccessful.
	FailureStatus = "FAILURE"

	// HistoricalBalanceSupported is whether
	// historical balance is supported.
	HistoricalBalanceSupported = true

	// UnclesRewardMultiplier is the uncle reward
	// multiplier.
	UnclesRewardMultiplier = 32

	// MaxUncleDepth is the maximum depth for
	// an uncle to be rewarded.
	MaxUncleDepth = 8

	// GenesisBlockIndex is the index of the
	// genesis block.
	GenesisBlockIndex = int64(0)

	// TransferGasLimit is the gas limit
	// of a transfer.
	TransferGasLimit = int64(21000) //nolint:gomnd

	// MainnetGethArguments are the arguments to start a mainnet geth instance.
	MainnetGethArguments = `--config=/app/ethereum/geth.toml --cache=8000 --gcmode=archive --graphql`

	// IncludeMempoolCoins does not apply to rosetta-core as it is not UTXO-based.
	IncludeMempoolCoins = false
)

// CoreChain Genesis hashes and Network configurations to enforce below configs on.
var (
	DevGenesisHash     = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
	CoreGenesisHash    = common.HexToHash("0xf7fc87f11e61508a5828cd1508060ed1714c8d32a92744ae10acb43c953357ad")
	BuffaloGenesisHash = common.HexToHash("0xd90508c51efd64e75363cdf51114d9f2a90a79e6cd0f78f3c3038b47695c034a")

	CoreChainConfig = &params.ChainConfig{
		ChainID: big.NewInt(1116),
	}

	BuffaloChainConfig = &params.ChainConfig{
		ChainID: big.NewInt(1115),
	}

	DevChainConfig = &params.ChainConfig{
		ChainID: big.NewInt(1112),
	}
)

var (
	// RopstenGethArguments are the arguments to start a ropsten geth instance.
	RopstenGethArguments = fmt.Sprintf("%s --ropsten", MainnetGethArguments)

	// RinkebyGethArguments are the arguments to start a rinkeby geth instance.
	RinkebyGethArguments = fmt.Sprintf("%s --rinkeby", MainnetGethArguments)

	// GoerliGethArguments are the arguments to start a ropsten geth instance.
	GoerliGethArguments = fmt.Sprintf("%s --goerli", MainnetGethArguments)

	// DevGethArguments are the arguments to start a dev geth instance.
	DevGethArguments = MainnetGethArguments

	// CoreGethArguments are the arguments to start a core geth instance
	CoreGethArguments = MainnetGethArguments

	// BuffaloGethArguments are the arguments to start a buffalo geth instance
	BuffaloGethArguments = MainnetGethArguments

	// MainnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the mainnet genesis block.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.MainnetGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// RopstenGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Ropsten genesis block.
	RopstenGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.RopstenGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// RinkebyGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Ropsten genesis block.
	RinkebyGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.RinkebyGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// GoerliGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Goerli genesis block.
	GoerliGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.GoerliGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// DevGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Corechain Devnet genesis block
	DevGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  DevGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// CoreGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Corechain Mainnet genesis block
	CoreGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  CoreGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// BuffaloGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the Corechain Testnet genesis block
	BuffaloGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  BuffaloGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// Currency is the *types.Currency for all
	// Ethereum networks.
	Currency = &types.Currency{
		Symbol:   Symbol,
		Decimals: Decimals,
	}

	// OperationTypes are all suppoorted operation types.
	OperationTypes = []string{
		MinerRewardOpType,
		UncleRewardOpType,
		FeeOpType,
		CallOpType,
		CreateOpType,
		Create2OpType,
		SelfDestructOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
		DestructOpType,
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	// CallMethods are all supported call methods.
	CallMethods = []string{
		"eth_getBlockByNumber",
		"eth_getTransactionReceipt",
		"eth_call",
		"eth_estimateGas",
	}
)

// JSONRPC is the interface for accessing go-ethereum's JSON RPC endpoint.
type JSONRPC interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Close()
}

// GraphQL is the interface for accessing go-ethereum's GraphQL endpoint.
type GraphQL interface {
	Query(ctx context.Context, input string) (string, error)
}

// CallType returns a boolean indicating
// if the provided trace type is a call type.
func CallType(t string) bool {
	callTypes := []string{
		CallOpType,
		CallCodeOpType,
		DelegateCallOpType,
		StaticCallOpType,
	}

	for _, callType := range callTypes {
		if callType == t {
			return true
		}
	}

	return false
}

// CreateType returns a boolean indicating
// if the provided trace type is a create type.
func CreateType(t string) bool {
	createTypes := []string{
		CreateOpType,
		Create2OpType,
	}

	for _, createType := range createTypes {
		if createType == t {
			return true
		}
	}

	return false
}
