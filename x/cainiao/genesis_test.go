package cainiao_test

import (
	"testing"

	keepertest "github.com/cosmonaut/cainiao/testutil/keeper"
	"github.com/cosmonaut/cainiao/testutil/nullify"
	"github.com/cosmonaut/cainiao/x/cainiao"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OrdersList: []types.Orders{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OrdersCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CainiaoKeeper(t)
	cainiao.InitGenesis(ctx, *k, genesisState)
	got := cainiao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OrdersList, got.OrdersList)
	require.Equal(t, genesisState.OrdersCount, got.OrdersCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
