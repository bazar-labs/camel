package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/el-goblino-foundation/turron/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IInventoryRegistry interface {
	GetItemDefinition(itemDefID *big.Int) (string, error)
	CreateItemDefinition(itemDefURI string) (*types.Transaction, error)
}

type InventoryRegistry struct {
	client  *ethclient.Client
	pk      *ecdsa.PrivateKey
	address common.Address
}

func (c *Client) InventoryRegistry(address common.Address) IInventoryRegistry {
	return &InventoryRegistry{c.client, c.pk, address}
}

func (c *InventoryRegistry) GetItemDefinition(itemDefID *big.Int) (string, error) {
	instance, err := contract.NewInventoryRegistryContract(c.address, c.client)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate 'InventoryRegistry' contract: %v", err)
	}

	itemDefURI, err := instance.DefinitionIDToURI(&bind.CallOpts{}, itemDefID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve the 'number' variable: %v", err)
	}

	// TODO also fetch token ids

	return itemDefURI, nil
}

func (c *InventoryRegistry) CreateItemDefinition(itemDefURI string) (*types.Transaction, error) {
	instance, err := contract.NewInventoryRegistryContract(c.address, c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate 'InventoryRegistry' contract: %v", err)
	}

	tx, err := instance.CreateItemDefinition(&bind.TransactOpts{}, itemDefURI)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the 'number' variable: %v", err)
	}

	return tx, nil
}
