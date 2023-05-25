package keeper

import (
	"lordverse/x/dao/types"
)

var _ types.QueryServer = Keeper{}
