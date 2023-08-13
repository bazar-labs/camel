package domain

import "github.com/ethereum/go-ethereum/common"

type MasterEconomyContracts struct {
	BoringFactory       common.Address  `json:"boring_factory"`
	InventoryRegistry   common.Address  `json:"inventory_registry"`
	InventoryController common.Address  `json:"inventory_controller"`
	Behaviors           MasterBehaviors `json:"behaviors"`
}

type GameEconomyContracts struct {
	InventoryRegistry   ContractInfo  `json:"inventory_registry"`
	InventoryController ContractInfo  `json:"inventory_controller"`
	Behaviors           GameBehaviors `json:"behaviors"`
}

type MasterBehaviors struct {
	PurchaseItem ContractInfo `json:"purchase_item"`
}

type GameBehaviors struct {
	PurchaseItem BehaviorContractInfo `json:"purchase_item"`
}

type ContractInfo struct {
	Address common.Address `json:"address"`
}

type BehaviorContractInfo struct {
	Address common.Address `json:"address"`
	Enabled bool           `json:"enabled"`
}
