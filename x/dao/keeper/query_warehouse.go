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

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lordverse/x/dao/types"
)

func (k Keeper) WarehouseAll(goCtx context.Context, req *types.QueryAllWarehouseRequest) (*types.QueryAllWarehouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var warehouses []types.Warehouse
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	warehouseStore := prefix.NewStore(store, types.KeyPrefix(types.WarehouseKey))

	pageRes, err := query.Paginate(warehouseStore, req.Pagination, func(key []byte, value []byte) error {
		var warehouse types.Warehouse
		if err := k.cdc.Unmarshal(value, &warehouse); err != nil {
			return err
		}

		warehouses = append(warehouses, warehouse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWarehouseResponse{Warehouse: warehouses, Pagination: pageRes}, nil
}

func (k Keeper) Warehouse(goCtx context.Context, req *types.QueryGetWarehouseRequest) (*types.QueryGetWarehouseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	warehouse, found := k.GetWarehouse(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetWarehouseResponse{Warehouse: warehouse}, nil
}
