package domain

import "math/big"

type Item struct {
	DefID *big.Int `json:"def_id"`
	URI   string   `json:"uri"`
}
