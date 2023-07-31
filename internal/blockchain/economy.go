package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/el-goblino-foundation/turron/contract"
	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IEconomy interface {
	Deploy() (domain.GameContractAddresses, error)
}

type Economy struct {
	client *ethclient.Client
	pk     *ecdsa.PrivateKey
}

func (c *Client) Economy() IEconomy {
	return &Economy{c.client, c.pk}
}

// FIXME
var MasterFactoryContractAddress = common.HexToAddress("")
var MasterInventoryRegistryContractAddress = common.HexToAddress("")
var OwnerAddress = common.HexToAddress("")

// func (c *Economy) Deploy(contracts strings...) (domain.GameContractAddresses, error) {
func (c *Economy) Deploy() (domain.GameContractAddresses, error) {
	instance, err := contract.NewBoringFactoryContract(MasterFactoryContractAddress, c.client)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to instantiate a 'BoringFactory' contract: %w", err)
	}

	// loop over contracts

	abi, err := abi.JSON(strings.NewReader(contract.InventoryRegistryContractABI))
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to create ABI: %w", err)
	}

	data, err := abi.Pack("init", OwnerAddress)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to pack data: %w", err)
	}

	_, err = instance.Deploy(&bind.TransactOpts{}, MasterInventoryRegistryContractAddress, data, false)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to deploy with 'BoringFactory' contract: %w", err)
	}

	return domain.GameContractAddresses{}, nil
}
