package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LandList: []Land{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in land
	landIdMap := make(map[uint64]bool)
	landCount := gs.GetLandCount()
	for _, elem := range gs.LandList {
		if _, ok := landIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for land")
		}
		if elem.Id >= landCount {
			return fmt.Errorf("land id should be lower or equal than the last id")
		}
		landIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
