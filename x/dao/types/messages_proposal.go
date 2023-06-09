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
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProposal  = "create_proposal"
	TypeMsgUpdateProposal  = "update_proposal"
	TypeMsgVoteProposal    = "vote_proposal"
	TypeMsgDeleteProposal  = "delete_proposal"
	TypeMsgExecuteProposal = "execute_proposal"
)

var _ sdk.Msg = &MsgCreateProposal{}

func NewMsgCreateProposal(creator string, warehouse uint64, title string, description string, expiration *time.Time) *MsgCreateProposal {
	return &MsgCreateProposal{
		Creator:     creator,
		Warehouse:   warehouse,
		Title:       title,
		Description: description,
		Expiration:  expiration,
	}
}

func (msg *MsgCreateProposal) Route() string {
	return RouterKey
}

func (msg *MsgCreateProposal) Type() string {
	return TypeMsgCreateProposal
}

func (msg *MsgCreateProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProposal{}

func NewMsgUpdateProposal(creator string, id uint64, description string) *MsgUpdateProposal {
	return &MsgUpdateProposal{
		Id:          id,
		Creator:     creator,
		Description: description,
	}
}

func (msg *MsgUpdateProposal) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProposal) Type() string {
	return TypeMsgUpdateProposal
}

func (msg *MsgUpdateProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgVoteProposal{}

func NewMsgVoteProposal(creator string, id uint64, voterID uint64, decision VoteType, explanation string) *MsgVoteProposal {
	return &MsgVoteProposal{
		Id:          id,
		VoterID:     voterID,
		Creator:     creator,
		Decision:    decision,
		Explanation: explanation,
	}
}

func (msg *MsgVoteProposal) Route() string {
	return RouterKey
}

func (msg *MsgVoteProposal) Type() string {
	return TypeMsgVoteProposal
}

func (msg *MsgVoteProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// validate vote type
	switch msg.Decision {
	case VoteType_VOTE_TYPE_ABSTAIN:
	case VoteType_VOTE_TYPE_YES:
	case VoteType_VOTE_TYPE_NO:
	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid vote type %d", msg.Decision)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteProposal{}

func NewMsgDeleteProposal(creator string, id uint64) *MsgDeleteProposal {
	return &MsgDeleteProposal{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteProposal) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProposal) Type() string {
	return TypeMsgDeleteProposal
}

func (msg *MsgDeleteProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgExecuteProposal{}

func NewMsgExecuteProposal(creator string, id uint64) *MsgDeleteProposal {
	return &MsgDeleteProposal{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgExecuteProposal) Route() string {
	return RouterKey
}

func (msg *MsgExecuteProposal) Type() string {
	return TypeMsgExecuteProposal
}

func (msg *MsgExecuteProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExecuteProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
