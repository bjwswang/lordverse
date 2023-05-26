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

const (
	// ModuleName defines the module name
	ModuleName = "dao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dao"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	WarehouseKey      = "Warehouse/value/"
	WarehouseCountKey = "Warehouse/count/"
)

const (
	VoteKey      = "Vote/value/"
	VoteCountKey = "Vote/count/"
)

const (
	ProposalKey      = "Proposal/value/"
	ProposalCountKey = "Proposal/count/"
)

const (
	VoterKey      = "Voter/value/"
	VoterCountKey = "Voter/count/"
)
