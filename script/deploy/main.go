package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const ANVIL_RPC_URL = "http://127.0.0.1:8545"
const ANVIL_PUBLIC_KEY = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const ANVIL_PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

type DeployArg struct {
	Path            string
	Name            string
	ConstructorArgs []string
	DeployTo        common.Address
}

var DEPLOY_ARGS = []DeployArg{
	{
		Path:            "lib/BoringSolidity/contracts/BoringFactory.sol",
		Name:            "BoringFactory",
		ConstructorArgs: []string{ANVIL_PUBLIC_KEY},
	},
	{
		Path:            "src/InventoryRegistry.sol",
		Name:            "InventoryRegistry",
		ConstructorArgs: []string{ANVIL_PUBLIC_KEY},
	},
}

func main() {
	// cfg := config.Load()
	// db := db.New(&cfg.Database)

	for i, arg := range DEPLOY_ARGS {
		contract := strings.Join([]string{arg.Path, arg.Name}, ":")

		command := exec.Command("forge", "create",
			contract,
			"--rpc-url", ANVIL_RPC_URL,
			"--private-key", ANVIL_PRIVATE_KEY,
			"--constructor-args", strings.Join(arg.ConstructorArgs, " "),
		)

		command.Dir = "protocol"

		output, err := command.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to deploy %s: %v\n%s", contract, err, string(output))
		}

		regex := regexp.MustCompile(`Deployed to: (0x[0-9a-fA-F]+)`)
		matches := regex.FindStringSubmatch(string(output))
		if len(matches) < 2 {
			log.Fatalf("Failed to parse deployed to address for %s: %s", contract, string(output))
		}

		DEPLOY_ARGS[i].DeployTo = common.HexToAddress(matches[1])
	}

	fmt.Printf("Deployed contracts: %v \n", DEPLOY_ARGS)
}
