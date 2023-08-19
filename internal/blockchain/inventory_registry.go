package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/bazar-labs/turron/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IInventoryRegistry interface {
	GetItem(itemDefID *big.Int) (string, error)
	CreateItem(itemDefURI string) (*big.Int, error)
}

type InventoryRegistry struct {
	client  *ethclient.Client
	pk      *ecdsa.PrivateKey
	address common.Address
}

func (c *Client) InventoryRegistry(address common.Address) IInventoryRegistry {
	return &InventoryRegistry{c.client, c.pk, address}
}

func (c *InventoryRegistry) GetItem(itemDefID *big.Int) (string, error) {
	instance, err := contract.NewInventoryRegistryContract(c.address, c.client)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate 'InventoryRegistry' contract: %v", err)
	}

	itemDefURI, err := instance.ItemDefinitionIDToURI(&bind.CallOpts{}, itemDefID)
	if err != nil {
		return "", fmt.Errorf("failed to call 'DefinitionIDToURI' function: %w", err)
	}

	// TODO also fetch token ids

	return itemDefURI, nil
}

func (c *InventoryRegistry) CreateItem(itemDefURI string) (*big.Int, error) {
	instance, err := contract.NewInventoryRegistryContract(c.address, c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate 'InventoryRegistry' contract: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.pk, big.NewInt(CHAIN_ID))
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}
	// FIXME
	auth.GasLimit = uint64(30000000)

	tx, err := instance.Create(auth, itemDefURI)
	if err != nil {
		return nil, fmt.Errorf("failed to create 'CreateItem' transaction: %w", err)
	}

	_, err = bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for contract deployment transaction to be mined: %w", err)
	}

	receipt, err := c.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contract.InventoryRegistryContractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	var itemDefID *big.Int
	eventSignature := parsedABI.Events["ItemDefinitionCreated"].ID
	for _, log := range receipt.Logs {
		log := log
		if log.Topics[0] == eventSignature {
			itemDefID = log.Topics[1].Big()
		}
	}

	return itemDefID, nil
}
