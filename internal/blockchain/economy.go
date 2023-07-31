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
var MasterBoringFactoryContractAddress = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
var MasterInventoryRegistryContractAddress = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
var OwnerAddress = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

// func (c *Economy) Deploy(contracts strings...) (domain.GameContractAddresses, error) {
func (c *Economy) Deploy() (domain.GameContractAddresses, error) {
	instance, err := contract.NewBoringFactoryContract(MasterBoringFactoryContractAddress, c.client)
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
