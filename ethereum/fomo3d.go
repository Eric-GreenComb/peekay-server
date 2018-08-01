package ethereum

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetPlayerInfoByAddress 查询用户信息
func GetPlayerInfoByAddress(conAddr, addr string) (string, error) {
	_client := Clients.Eth

	_method := crypto.Keccak256([]byte("getPlayerInfoByAddress(address)"))[:4]

	_addr := PadLeft0(addr)

	_arguments, _ := hex.DecodeString(_addr)
	_input := append(_method, _arguments...)

	_to := ethcommon.HexToAddress(conAddr)

	var (
		msg = ethereum.CallMsg{To: &_to, Data: _input}
	)

	_arg := toCallArg(msg)

	var result string
	err := _client.Call(&result, "eth_call", _arg, "latest")
	if err != nil {
		return "", err
	}

	return result, nil
}

// GetTimeLeft GetTimeLeft
func GetTimeLeft(conAddr string) (string, error) {
	_client := Clients.Eth

	_method := crypto.Keccak256([]byte("getTimeLeft()"))[:4]
	_input := append(_method)

	_to := ethcommon.HexToAddress(conAddr)

	var (
		msg = ethereum.CallMsg{To: &_to, Data: _input}
	)
	_arg := toCallArg(msg)

	var result string
	err := _client.Call(&result, "eth_call", _arg, "latest")
	if err != nil {
		return "", err
	}

	return result, nil
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
