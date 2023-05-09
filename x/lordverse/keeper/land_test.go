package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "lordverse/testutil/keeper"
	"lordverse/testutil/nullify"
	"lordverse/x/lordverse/keeper"
	"lordverse/x/lordverse/types"
)

func createNLand(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Land {
	items := make([]types.Land, n)
	for i := range items {
		items[i].Id = keeper.AppendLand(ctx, items[i])
	}
	return items
}

func TestLandGet(t *testing.T) {
	keeper, ctx := keepertest.LordverseKeeper(t)
	items := createNLand(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLand(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLandRemove(t *testing.T) {
	keeper, ctx := keepertest.LordverseKeeper(t)
	items := createNLand(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLand(ctx, item.Id)
		_, found := keeper.GetLand(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLandGetAll(t *testing.T) {
	keeper, ctx := keepertest.LordverseKeeper(t)
	items := createNLand(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLand(ctx)),
	)
}

func TestLandCount(t *testing.T) {
	keeper, ctx := keepertest.LordverseKeeper(t)
	items := createNLand(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLandCount(ctx))
}
