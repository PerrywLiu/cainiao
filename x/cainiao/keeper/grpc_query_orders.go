package keeper

import (
	"context"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrdersAll(c context.Context, req *types.QueryAllOrdersRequest) (*types.QueryAllOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orderss []types.Orders
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ordersStore := prefix.NewStore(store, types.KeyPrefix(types.OrdersKey))

	pageRes, err := query.Paginate(ordersStore, req.Pagination, func(key []byte, value []byte) error {
		var orders types.Orders
		if err := k.cdc.Unmarshal(value, &orders); err != nil {
			return err
		}

		orderss = append(orderss, orders)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersResponse{Orders: orderss, Pagination: pageRes}, nil
}

func (k Keeper) Orders(c context.Context, req *types.QueryGetOrdersRequest) (*types.QueryGetOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	orders, found := k.GetOrders(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrdersResponse{Orders: orders}, nil
}
