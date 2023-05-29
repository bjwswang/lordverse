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

func (k msgServer) CreateVote(goCtx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vote = types.Vote{
		Creator:  msg.Creator,
		Decision: msg.Decision,
	}

	// Check proposal
	proposal, found := k.GetProposal(ctx, msg.ProposalId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("proposal %d doesn't exist", msg.ProposalId))
	}

	// Check proposal status
	if proposal.Status != types.ProposalStatus_PROPOSAL_STATUS_VOTING && proposal.Status != types.ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal has finished")
	}

	// proposal not started yet
	if ctx.BlockTime().Before(*proposal.Expiration) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal expired")
	}

	// proposal ended

	// TODO: Check this voter in proposal's voter list

	id := k.AppendVote(
		ctx,
		vote,
	)

	return &types.MsgCreateVoteResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateVote(goCtx context.Context, msg *types.MsgUpdateVote) (*types.MsgUpdateVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var vote = types.Vote{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Decision: msg.Decision,
	}

	// Checks that the element exists
	val, found := k.GetVote(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetVote(ctx, vote)

	return &types.MsgUpdateVoteResponse{}, nil
}

func (k msgServer) DeleteVote(goCtx context.Context, msg *types.MsgDeleteVote) (*types.MsgDeleteVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetVote(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveVote(ctx, msg.Id)

	return &types.MsgDeleteVoteResponse{}, nil
}
