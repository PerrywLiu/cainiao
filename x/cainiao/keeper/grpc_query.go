package keeper

import (
	"github.com/cosmonaut/cainiao/x/cainiao/types"
)

var _ types.QueryServer = Keeper{}
