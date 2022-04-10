package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cainiao module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrOrderHasReceived     = sdkerrors.Register(ModuleName, 1001, "orderHasReceived")
	ErrGetMsgCreatorAddress = sdkerrors.Register(ModuleName, 1002, "GetAddress faild")
)
