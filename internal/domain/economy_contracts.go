package domain

import "github.com/ethereum/go-ethereum/common"

type MasterEconomyContracts struct {
	BoringFactory       common.Address        `json:"boring_factory"`
	InventoryRegistry   common.Address        `json:"inventory_registry"`
	InventoryController common.Address        `json:"inventory_controller"`
	Behaviors           BehaviourContractInfo `json:"behaviors"`
}

type GameEconomyContracts struct {
	InventoryRegistry   ContractInfo          `json:"inventory_registry"`
	InventoryController ContractInfo          `json:"inventory_controller"`
	Behaviors           BehaviourContractInfo `json:"behaviors"`
}

type BehaviourContractInfo struct {
	PurchaseItemWithETH ContractInfo `json:"purchase_item_with_eth"`
}

type ContractInfo struct {
	Address common.Address `json:"address"`
}
