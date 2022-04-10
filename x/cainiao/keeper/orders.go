package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/cainiao/x/cainiao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetOrdersCount get the total number of orders
func (k Keeper) GetOrdersCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OrdersCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOrdersCount set the total number of orders
func (k Keeper) SetOrdersCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OrdersCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOrders appends a orders in the store with a new id and update the count
func (k Keeper) AppendOrders(
	ctx sdk.Context,
	orders types.Orders,
) uint64 {
	// Create the orders
	count := k.GetOrdersCount(ctx)

	// Set the ID of the appended value
	orders.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	appendedValue := k.cdc.MustMarshal(&orders)
	store.Set(GetOrdersIDBytes(orders.Id), appendedValue)

	// Update orders count
	k.SetOrdersCount(ctx, count+1)

	return count
}

// SetOrders set a specific orders in the store
func (k Keeper) SetOrders(ctx sdk.Context, orders types.Orders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	b := k.cdc.MustMarshal(&orders)
	store.Set(GetOrdersIDBytes(orders.Id), b)
}

// GetOrders returns a orders from its id
func (k Keeper) GetOrders(ctx sdk.Context, id uint64) (val types.Orders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	b := store.Get(GetOrdersIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrders removes a orders from the store
func (k Keeper) RemoveOrders(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	store.Delete(GetOrdersIDBytes(id))
}

// GetAllOrders returns all orders
func (k Keeper) GetAllOrders(ctx sdk.Context) (list []types.Orders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Orders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOrdersIDBytes returns the byte representation of the ID
func GetOrdersIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetOrdersIDFromBytes returns ID in uint64 format from a byte array
func GetOrdersIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
