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

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lordverse/x/dao/types"
)

func (k Keeper) VoterAll(goCtx context.Context, req *types.QueryAllVoterRequest) (*types.QueryAllVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var voters []types.Voter
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	voterStore := prefix.NewStore(store, types.KeyPrefix(types.VoterKey))

	pageRes, err := query.Paginate(voterStore, req.Pagination, func(key []byte, value []byte) error {
		var voter types.Voter
		if err := k.cdc.Unmarshal(value, &voter); err != nil {
			return err
		}

		voters = append(voters, voter)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoterResponse{Voter: voters, Pagination: pageRes}, nil
}

func (k Keeper) Voter(goCtx context.Context, req *types.QueryGetVoterRequest) (*types.QueryGetVoterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	voter, found := k.GetVoter(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVoterResponse{Voter: voter}, nil
}
