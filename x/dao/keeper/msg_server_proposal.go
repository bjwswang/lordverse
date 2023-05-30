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

	// Check that the proposal expiration time
	if ctx.BlockTime().After(*msg.Expiration) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal expired time should be after current block time")
	}
	// Check that the warehouse exists and has been activated
	warehouse, found := k.GetWarehouse(ctx, msg.Warehouse)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("warehouse %d doesn't exist", msg.Warehouse))
	}
	if !warehouse.Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("warehouse %d is not active", msg.Warehouse))
	}

	var proposal = types.Proposal{
		Warehouse:   msg.Warehouse,
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Expiration:  msg.Expiration,
		Votes:       make([]uint64, types.VoteType_VOTE_TYPE_NO+1),
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

	// Checks that the element exists
	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("proposal %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != proposal.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check that the proposal status
	if proposal.Status != types.ProposalStatus_PROPOSAL_STATUS_VOTING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal status should be PROPOSAL_STATUS_VOTING")
	}
	// Check that the proposal expiration time
	if !proposal.Expiration.After(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal has expired")
	}

	// Update the proposal
	if proposal.Description != msg.Description {
		proposal.Description = msg.Description
		k.SetProposal(ctx, proposal)
	}

	return &types.MsgUpdateProposalResponse{}, nil
}

func (k msgServer) VoteProposal(goCtx context.Context, msg *types.MsgVoteProposal) (*types.MsgVoteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the proposal exists
	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("proposal %d doesn't exist", msg.Id))
	}
	// Check that the proposal status
	if proposal.Status != types.ProposalStatus_PROPOSAL_STATUS_VOTING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal status should be PROPOSAL_STATUS_VOTING")
	}

	// Check whether the voter already voted
	votedVoters := proposal.GetVotedVoters()
	if votedVoters == nil {
		votedVoters = make(map[uint64]types.VoteType)
	}
	if _, ok := votedVoters[msg.VoterID]; ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "already voted")
	}

	// Check that the proposal expiration time
	// Current status is voting,but excceeds the expiration time
	if !proposal.Expiration.After(ctx.BlockTime()) {
		// set proposal as expired
		proposal.Status = types.ProposalStatus_PROPOSAL_STATUS_FAILED
		k.SetProposal(ctx, proposal)

		return &types.MsgVoteProposalResponse{
			Status: proposal.Status,
		}, nil
	}

	// Check the voter
	// 1. voter exists
	voter, found := k.GetVoter(ctx, msg.VoterID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("voter %d doesn't exist", msg.VoterID))
	}
	// 2. message sender is the voter owner
	if voter.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect voter owner")
	}
	// 3. warehouse exists
	warehouse, found := k.GetWarehouse(ctx, proposal.Warehouse)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("proposal %d 's warehouse %d doesn't exist", proposal.Id, proposal.Warehouse))
	}
	// 4. voter is in warehouse voter list
	if !warehouse.SignedVoters[msg.VoterID] {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "not a correct warehouse voter")
	}

	// All checks passed, record vote
	votedVoters[msg.VoterID] = msg.Decision
	votes := proposal.GetVotes()
	if votes == nil {
		votes = make([]uint64, types.VoteType_VOTE_TYPE_NO+1)
	}
	votes[msg.Decision]++

	// Update the proposal
	if votes[types.VoteType_VOTE_TYPE_YES] >= warehouse.Threshold {
		proposal.Status = types.ProposalStatus_PROPOSAL_STATUS_SUCCESS
	} else {
		var totalVotes uint64
		for _, voterID := range warehouse.Voters {
			voter, found := k.GetVoter(ctx, voterID)
			if !found {
				return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("voter %d in warehouse doesn't exist", voterID))
			}
			totalVotes += voter.Weight
		}
		// When Votes(NO+Abstain) > (TotalVotes-Threshold), threshold never reached,mark proposal as failed
		if votes[types.VoteType_VOTE_TYPE_NO]+votes[types.VoteType_VOTE_TYPE_ABSTAIN]+warehouse.Threshold > totalVotes {
			proposal.Status = types.ProposalStatus_PROPOSAL_STATUS_FAILED
		}
	}

	proposal.VotedVoters = votedVoters
	proposal.Votes = votes
	k.SetProposal(ctx, proposal)

	return &types.MsgVoteProposalResponse{
		Status: proposal.Status,
	}, nil
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

	// Check that the proposal status
	if val.Status != types.ProposalStatus_PROPOSAL_STATUS_VOTING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal can be canceled only when status should be PROPOSAL_STATUS_VOTING")
	}

	val.Status = types.ProposalStatus_PROPOSAL_STATUS_CANCELLED

	k.SetProposal(ctx, val)

	return &types.MsgDeleteProposalResponse{}, nil
}

func (k msgServer) ExecuteProposal(goCtx context.Context, msg *types.MsgExecuteProposal) (*types.MsgExecuteProposalResponse, error) {
	// TODO: implement code
	return &types.MsgExecuteProposalResponse{}, nil
}
