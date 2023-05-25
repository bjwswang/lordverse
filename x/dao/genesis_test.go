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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DaoKeeper(t)
	dao.InitGenesis(ctx, *k, genesisState)
	got := dao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
