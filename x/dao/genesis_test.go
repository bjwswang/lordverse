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

package dao_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lordverse/testutil/keeper"
	"lordverse/testutil/nullify"
	"lordverse/x/dao"
	"lordverse/x/dao/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WarehouseList: []types.Warehouse{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		WarehouseCount: 2,
		VoteList: []types.Vote{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VoteCount: 2,
		ProposalList: []types.Proposal{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ProposalCount: 2,
		VoterList: []types.Voter{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VoterCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DaoKeeper(t)
	dao.InitGenesis(ctx, *k, genesisState)
	got := dao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WarehouseList, got.WarehouseList)
	require.Equal(t, genesisState.WarehouseCount, got.WarehouseCount)
	require.ElementsMatch(t, genesisState.VoteList, got.VoteList)
	require.Equal(t, genesisState.VoteCount, got.VoteCount)
	require.ElementsMatch(t, genesisState.ProposalList, got.ProposalList)
	require.Equal(t, genesisState.ProposalCount, got.ProposalCount)
	require.ElementsMatch(t, genesisState.VoterList, got.VoterList)
	require.Equal(t, genesisState.VoterCount, got.VoterCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
