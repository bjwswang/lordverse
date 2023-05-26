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

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "lordverse/testutil/keeper"
	"lordverse/testutil/nullify"
	"lordverse/x/dao/types"
)

func TestProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposal(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProposalRequest
		response *types.QueryGetProposalResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetProposalRequest{Id: msgs[0].Id},
			response: &types.QueryGetProposalResponse{Proposal: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetProposalRequest{Id: msgs[1].Id},
			response: &types.QueryGetProposalResponse{Proposal: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetProposalRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Proposal(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestProposalQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProposal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProposalRequest {
		return &types.QueryAllProposalRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Proposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Proposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Proposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
