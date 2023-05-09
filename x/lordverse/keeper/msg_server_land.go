package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lordverse/x/lordverse/types"
)

func (k msgServer) CreateLand(goCtx context.Context, msg *types.MsgCreateLand) (*types.MsgCreateLandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var land = types.Land{
		Creator:    msg.Creator,
		Owner:      msg.Owner,
		Location:   msg.Location,
		Grade:      msg.Grade,
		Assessment: msg.Assessment,
		Extra:      msg.Extra,
	}

	id := k.AppendLand(
		ctx,
		land,
	)

	return &types.MsgCreateLandResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateLand(goCtx context.Context, msg *types.MsgUpdateLand) (*types.MsgUpdateLandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var land = types.Land{
		Creator:    msg.Creator,
		Id:         msg.Id,
		Owner:      msg.Owner,
		Location:   msg.Location,
		Grade:      msg.Grade,
		Assessment: msg.Assessment,
		Extra:      msg.Extra,
	}

	// Checks that the element exists
	val, found := k.GetLand(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetLand(ctx, land)

	return &types.MsgUpdateLandResponse{}, nil
}

func (k msgServer) DeleteLand(goCtx context.Context, msg *types.MsgDeleteLand) (*types.MsgDeleteLandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetLand(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLand(ctx, msg.Id)

	return &types.MsgDeleteLandResponse{}, nil
}
