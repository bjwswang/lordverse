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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWarehouse = "create_warehouse"
	TypeMsgSignWarehouse   = "sign_warehouse"
	TypeMsgUpdateWarehouse = "update_warehouse"
	TypeMsgDeleteWarehouse = "delete_warehouse"
)

var _ sdk.Msg = &MsgCreateWarehouse{}

func NewMsgCreateWarehouse(creator string, voters []uint64, threshold uint64) *MsgCreateWarehouse {
	return &MsgCreateWarehouse{
		Creator:   creator,
		Voters:    voters,
		Threshold: threshold,
	}
}

func (msg *MsgCreateWarehouse) Route() string {
	return RouterKey
}

func (msg *MsgCreateWarehouse) Type() string {
	return TypeMsgCreateWarehouse
}

func (msg *MsgCreateWarehouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWarehouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWarehouse{}

func NewMsgUpdateWarehouse(creator string, id uint64, voters []uint64, threshold uint64) *MsgUpdateWarehouse {
	return &MsgUpdateWarehouse{
		Id:        id,
		Creator:   creator,
		Voters:    voters,
		Threshold: threshold,
	}
}

func (msg *MsgUpdateWarehouse) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWarehouse) Type() string {
	return TypeMsgUpdateWarehouse
}

func (msg *MsgUpdateWarehouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWarehouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgSignWarehouse{}

func NewMsgSignWarehouse(creator string, warehouseID uint64, voterID uint64) *MsgSignWarehouse {
	return &MsgSignWarehouse{
		Id:      warehouseID,
		Voter:   voterID,
		Creator: creator,
	}
}

func (msg *MsgSignWarehouse) Route() string {
	return RouterKey
}

func (msg *MsgSignWarehouse) Type() string {
	return TypeMsgSignWarehouse
}

func (msg *MsgSignWarehouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignWarehouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWarehouse{}

func NewMsgDeleteWarehouse(creator string, id uint64) *MsgDeleteWarehouse {
	return &MsgDeleteWarehouse{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteWarehouse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWarehouse) Type() string {
	return TypeMsgDeleteWarehouse
}

func (msg *MsgDeleteWarehouse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWarehouse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWarehouse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
