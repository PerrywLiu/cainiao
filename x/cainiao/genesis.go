package cainiao

import (
	"github.com/cosmonaut/cainiao/x/cainiao/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the orders
	for _, elem := range genState.OrdersList {
		k.SetOrders(ctx, elem)
	}

	// Set orders count
	k.SetOrdersCount(ctx, genState.OrdersCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OrdersList = k.GetAllOrders(ctx)
	genesis.OrdersCount = k.GetOrdersCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
