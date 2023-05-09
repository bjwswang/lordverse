package lordverse

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lordverse/x/lordverse/keeper"
	"lordverse/x/lordverse/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the land
	for _, elem := range genState.LandList {
		k.SetLand(ctx, elem)
	}

	// Set land count
	k.SetLandCount(ctx, genState.LandCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LandList = k.GetAllLand(ctx)
	genesis.LandCount = k.GetLandCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
