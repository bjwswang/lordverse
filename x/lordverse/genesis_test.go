package lordverse_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lordverse/testutil/keeper"
	"lordverse/testutil/nullify"
	"lordverse/x/lordverse"
	"lordverse/x/lordverse/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LandList: []types.Land{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LandCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LordverseKeeper(t)
	lordverse.InitGenesis(ctx, *k, genesisState)
	got := lordverse.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LandList, got.LandList)
	require.Equal(t, genesisState.LandCount, got.LandCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
