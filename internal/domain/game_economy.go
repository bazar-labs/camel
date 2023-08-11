package domain

import "github.com/ethereum/go-ethereum/common"

type GameEconomy struct {
	ID                int64                        `json:"id"`
	GameID            int64                        `json:"game_id"`
	ChainNetwork      ChainNetwork                 `json:"chain_network"`
	ContractAddresses GameEconomyContractAddresses `json:"contract_addresses"`
}

type GameEconomyContractAddresses struct {
	InventoryRegistry    common.Address `json:"inventory_registry"`
	InventoryController  common.Address `json:"inventory_controller"`
	BehaviorItemPurchase common.Address `json:"behavior_item_purchase"`
}
