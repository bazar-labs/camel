package domain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Game struct {
	ID                uuid.UUID              `json:"id"`
	Name              string                 `json:"name"`
	ContractAddresses *GameContractAddresses `json:"contract_addresses"`
}

type GameContractAddresses struct {
	InventoryRegistry common.Address `json:"inventory_registry"`
}
