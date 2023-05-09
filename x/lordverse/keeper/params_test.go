package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lordverse/testutil/keeper"
	"lordverse/x/lordverse/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LordverseKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
