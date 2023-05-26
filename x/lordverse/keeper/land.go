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
	"lordverse/x/lordverse/types"
)

// GetLandCount get the total number of land
func (k Keeper) GetLandCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LandCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLandCount set the total number of land
func (k Keeper) SetLandCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LandCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLand appends a land in the store with a new id and update the count
func (k Keeper) AppendLand(
	ctx sdk.Context,
	land types.Land,
) uint64 {
	// Create the land
	count := k.GetLandCount(ctx)

	// Set the ID of the appended value
	land.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LandKey))
	appendedValue := k.cdc.MustMarshal(&land)
	store.Set(GetLandIDBytes(land.Id), appendedValue)

	// Update land count
	k.SetLandCount(ctx, count+1)

	return count
}

// SetLand set a specific land in the store
func (k Keeper) SetLand(ctx sdk.Context, land types.Land) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LandKey))
	b := k.cdc.MustMarshal(&land)
	store.Set(GetLandIDBytes(land.Id), b)
}

// GetLand returns a land from its id
func (k Keeper) GetLand(ctx sdk.Context, id uint64) (val types.Land, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LandKey))
	b := store.Get(GetLandIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLand removes a land from the store
func (k Keeper) RemoveLand(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LandKey))
	store.Delete(GetLandIDBytes(id))
}

// GetAllLand returns all land
func (k Keeper) GetAllLand(ctx sdk.Context) (list []types.Land) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LandKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Land
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLandIDBytes returns the byte representation of the ID
func GetLandIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLandIDFromBytes returns ID in uint64 format from a byte array
func GetLandIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
