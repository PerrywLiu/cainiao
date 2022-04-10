package keeper

import (
	"context"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddOrder(goCtx context.Context, msg *types.MsgAddOrder) (*types.MsgAddOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order := types.Orders{
		SendName: msg.SendName,
		SendAddress: msg.SendAddress,
		SendTel: msg.SendTel,
		TargetName: msg.ReceiveName,
		TargetAddress: msg.ReceiveAddress,
		TargetTel: msg.ReceiveTel,
		State: types.StateWait,
		Station: "",
	}

	k.AppendOrders(ctx,order)

	return &types.MsgAddOrderResponse{}, nil
}
