package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/eth-server/bean"
)

// EthereumConfig Ethereum Config
var EthereumConfig bean.EthereumConfig

// ServerConfig Server Config
var ServerConfig bean.ServerConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	ServerConfig.Port = strings.Split(viper.GetString("server.port"), ",")
	ServerConfig.Mode = viper.GetString("server.mode")
	ServerConfig.GormLogMode = viper.GetString("server.gorm.LogMode")
	ServerConfig.ViewLimit = viper.GetInt("server.view.limit")

	EthereumConfig.ChainID = viper.GetInt64("ethereum.chainID")
	EthereumConfig.Host = viper.GetString("ethereum.host")
	EthereumConfig.Passphrase = viper.GetString("ethereum.passphrase")
}
