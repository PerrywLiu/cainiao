package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OrdersList: []Orders{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in orders
	ordersIdMap := make(map[uint64]bool)
	ordersCount := gs.GetOrdersCount()
	for _, elem := range gs.OrdersList {
		if _, ok := ordersIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for orders")
		}
		if elem.Id >= ordersCount {
			return fmt.Errorf("orders id should be lower or equal than the last id")
		}
		ordersIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
