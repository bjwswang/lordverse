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
	"github.com/stretchr/testify/require"
	keepertest "lordverse/testutil/keeper"
	"lordverse/testutil/nullify"
	"lordverse/x/dao/keeper"
	"lordverse/x/dao/types"
)

func createNProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Id = keeper.AppendProposal(ctx, items[i])
	}
	return items
}

func TestProposalGet(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProposal(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposal(ctx, item.Id)
		_, found := keeper.GetProposal(ctx, item.Id)
		require.False(t, found)
	}
}

func TestProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposal(ctx)),
	)
}

func TestProposalCount(t *testing.T) {
	keeper, ctx := keepertest.DaoKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetProposalCount(ctx))
}
