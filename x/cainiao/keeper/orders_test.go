package keeper_test

import (
	"testing"

	keepertest "github.com/cosmonaut/cainiao/testutil/keeper"
	"github.com/cosmonaut/cainiao/testutil/nullify"
	"github.com/cosmonaut/cainiao/x/cainiao/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNOrders(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Orders {
	items := make([]types.Orders, n)
	for i := range items {
		items[i].Id = keeper.AppendOrders(ctx, items[i])
	}
	return items
}

func TestOrdersGet(t *testing.T) {
	keeper, ctx := keepertest.CainiaoKeeper(t)
	items := createNOrders(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetOrders(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestOrdersRemove(t *testing.T) {
	keeper, ctx := keepertest.CainiaoKeeper(t)
	items := createNOrders(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrders(ctx, item.Id)
		_, found := keeper.GetOrders(ctx, item.Id)
		require.False(t, found)
	}
}

func TestOrdersGetAll(t *testing.T) {
	keeper, ctx := keepertest.CainiaoKeeper(t)
	items := createNOrders(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrders(ctx)),
	)
}

func TestOrdersCount(t *testing.T) {
	keeper, ctx := keepertest.CainiaoKeeper(t)
	items := createNOrders(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetOrdersCount(ctx))
}
