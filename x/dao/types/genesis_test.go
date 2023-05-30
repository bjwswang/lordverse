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

package types_test

import (
	"testing"

	"lordverse/x/dao/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				WarehouseList: []types.Warehouse{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				WarehouseCount: 2,
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated warehouse",
			genState: &types.GenesisState{
				WarehouseList: []types.Warehouse{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid warehouse count",
			genState: &types.GenesisState{
				WarehouseList: []types.Warehouse{
					{
						Id: 1,
					},
				},
				WarehouseCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated proposal",
			genState: &types.GenesisState{
				ProposalList: []types.Proposal{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid proposal count",
			genState: &types.GenesisState{
				ProposalList: []types.Proposal{
					{
						Id: 1,
					},
				},
				ProposalCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated voter",
			genState: &types.GenesisState{
				VoterList: []types.Voter{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid voter count",
			genState: &types.GenesisState{
				VoterList: []types.Voter{
					{
						Id: 1,
					},
				},
				VoterCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
