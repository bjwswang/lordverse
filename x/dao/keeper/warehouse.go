/*
Copyright 2023 The Bestchains Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lordverse/x/dao/types"
)

// GetWarehouseCount get the total number of warehouse
func (k Keeper) GetWarehouseCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WarehouseCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWarehouseCount set the total number of warehouse
func (k Keeper) SetWarehouseCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.WarehouseCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWarehouse appends a warehouse in the store with a new id and update the count
func (k Keeper) AppendWarehouse(
	ctx sdk.Context,
	warehouse types.Warehouse,
) uint64 {
	// Create the warehouse
	count := k.GetWarehouseCount(ctx)

	// Set the ID of the appended value
	warehouse.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WarehouseKey))
	appendedValue := k.cdc.MustMarshal(&warehouse)
	store.Set(GetWarehouseIDBytes(warehouse.Id), appendedValue)

	// Update warehouse count
	k.SetWarehouseCount(ctx, count+1)

	return count
}

// SetWarehouse set a specific warehouse in the store
func (k Keeper) SetWarehouse(ctx sdk.Context, warehouse types.Warehouse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WarehouseKey))
	b := k.cdc.MustMarshal(&warehouse)
	store.Set(GetWarehouseIDBytes(warehouse.Id), b)
}

// GetWarehouse returns a warehouse from its id
func (k Keeper) GetWarehouse(ctx sdk.Context, id uint64) (val types.Warehouse, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WarehouseKey))
	b := store.Get(GetWarehouseIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWarehouse removes a warehouse from the store
func (k Keeper) RemoveWarehouse(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WarehouseKey))
	store.Delete(GetWarehouseIDBytes(id))
}

// GetAllWarehouse returns all warehouse
func (k Keeper) GetAllWarehouse(ctx sdk.Context) (list []types.Warehouse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WarehouseKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Warehouse
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWarehouseIDBytes returns the byte representation of the ID
func GetWarehouseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWarehouseIDFromBytes returns ID in uint64 format from a byte array
func GetWarehouseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
