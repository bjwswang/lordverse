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

const (
	DefaultVoterWeight = 1
)

func (k msgServer) CreateVoter(goCtx context.Context, msg *types.MsgCreateVoter) (*types.MsgCreateVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var voter = types.Voter{
		Creator: msg.Creator,
		Weight:  msg.Weight,
	}

	id := k.AppendVoter(
		ctx,
		voter,
	)

	return &types.MsgCreateVoterResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateVoter(goCtx context.Context, msg *types.MsgUpdateVoter) (*types.MsgUpdateVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var voter = types.Voter{
		Creator: msg.Creator,
		Id:      msg.Id,
		Weight:  msg.Weight,
	}

	// Checks that the element exists
	val, found := k.GetVoter(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetVoter(ctx, voter)

	return &types.MsgUpdateVoterResponse{}, nil
}

func (k msgServer) DeleteVoter(goCtx context.Context, msg *types.MsgDeleteVoter) (*types.MsgDeleteVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetVoter(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveVoter(ctx, msg.Id)

	return &types.MsgDeleteVoterResponse{}, nil
}
