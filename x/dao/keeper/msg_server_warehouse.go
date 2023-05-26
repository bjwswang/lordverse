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

package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lordverse/x/dao/types"
)

func (k msgServer) CreateWarehouse(goCtx context.Context, msg *types.MsgCreateWarehouse) (*types.MsgCreateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var warehouse = types.Warehouse{
		Creator:   msg.Creator,
		Voters:    msg.Voters,
		Threshold: msg.Threshold,
		Active:    msg.Active,
	}

	id := k.AppendWarehouse(
		ctx,
		warehouse,
	)

	return &types.MsgCreateWarehouseResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateWarehouse(goCtx context.Context, msg *types.MsgUpdateWarehouse) (*types.MsgUpdateWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var warehouse = types.Warehouse{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Voters:    msg.Voters,
		Threshold: msg.Threshold,
		Active:    msg.Active,
	}

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWarehouse(ctx, warehouse)

	return &types.MsgUpdateWarehouseResponse{}, nil
}

func (k msgServer) DeleteWarehouse(goCtx context.Context, msg *types.MsgDeleteWarehouse) (*types.MsgDeleteWarehouseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetWarehouse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWarehouse(ctx, msg.Id)

	return &types.MsgDeleteWarehouseResponse{}, nil
}
