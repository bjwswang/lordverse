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
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		WarehouseList: []Warehouse{},
		VoteList:      []Vote{},
		ProposalList:  []Proposal{},
		VoterList:     []Voter{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in warehouse
	warehouseIdMap := make(map[uint64]bool)
	warehouseCount := gs.GetWarehouseCount()
	for _, elem := range gs.WarehouseList {
		if _, ok := warehouseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for warehouse")
		}
		if elem.Id >= warehouseCount {
			return fmt.Errorf("warehouse id should be lower or equal than the last id")
		}
		warehouseIdMap[elem.Id] = true
	}
	// Check for duplicated ID in vote
	voteIdMap := make(map[uint64]bool)
	voteCount := gs.GetVoteCount()
	for _, elem := range gs.VoteList {
		if _, ok := voteIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vote")
		}
		if elem.Id >= voteCount {
			return fmt.Errorf("vote id should be lower or equal than the last id")
		}
		voteIdMap[elem.Id] = true
	}
	// Check for duplicated ID in proposal
	proposalIdMap := make(map[uint64]bool)
	proposalCount := gs.GetProposalCount()
	for _, elem := range gs.ProposalList {
		if _, ok := proposalIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for proposal")
		}
		if elem.Id >= proposalCount {
			return fmt.Errorf("proposal id should be lower or equal than the last id")
		}
		proposalIdMap[elem.Id] = true
	}
	// Check for duplicated ID in voter
	voterIdMap := make(map[uint64]bool)
	voterCount := gs.GetVoterCount()
	for _, elem := range gs.VoterList {
		if _, ok := voterIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for voter")
		}
		if elem.Id >= voterCount {
			return fmt.Errorf("voter id should be lower or equal than the last id")
		}
		voterIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
