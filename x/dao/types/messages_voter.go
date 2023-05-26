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
	TypeMsgCreateVoter = "create_voter"
	TypeMsgUpdateVoter = "update_voter"
	TypeMsgDeleteVoter = "delete_voter"
)

var _ sdk.Msg = &MsgCreateVoter{}

func NewMsgCreateVoter(creator string, weight uint32) *MsgCreateVoter {
	return &MsgCreateVoter{
		Creator: creator,
		Weight:  weight,
	}
}

func (msg *MsgCreateVoter) Route() string {
	return RouterKey
}

func (msg *MsgCreateVoter) Type() string {
	return TypeMsgCreateVoter
}

func (msg *MsgCreateVoter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVoter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVoter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVoter{}

func NewMsgUpdateVoter(creator string, id uint64, weight uint32) *MsgUpdateVoter {
	return &MsgUpdateVoter{
		Id:      id,
		Creator: creator,
		Weight:  weight,
	}
}

func (msg *MsgUpdateVoter) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVoter) Type() string {
	return TypeMsgUpdateVoter
}

func (msg *MsgUpdateVoter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVoter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVoter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVoter{}

func NewMsgDeleteVoter(creator string, id uint64) *MsgDeleteVoter {
	return &MsgDeleteVoter{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteVoter) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVoter) Type() string {
	return TypeMsgDeleteVoter
}

func (msg *MsgDeleteVoter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVoter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVoter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
