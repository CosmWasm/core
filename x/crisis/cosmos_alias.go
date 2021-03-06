// nolint
package crisis

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
)

const (
	DefaultCodespace  = crisis.DefaultCodespace
	CodeInvalidInput  = crisis.CodeInvalidInput
	ModuleName        = crisis.ModuleName
	DefaultParamspace = crisis.DefaultParamspace
)

var (
	// functions aliases
	ErrNilSender          = crisis.ErrNilSender
	ErrUnknownInvariant   = crisis.ErrUnknownInvariant
	NewGenesisState       = crisis.NewGenesisState
	DefaultGenesisState   = crisis.DefaultGenesisState
	NewMsgVerifyInvariant = crisis.NewMsgVerifyInvariant
	ParamKeyTable         = crisis.ParamKeyTable
	NewInvarRoute         = crisis.NewInvarRoute
	NewKeeper             = crisis.NewKeeper
	NewCosmosAppModule    = crisis.NewAppModule

	// variable aliases
	CosmosModuleCdc          = crisis.ModuleCdc
	ParamStoreKeyConstantFee = crisis.ParamStoreKeyConstantFee
)

type (
	GenesisState         = crisis.GenesisState
	MsgVerifyInvariant   = crisis.MsgVerifyInvariant
	InvarRoute           = crisis.InvarRoute
	Keeper               = crisis.Keeper
	CosmosAppModule      = crisis.AppModule
	CosmosAppModuleBasic = crisis.AppModuleBasic
)
