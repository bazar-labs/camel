package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/el-goblino-foundation/turron/contract"
	"github.com/el-goblino-foundation/turron/internal/config"
	"github.com/el-goblino-foundation/turron/internal/domain"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const CHAIN_ID = 31337

type IEconomy interface {
	Deploy() (domain.GameContractAddresses, error)
}

type Economy struct {
	client                  *ethclient.Client
	pk                      *ecdsa.PrivateKey
	masterContractAddresses *config.MasterContractAddresses
}

func (c *Client) Economy() IEconomy {
	return &Economy{c.client, c.pk, c.masterContractAddresses}
}

func (c *Economy) Deploy() (domain.GameContractAddresses, error) {
	instance, err := contract.NewBoringFactoryContract(c.masterContractAddresses.BoringFactory, c.client)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to instantiate a 'BoringFactory' contract: %w", err)
	}

	// fit owner into 32 bytes
	owner := crypto.PubkeyToAddress(c.pk.PublicKey).Bytes()
	data := make([]byte, 32)
	copy(data[32-len(owner):], owner)

	auth, err := bind.NewKeyedTransactorWithChainID(c.pk, big.NewInt(CHAIN_ID))
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to create authorized transactor: %v", err)
	}
	auth.GasLimit = uint64(30000000)

	tx, err := instance.Deploy(auth, c.masterContractAddresses.InventoryRegistry, data, false)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to deploy 'InventoryRegistry' contract: %w", err)
	}

	_, err = bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to wait for contract deployment transaction to be mined: %w", err)
	}

	receipt, err := c.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contract.BoringFactoryContractABI))
	if err != nil {
		return domain.GameContractAddresses{}, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	gca := domain.GameContractAddresses{}
	eventSignature := parsedABI.Events["LogDeploy"].ID
	for _, log := range receipt.Logs {
		if log.Topics[0] == eventSignature {
			gca.InventoryRegistry = common.BytesToAddress(log.Topics[2].Bytes())
		}
	}

	return gca, nil
}
