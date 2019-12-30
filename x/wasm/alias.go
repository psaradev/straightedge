// nolint
package wasm

import (
	"github.com/heystraightedge/straightedge/x/wasm/internal/keeper"
	"github.com/heystraightedge/straightedge/x/wasm/internal/types"
)

// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmwasm/wasmd/x/wasm/internal/keeper/
// ALIASGEN: github.com/cosmwasm/wasmd/x/wasm/internal/types/

const (
	GasMultiplier         = keeper.GasMultiplier
	MaxGas                = keeper.MaxGas
	QueryListContracts    = keeper.QueryListContracts
	QueryGetContract      = keeper.QueryGetContract
	QueryGetContractState = keeper.QueryGetContractState
	QueryGetCode          = keeper.QueryGetCode
	QueryListCode         = keeper.QueryListCode
	DefaultCodespace      = types.DefaultCodespace
	CodeCreatedFailed     = types.CodeCreatedFailed
	CodeAccountExists     = types.CodeAccountExists
	CodeInstantiateFailed = types.CodeInstantiateFailed
	CodeExecuteFailed     = types.CodeExecuteFailed
	CodeGasLimit          = types.CodeGasLimit
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	TStoreKey             = types.TStoreKey
	QuerierRoute          = types.QuerierRoute
	RouterKey             = types.RouterKey
	MaxWasmSize           = types.MaxWasmSize
)

var (
	// functions aliases
	NewKeeper                 = keeper.NewKeeper
	NewQuerier                = keeper.NewQuerier
	MakeTestCodec             = keeper.MakeTestCodec
	CreateTestInput           = keeper.CreateTestInput
	RegisterCodec             = types.RegisterCodec
	ErrCreateFailed           = types.ErrCreateFailed
	ErrAccountExists          = types.ErrAccountExists
	ErrInstantiateFailed      = types.ErrInstantiateFailed
	ErrExecuteFailed          = types.ErrExecuteFailed
	ErrGasLimit               = types.ErrGasLimit
	GetCodeKey                = types.GetCodeKey
	GetContractAddressKey     = types.GetContractAddressKey
	GetContractStorePrefixKey = types.GetContractStorePrefixKey
	NewCodeInfo               = types.NewCodeInfo
	NewParams                 = types.NewParams
	NewWasmCoins              = types.NewWasmCoins
	NewContract               = types.NewContract
	CosmosResult              = types.CosmosResult

	// variable aliases
	ModuleCdc           = types.ModuleCdc
	KeyLastCodeID       = types.KeyLastCodeID
	KeyLastInstanceID   = types.KeyLastInstanceID
	CodeKeyPrefix       = types.CodeKeyPrefix
	ContractKeyPrefix   = types.ContractKeyPrefix
	ContractStorePrefix = types.ContractStorePrefix
)

type (
	Keeper                 = keeper.Keeper
	GetCodeResponse        = keeper.GetCodeResponse
	MsgStoreCode           = types.MsgStoreCode
	MsgInstantiateContract = types.MsgInstantiateContract
	MsgExecuteContract     = types.MsgExecuteContract
	CodeInfo               = types.CodeInfo
	Contract               = types.Contract
)
