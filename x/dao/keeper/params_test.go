package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lordverse/testutil/keeper"
	"lordverse/x/dao/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DaoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
