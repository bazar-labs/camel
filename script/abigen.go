package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

// Define the structure of JSON data
type contract struct {
	ABI []interface{} `json:"abi"`
}

const outDir = "contract"
const artifactDir = "artifacts"

func abigen(name string, abi []interface{}) {
	abiData, err := json.Marshal(abi)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(artifactDir+"/"+name+"_abi.json", abiData, 0644)
	if err != nil {
		panic(err)
	}

	err = exec.
		Command("abigen", "--abi="+artifactDir+"/"+name+"_abi.json", "--pkg=contract", "--out="+outDir+"/"+name+".go", "--type="+name+"Contract").
		Run()

	if err != nil {
		panic(err)
	}
}

func tempDir(dir string) func() {
	err := os.Mkdir(artifactDir, 0755)
	if err != nil {
		panic(err)
	}

	return func() {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			err := os.RemoveAll(dir)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	names := []string{"Counter"}

	cleanup := tempDir(artifactDir)
	defer cleanup()

	for _, name := range names {
		file, err := os.ReadFile("../protocol/out/" + name + ".sol/" + name + ".json")
		if err != nil {
			panic(err)
		}
		var contract contract
		json.Unmarshal(file, &contract)
		abigen(name, contract.ABI)
	}
}
