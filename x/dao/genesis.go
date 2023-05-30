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

package dao

import (
	"lordverse/x/dao/keeper"
	"lordverse/x/dao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the warehouse
	for _, elem := range genState.WarehouseList {
		k.SetWarehouse(ctx, elem)
	}

	// Set warehouse count
	k.SetWarehouseCount(ctx, genState.WarehouseCount)

	// Set proposal count
	k.SetProposalCount(ctx, genState.ProposalCount)
	// Set all the voter
	for _, elem := range genState.VoterList {
		k.SetVoter(ctx, elem)
	}

	// Set voter count
	k.SetVoterCount(ctx, genState.VoterCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WarehouseList = k.GetAllWarehouse(ctx)
	genesis.WarehouseCount = k.GetWarehouseCount(ctx)
	genesis.ProposalList = k.GetAllProposal(ctx)
	genesis.ProposalCount = k.GetProposalCount(ctx)
	genesis.VoterList = k.GetAllVoter(ctx)
	genesis.VoterCount = k.GetVoterCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
