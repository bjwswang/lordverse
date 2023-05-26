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
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateWarehouse{}, "dao/CreateWarehouse", nil)
	cdc.RegisterConcrete(&MsgUpdateWarehouse{}, "dao/UpdateWarehouse", nil)
	cdc.RegisterConcrete(&MsgDeleteWarehouse{}, "dao/DeleteWarehouse", nil)
	cdc.RegisterConcrete(&MsgCreateVote{}, "dao/CreateVote", nil)
	cdc.RegisterConcrete(&MsgUpdateVote{}, "dao/UpdateVote", nil)
	cdc.RegisterConcrete(&MsgDeleteVote{}, "dao/DeleteVote", nil)
	cdc.RegisterConcrete(&MsgCreateProposal{}, "dao/CreateProposal", nil)
	cdc.RegisterConcrete(&MsgUpdateProposal{}, "dao/UpdateProposal", nil)
	cdc.RegisterConcrete(&MsgDeleteProposal{}, "dao/DeleteProposal", nil)
	cdc.RegisterConcrete(&MsgCreateVoter{}, "dao/CreateVoter", nil)
	cdc.RegisterConcrete(&MsgUpdateVoter{}, "dao/UpdateVoter", nil)
	cdc.RegisterConcrete(&MsgDeleteVoter{}, "dao/DeleteVoter", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWarehouse{},
		&MsgUpdateWarehouse{},
		&MsgDeleteWarehouse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVote{},
		&MsgUpdateVote{},
		&MsgDeleteVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProposal{},
		&MsgUpdateProposal{},
		&MsgDeleteProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVoter{},
		&MsgUpdateVoter{},
		&MsgDeleteVoter{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
