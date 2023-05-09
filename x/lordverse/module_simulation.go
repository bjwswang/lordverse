package lordverse

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"lordverse/testutil/sample"
	lordversesimulation "lordverse/x/lordverse/simulation"
	"lordverse/x/lordverse/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lordversesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateLand = "op_weight_msg_land"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLand int = 100

	opWeightMsgUpdateLand = "op_weight_msg_land"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLand int = 100

	opWeightMsgDeleteLand = "op_weight_msg_land"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLand int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lordverseGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		LandList: []types.Land{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		LandCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lordverseGenesis)
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

	var weightMsgCreateLand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLand, &weightMsgCreateLand, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLand = defaultWeightMsgCreateLand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLand,
		lordversesimulation.SimulateMsgCreateLand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLand, &weightMsgUpdateLand, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLand = defaultWeightMsgUpdateLand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLand,
		lordversesimulation.SimulateMsgUpdateLand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteLand, &weightMsgDeleteLand, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLand = defaultWeightMsgDeleteLand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLand,
		lordversesimulation.SimulateMsgDeleteLand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateLand,
			defaultWeightMsgCreateLand,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lordversesimulation.SimulateMsgCreateLand(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLand,
			defaultWeightMsgUpdateLand,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lordversesimulation.SimulateMsgUpdateLand(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteLand,
			defaultWeightMsgDeleteLand,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				lordversesimulation.SimulateMsgDeleteLand(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
