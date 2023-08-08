package domain

import (
	"github.com/ethereum/go-ethereum/common"
)

type Game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ContractAddresses *GameContractAddresses `json:"contract_addresses"`
type GameContractAddresses struct {
	InventoryRegistry common.Address `json:"inventory_registry"`
}
