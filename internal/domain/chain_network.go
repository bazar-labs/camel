package domain

type ChainNetwork string

const (
	Localhost       ChainNetwork = "localhost"
	EthereumMainnet ChainNetwork = "ethereum_mainnet"
	EthereumGoerli  ChainNetwork = "ethereum_goerli"
	PolygonMainnet  ChainNetwork = "polygon_mainnet"
	PolygonMumbai   ChainNetwork = "polygon_mumbai"
)
