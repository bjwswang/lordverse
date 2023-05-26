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
	TypeMsgCreateLand = "create_land"
	TypeMsgUpdateLand = "update_land"
	TypeMsgDeleteLand = "delete_land"
)

var _ sdk.Msg = &MsgCreateLand{}

func NewMsgCreateLand(creator string, owner string, location string, grade string, assessment string, extra string) *MsgCreateLand {
	return &MsgCreateLand{
		Creator:    creator,
		Owner:      owner,
		Location:   location,
		Grade:      grade,
		Assessment: assessment,
		Extra:      extra,
	}
}

func (msg *MsgCreateLand) Route() string {
	return RouterKey
}

func (msg *MsgCreateLand) Type() string {
	return TypeMsgCreateLand
}

func (msg *MsgCreateLand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateLand{}

func NewMsgUpdateLand(creator string, id uint64, owner string, location string, grade string, assessment string, extra string) *MsgUpdateLand {
	return &MsgUpdateLand{
		Id:         id,
		Creator:    creator,
		Owner:      owner,
		Location:   location,
		Grade:      grade,
		Assessment: assessment,
		Extra:      extra,
	}
}

func (msg *MsgUpdateLand) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLand) Type() string {
	return TypeMsgUpdateLand
}

func (msg *MsgUpdateLand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLand{}

func NewMsgDeleteLand(creator string, id uint64) *MsgDeleteLand {
	return &MsgDeleteLand{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteLand) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLand) Type() string {
	return TypeMsgDeleteLand
}

func (msg *MsgDeleteLand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
