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

// GetProposalCount get the total number of proposal
func (k Keeper) GetProposalCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProposalCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetProposalCount set the total number of proposal
func (k Keeper) SetProposalCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProposalCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendProposal appends a proposal in the store with a new id and update the count
func (k Keeper) AppendProposal(
	ctx sdk.Context,
	proposal types.Proposal,
) uint64 {
	// Create the proposal
	count := k.GetProposalCount(ctx)

	// Set the ID of the appended value
	proposal.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKey))
	appendedValue := k.cdc.MustMarshal(&proposal)
	store.Set(GetProposalIDBytes(proposal.Id), appendedValue)

	// Update proposal count
	k.SetProposalCount(ctx, count+1)

	return count
}

// SetProposal set a specific proposal in the store
func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKey))
	b := k.cdc.MustMarshal(&proposal)
	store.Set(GetProposalIDBytes(proposal.Id), b)
}

// GetProposal returns a proposal from its id
func (k Keeper) GetProposal(ctx sdk.Context, id uint64) (val types.Proposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKey))
	b := store.Get(GetProposalIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposal removes a proposal from the store
func (k Keeper) RemoveProposal(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKey))
	store.Delete(GetProposalIDBytes(id))
}

// GetAllProposal returns all proposal
func (k Keeper) GetAllProposal(ctx sdk.Context) (list []types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProposalIDBytes returns the byte representation of the ID
func GetProposalIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetProposalIDFromBytes returns ID in uint64 format from a byte array
func GetProposalIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
