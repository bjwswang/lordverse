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
	"context"
	"fmt"

	"lordverse/x/dao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateWarehouse(goCtx context.Context, msg *types.MsgCreateWarehouse) (*types.MsgCreateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check voters
	// 1. voter with same creator exist
	// 2. voter is not duplicated
	var voters = make(map[string]struct{})
	var totalWeight uint64
	for _, voterID := range msg.Voters {
		voter, found := k.GetVoter(ctx, voterID)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("voter %d doesn't exist", voterID))
		}
		_, found = voters[voter.Creator]
		if found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("voter %d already exists", voterID))
		}
		voters[voter.Creator] = struct{}{}
		totalWeight = voter.Weight + totalWeight
	}
	// Check threshold
	if msg.Threshold > totalWeight {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("threshold %d is greater than total weight %d", msg.Threshold, totalWeight))
	}

	var warehouse = types.Warehouse{
		Creator:   msg.Creator,
		Voters:    msg.Voters,
		Threshold: msg.Threshold,
		// Set to false by default
		Active: false,
	}

	id := k.AppendWarehouse(
		ctx,
		warehouse,
	)

	return &types.MsgCreateWarehouseResponse{
		Id: id,
	}, nil
}

func (k msgServer) SignWarehouse(goCtx context.Context, msg *types.MsgSignWarehouse) (*types.MsgSignWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	warehouse, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("warehouse %d doesn't exist", msg.Id))
	}

	// Check warehouse has been activated
	if warehouse.Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("warehouse %d is already active", msg.Id))
	}

	// Check voter already signed
	signedVoters := warehouse.SignedVoters
	if signedVoters == nil {
		signedVoters = make(map[uint64]bool)
	}
	if signedVoters[msg.Voter] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("voter %d already signed", msg.Voter))
	}

	// Check voter
	voter, found := k.GetVoter(ctx, msg.Voter)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("voter %d doesn't exist", msg.Voter))
	}
	if voter.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner of voter")
	}

	// Update warehouse's signed voters
	signedVoters[msg.Voter] = true
	warehouse.SignedVoters = signedVoters
	// Set active to true if threshold is reached
	if len(warehouse.SignedVoters) == len(warehouse.Voters) {
		warehouse.Active = true
	}
	k.SetWarehouse(ctx, warehouse)

	return &types.MsgSignWarehouseResponse{
		SignedVoters: warehouse.SignedVoters,
	}, nil
}

func (k msgServer) UpdateWarehouse(goCtx context.Context, msg *types.MsgUpdateWarehouse) (*types.MsgUpdateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var warehouse = types.Warehouse{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Voters:    msg.Voters,
		Threshold: msg.Threshold,
	}

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWarehouse(ctx, warehouse)

	return &types.MsgUpdateWarehouseResponse{}, nil
}

func (k msgServer) DeleteWarehouse(goCtx context.Context, msg *types.MsgDeleteWarehouse) (*types.MsgDeleteWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWarehouse(ctx, msg.Id)

	return &types.MsgDeleteWarehouseResponse{}, nil
}
