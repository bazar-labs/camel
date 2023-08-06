package domain

import "math/big"

type ItemDefinition struct {
	ID  *big.Int `json:"id"`
	URI string   `json:"uri"`
}
