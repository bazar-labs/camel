package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/bazar-labs/camel/contract"
	"github.com/bazar-labs/camel/internal/domain"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// FIXME
const CHAIN_ID = 31337

type IEconomy interface {
	Deploy(domain.MasterEconomyContracts) (domain.GameEconomyContracts, error)
}

type Economy struct {
	client *ethclient.Client
	pk     *ecdsa.PrivateKey
}

type deployarg struct {
	Name    string
	Address common.Address
	Data    []byte
}

func (c *Client) Economy() IEconomy {
	return &Economy{c.client, c.pk}
}

func (c *Economy) Deploy(mc domain.MasterEconomyContracts) (domain.GameEconomyContracts, error) {
	addresses := domain.GameEconomyContracts{}

	instance, err := contract.NewBoringFactoryContract(mc.BoringFactory, c.client)
	if err != nil {
		return domain.GameEconomyContracts{}, fmt.Errorf("failed to instantiate a 'BoringFactory' contract: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.pk, big.NewInt(CHAIN_ID))
	if err != nil {
		return domain.GameEconomyContracts{}, fmt.Errorf("failed to create authorized transactor: %v", err)
	}
	auth.GasLimit = uint64(30000000)

	args := []deployarg{
		{
			Name:    "InventoryRegistry",
			Address: mc.InventoryRegistry,
			Data: packArgsTo32Bytes(
				crypto.PubkeyToAddress(c.pk.PublicKey).Bytes(),
			),
		},
		{
			Name:    "InventoryController",
			Address: mc.InventoryController,
			Data: packArgsTo32Bytes(
				crypto.PubkeyToAddress(c.pk.PublicKey).Bytes(),
			),
		},
		{
			Name:    "BehaviorPurchaseItem",
			Address: mc.Behaviors.PurchaseItemWithETH,
			Data: packArgsTo32Bytes(
				crypto.PubkeyToAddress(c.pk.PublicKey).Bytes(),
				addresses.InventoryRegistry.Bytes(),
				addresses.InventoryController.Bytes(),
			),
		},
	}

	for _, arg := range args {
		address, err := c.deploy(instance, auth, arg)
		if err != nil {
			return domain.GameEconomyContracts{}, err
		}

		switch arg.Name {
		case "InventoryRegistry":
			addresses.InventoryRegistry = address
		case "InventoryController":
			addresses.InventoryController = address
		case "BehaviorPurchaseItem":
			addresses.Behaviors.PurchaseItemWithETH = address
		}
	}

	return addresses, nil
}

func (c *Economy) deploy(instance *contract.BoringFactoryContract, auth *bind.TransactOpts, arg deployarg) (common.Address, error) {
	tx, err := instance.Deploy(auth, arg.Address, arg.Data, false)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy '%s' contract: %w", arg.Name, err)
	}

	_, err = bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to wait for contract deployment transaction to be mined: %w", err)
	}

	receipt, err := c.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contract.BoringFactoryContractABI))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	var address common.Address
	signature := parsedABI.Events["LogDeploy"].ID
	for _, log := range receipt.Logs {
		log := log
		if log.Topics[0] == signature {
			address = common.BytesToAddress(log.Topics[2].Bytes())
		}
	}

	return address, nil
}

// Pads each byte slice to a length of 32 bytes, then concatenates them all together.
func packArgsTo32Bytes(args ...[]byte) []byte {
	var result bytes.Buffer
	for _, arg := range args {
		padded := make([]byte, 32)      // Start with 32 zero bytes
		copy(padded[32-len(arg):], arg) // Overwrite the end of the padded slice with the arg bytes
		result.Write(padded)            // Add the padded bytes to the result
	}
	return result.Bytes()
}
