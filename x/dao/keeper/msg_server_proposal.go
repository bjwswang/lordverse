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

func (k msgServer) CreateProposal(goCtx context.Context, msg *types.MsgCreateProposal) (*types.MsgCreateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var proposal = types.Proposal{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Expiration:  msg.Expiration,
	}

	// Check that the proposal expiration time
	if ctx.BlockTime().After(*msg.Expiration) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal expired time should be after current block time")
	}

	id := k.AppendProposal(
		ctx,
		proposal,
	)

	return &types.MsgCreateProposalResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateProposal(goCtx context.Context, msg *types.MsgUpdateProposal) (*types.MsgUpdateProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var proposal = types.Proposal{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Title:       msg.Title,
		Description: msg.Description,
		Expiration:  msg.Expiration,
	}

	// Checks that the element exists
	val, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check that the proposal expiration time not changed
	if !msg.Expiration.Equal(*val.Expiration) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal expired time can't be changed")
	}

	k.SetProposal(ctx, proposal)

	return &types.MsgUpdateProposalResponse{}, nil
}

func (k msgServer) DeleteProposal(goCtx context.Context, msg *types.MsgDeleteProposal) (*types.MsgDeleteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProposal(ctx, msg.Id)

	return &types.MsgDeleteProposalResponse{}, nil
}

func (k msgServer) ExecuteProposal(goCtx context.Context, msg *types.MsgExecuteProposal) (*types.MsgExecuteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProposal(ctx, msg.Id)

	return &types.MsgExecuteProposalResponse{}, nil
}
