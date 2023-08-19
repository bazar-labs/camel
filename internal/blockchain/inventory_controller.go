package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/bazar-labs/camel/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IInventoryController interface {
	IsInventoryBehavior(behavior common.Address) (bool, error)
	SetBehavior(ctx context.Context, behavior common.Address, state bool) error
}

type InventoryController struct {
	client  *ethclient.Client
	pk      *ecdsa.PrivateKey
	address common.Address
}

func (c *Client) InventoryController(address common.Address) IInventoryController {
	return &InventoryController{c.client, c.pk, address}
}

func (c *InventoryController) IsInventoryBehavior(behavior common.Address) (bool, error) {
	instance, err := contract.NewInventoryControllerContract(c.address, c.client)
	if err != nil {
		return false, fmt.Errorf("failed to instantiate 'InventoryController' contract: %v", err)
	}

	state, err := instance.IsInventoryBehavior(&bind.CallOpts{}, behavior)
	if err != nil {
		return false, fmt.Errorf("failed to get behavior state: %v", err)
	}

	return state, nil
}

func (c *InventoryController) SetBehavior(ctx context.Context, behavior common.Address, state bool) error {
	instance, err := contract.NewInventoryControllerContract(c.address, c.client)
	if err != nil {
		return fmt.Errorf("failed to instantiate 'InventoryController' contract: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.pk, big.NewInt(CHAIN_ID))
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}
	// FIXME
	auth.GasLimit = uint64(30000000)

	tx, err := instance.SetBehavior(auth, behavior, state)
	if err != nil {
		return fmt.Errorf("failed to create 'SetBehavior' transaction: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for contract deployment transaction to be mined: %w", err)
	}

	return nil
}
