package bean

// ServerConfig ServerConfig Struct
type ServerConfig struct {
	Port []string
	Mode string
}

// EthereumConfig EthereumConfig
type EthereumConfig struct {
	ChainID int64
	Host    string
}
