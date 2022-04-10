package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	time2 "time"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateOrderState(goCtx context.Context, msg *types.MsgUpdateOrderState) (*types.MsgUpdateOrderStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id, _ := strconv.Atoi(msg.Id)
	order, found := k.GetOrders(ctx, uint64(id))
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}

	if order.State == types.StateReceive {
		return nil, sdkerrors.Wrapf(types.ErrOrderHasReceived, "")
	}

	//更新状态
	//updateAccount, err := sdk.AccAddressFromBech32(msg.Creator)
	//if err != nil {
	//	return nil, sdkerrors.Wrapf(types.ErrGetMsgCreatorAddress, "")
	//}

	time := time2.Now().Format("2006-01-02 15:04:05")
	station := "操作者：" + msg.Creator + " 时间：" + time + "位置：" + msg.Station + "\n"
	order.Station = order.Station + station

	if order.State == types.StateWait {
		order.State = types.StateStart
	} else {
		order.State = types.StateOntheWay
	}

	k.SetOrders(ctx, order)

	return &types.MsgUpdateOrderStateResponse{}, nil
}
