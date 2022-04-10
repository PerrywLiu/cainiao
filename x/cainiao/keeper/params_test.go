package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/cainiao/testutil/keeper"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CainiaoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
