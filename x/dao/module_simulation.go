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
	"math/rand"

	"lordverse/testutil/sample"
	daosimulation "lordverse/x/dao/simulation"
	"lordverse/x/dao/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = daosimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWarehouse int = 100

	opWeightMsgUpdateWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWarehouse int = 100

	opWeightMsgDeleteWarehouse = "op_weight_msg_warehouse"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWarehouse int = 100

	opWeightMsgCreateProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProposal int = 100

	opWeightMsgUpdateProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProposal int = 100

	opWeightMsgDeleteProposal = "op_weight_msg_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProposal int = 100

	opWeightMsgCreateVoter = "op_weight_msg_voter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVoter int = 100

	opWeightMsgUpdateVoter = "op_weight_msg_voter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVoter int = 100

	opWeightMsgDeleteVoter = "op_weight_msg_voter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVoter int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	daoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		WarehouseList: []types.Warehouse{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		WarehouseCount: 2,
		ProposalList: []types.Proposal{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ProposalCount: 2,
		VoterList: []types.Voter{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		VoterCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&daoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateWarehouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWarehouse, &weightMsgCreateWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWarehouse = defaultWeightMsgCreateWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWarehouse,
		daosimulation.SimulateMsgCreateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWarehouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWarehouse, &weightMsgUpdateWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWarehouse = defaultWeightMsgUpdateWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWarehouse,
		daosimulation.SimulateMsgUpdateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWarehouse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWarehouse, &weightMsgDeleteWarehouse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWarehouse = defaultWeightMsgDeleteWarehouse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWarehouse,
		daosimulation.SimulateMsgDeleteWarehouse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProposal, &weightMsgCreateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProposal = defaultWeightMsgCreateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProposal,
		daosimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProposal, &weightMsgUpdateProposal, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProposal = defaultWeightMsgUpdateProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProposal,
		daosimulation.SimulateMsgUpdateProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProposal, &weightMsgDeleteProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProposal = defaultWeightMsgDeleteProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProposal,
		daosimulation.SimulateMsgDeleteProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVoter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVoter, &weightMsgCreateVoter, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVoter = defaultWeightMsgCreateVoter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVoter,
		daosimulation.SimulateMsgCreateVoter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVoter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVoter, &weightMsgUpdateVoter, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVoter = defaultWeightMsgUpdateVoter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVoter,
		daosimulation.SimulateMsgUpdateVoter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVoter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVoter, &weightMsgDeleteVoter, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVoter = defaultWeightMsgDeleteVoter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVoter,
		daosimulation.SimulateMsgDeleteVoter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateWarehouse,
			defaultWeightMsgCreateWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgCreateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateWarehouse,
			defaultWeightMsgUpdateWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgUpdateWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteWarehouse,
			defaultWeightMsgDeleteWarehouse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgDeleteWarehouse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProposal,
			defaultWeightMsgCreateProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgCreateProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProposal,
			defaultWeightMsgUpdateProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgUpdateProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProposal,
			defaultWeightMsgDeleteProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgDeleteProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateVoter,
			defaultWeightMsgCreateVoter,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgCreateVoter(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateVoter,
			defaultWeightMsgUpdateVoter,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgUpdateVoter(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteVoter,
			defaultWeightMsgDeleteVoter,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgDeleteVoter(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
