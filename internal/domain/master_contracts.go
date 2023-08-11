package domain

import "github.com/ethereum/go-ethereum/common"

type MasterContracts struct {
	BoringFactory        common.Address `json:"boring_factory"`
	InventoryRegistry    common.Address `json:"inventory_registry"`
	InventoryController  common.Address `json:"inventory_controller"`
	BehaviorItemPurchase common.Address `json:"behavior_item_purchase"`
}
