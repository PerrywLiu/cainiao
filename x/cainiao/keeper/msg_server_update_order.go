package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	time2 "time"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateOrder(goCtx context.Context, msg *types.MsgUpdateOrder) (*types.MsgUpdateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.GetOrders(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	if order.State == types.StateReceive {
		return nil, sdkerrors.Wrapf(types.ErrOrderHasReceived, "")
	}

	//更新状态
	updateAccount, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrGetMsgCreatorAddress, "")
	}

	time := time2.Now().Format("2006-01-02 15:04:05")
	station := "操作者：" + string(updateAccount) + " 时间：" + time + "时间：" + msg.Station
	order.Station = order.Station + "\n" + station

	if order.State == types.StateWait {
		order.State = types.StateStart
	} else {
		order.State = types.StateOntheWay
	}

	k.SetOrders(ctx, order)

	return &types.MsgUpdateOrderResponse{}, nil
}
