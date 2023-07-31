package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const OutDir = "contract"
const TempDir = "temp"

var Names = []string{"Counter", "BoringFactory"}

type contract struct {
	ABI []interface{} `json:"abi"`
}

func main() {
	remove := dir(TempDir)
	defer remove()

	for _, name := range Names {
		file, err := os.ReadFile("../protocol/out/" + name + ".sol/" + name + ".json")
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		var contract contract
		err = json.Unmarshal(file, &contract)
		if err != nil {
			log.Fatalf("failed to unmarshal json: %v", err)
		}

		err = abigen(name, contract.ABI)
		if err != nil {
			log.Fatalf("failed to abigen: %v", err)
		}
	}
}

func abigen(name string, abi []interface{}) error {
	data, err := json.Marshal(abi)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}

	err = os.WriteFile(TempDir+"/"+name+"_abi.json", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	args := []string{"--abi=" + TempDir + "/" + name + "_abi.json", "--pkg=contract", "--out=" + OutDir + "/" + name + ".go", "--type=" + name + "Contract"}

	err = exec.Command("abigen", args...).Run()
	if err != nil {
		return fmt.Errorf("failed to run abigen command: %v", err)
	}

	return nil
}

func dir(dir string) func() {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatalf("failed to create dir: %v", err)
	}

	return func() {
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatalf("failed to remove dir: %v", err)
		}
	}
}
