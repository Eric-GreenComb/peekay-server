package ethereum

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Eric-GreenComb/eth-server/config"
)

// Clients global client connection
var Clients = &Conn{}

// Conn 各种链的链接走这里
type Conn struct {
	Eth *rpc.Client
}

// Init 初始化各种链接
// TODO 维护多个连接池（每个链考虑多个节点)
func Init() error {
	fmt.Println(">>>>>>>>> start dial ethereum >>>>>>>>>>>")
	eth, err := rpc.Dial(config.EthereumConfig.Host)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(">>>>>>>>> dialed ethereum >>>>>>>>>>>")

	Clients.Eth = eth

	return nil
}
