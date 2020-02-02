package market

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	core "github.com/terra-project/core/types"
	"github.com/terra-project/core/x/market/internal/keeper"
)

func TestMarketFilters(t *testing.T) {
	input, h := setup(t)

	// Case 1: non-oracle message being sent fails
	bankMsg := bank.MsgSend{}
	_, err := h(input.Ctx, bankMsg)
	require.Error(t, err)

	// Case 2: Normal MsgSwap submission goes through
	offerCoin := sdk.NewCoin(core.MicroLunaDenom, sdk.NewInt(10))
	prevoteMsg := NewMsgSwap(keeper.Addrs[0], offerCoin, core.MicroSDRDenom)
	_, err = h(input.Ctx, prevoteMsg)
	require.NoError(t, err)
}

func TestSwapMsg(t *testing.T) {
	input, h := setup(t)

	beforeTerraPoolDelta := input.MarketKeeper.GetTerraPoolDelta(input.Ctx)

	amt := sdk.NewInt(10)
	offerCoin := sdk.NewCoin(core.MicroLunaDenom, amt)
	swapMsg := NewMsgSwap(keeper.Addrs[0], offerCoin, core.MicroSDRDenom)
	_, err := h(input.Ctx, swapMsg)
	require.NoError(t, err)

	afterTerraPoolDelta := input.MarketKeeper.GetTerraPoolDelta(input.Ctx)
	diff := beforeTerraPoolDelta.Sub(afterTerraPoolDelta)
	price, _ := input.OracleKeeper.GetLunaExchangeRate(input.Ctx, core.MicroSDRDenom)
	require.Equal(t, price.MulInt(amt), diff.Abs())

	swapMsg = NewMsgSwap(keeper.Addrs[0], offerCoin, core.MicroLunaDenom)
	_, err = h(input.Ctx, swapMsg)
	require.Error(t, err)
}
