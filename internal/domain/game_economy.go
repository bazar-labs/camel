package domain

type GameEconomy struct {
	ID           int64                `json:"id"`
	GameID       int64                `json:"game_id"`
	ChainNetwork ChainNetwork         `json:"chain_network"`
	Contracts    GameEconomyContracts `json:"contracts"`
}
