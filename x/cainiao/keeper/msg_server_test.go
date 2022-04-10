package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/cainiao/testutil/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CainiaoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
