package keeper

import (
	"context"
	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReceiveOrder(goCtx context.Context, msg *types.MsgReceiveOrder) (*types.MsgReceiveOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.GetOrders(ctx, uint64(msg.Id))
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	if order.State == types.StateReceive {
		return nil, sdkerrors.Wrapf(types.ErrOrderHasReceived, "")
	}

	order.State = types.StateReceive
	k.SetOrders(ctx,order)

	return &types.MsgReceiveOrderResponse{}, nil
}
