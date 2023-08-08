package domain

type ChainNetworkID int

// TODO think about this being int instead of string
const (
	Localhost ChainNetworkID = iota
	EthereumMainnet
	EthereumGoerli
	PolygonMainnet
	PolygonMumbai
)
