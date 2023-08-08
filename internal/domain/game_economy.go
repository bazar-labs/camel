package domain

import "github.com/ethereum/go-ethereum/common"

type GameEconomy struct {
	ID                int64                        `json:"id"`
	GameID            int64                        `json:"game_id"`
	ChainNetworkID    ChainNetworkID               `json:"chain_network_id"`
	ContractAddresses GameEconomyContractAddresses `json:"contract_addresses"`
}

type GameEconomyContractAddresses struct {
	InventoryRegistry common.Address `json:"inventory_registry"`
}
