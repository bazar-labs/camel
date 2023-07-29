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

type ICounter interface {
	GetNumber() (*big.Int, error)
	SetNumber(number int) (*types.Transaction, error)
}

type Counter struct {
	client  *ethclient.Client
	pk      *ecdsa.PrivateKey
	address common.Address
}

func (c *Client) Counter(address common.Address) ICounter {
	return &Counter{c.client, c.pk, address}
}

func (c *Counter) GetNumber() (*big.Int, error) {
	instance, err := contract.NewCounterContract(c.address, c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate 'Counter' contract: %v", err)
	}

	number, err := instance.Number(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the 'number' variable: %v", err)
	}

	return number, nil
}

func (c *Counter) SetNumber(number int) (*types.Transaction, error) {
	instance, err := contract.NewCounterContract(c.address, c.client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate 'Counter' contract: %v", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(c.pk, new(big.Int).SetInt64(31337))
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	tx, err := instance.SetNumber(transactOpts, big.NewInt(int64(number)))
	if err != nil {
		return nil, fmt.Errorf("failed to increment the number: %v", err)
	}

	return tx, nil
}
