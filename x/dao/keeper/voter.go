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

// GetVoterCount get the total number of voter
func (k Keeper) GetVoterCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VoterCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVoterCount set the total number of voter
func (k Keeper) SetVoterCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VoterCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVoter appends a voter in the store with a new id and update the count
func (k Keeper) AppendVoter(
	ctx sdk.Context,
	voter types.Voter,
) uint64 {
	// Create the voter
	count := k.GetVoterCount(ctx)

	// Set the ID of the appended value
	voter.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoterKey))
	appendedValue := k.cdc.MustMarshal(&voter)
	store.Set(GetVoterIDBytes(voter.Id), appendedValue)

	// Update voter count
	k.SetVoterCount(ctx, count+1)

	return count
}

// SetVoter set a specific voter in the store
func (k Keeper) SetVoter(ctx sdk.Context, voter types.Voter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoterKey))
	b := k.cdc.MustMarshal(&voter)
	store.Set(GetVoterIDBytes(voter.Id), b)
}

// GetVoter returns a voter from its id
func (k Keeper) GetVoter(ctx sdk.Context, id uint64) (val types.Voter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoterKey))
	b := store.Get(GetVoterIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVoter removes a voter from the store
func (k Keeper) RemoveVoter(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoterKey))
	store.Delete(GetVoterIDBytes(id))
}

// GetAllVoter returns all voter
func (k Keeper) GetAllVoter(ctx sdk.Context) (list []types.Voter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoterKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Voter
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVoterIDBytes returns the byte representation of the ID
func GetVoterIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVoterIDFromBytes returns ID in uint64 format from a byte array
func GetVoterIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
