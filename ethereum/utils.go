package ethereum

import (
	"encoding/hex"
	// "errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Eric-GreenComb/peekay-server/bean"
)

// TrimLeftZeros TrimLeftZeros
func TrimLeftZeros(str string) string {
	return strings.TrimLeft(strings.TrimPrefix(str, "0x"), "0")
}

// ParseFomoPlayerInfo ParseFomoPlayerInfo
func ParseFomoPlayerInfo(str string) (bean.FomoPlayerInfo, error) {
	var _fomoPlayerInfo bean.FomoPlayerInfo

	_str := strings.TrimPrefix(str, "0x")
	fmt.Println(len(_str))
	// if len(str) != 448 {
	// 	return _fomoPlayerInfo, errors.New("input len is not 448")
	// }

	_data1 := _str[0:64]
	bigInt1 := new(big.Int)
	bigInt1.SetString(_data1, 16)
	_fomoPlayerInfo.PlayerID = bigInt1

	_data2 := _str[64:128]
	_hex2byte, _ := hex.DecodeString(_data2)
	_fomoPlayerInfo.PlayerName = strings.Trim(string(_hex2byte), " ")

	_data3 := _str[128:192]
	bigInt3 := new(big.Int)
	bigInt3.SetString(_data3, 16)
	_fomoPlayerInfo.KeysOwned = bigInt3

	_data4 := _str[192:256]
	bigInt4 := new(big.Int)
	bigInt4.SetString(_data4, 16)
	_fomoPlayerInfo.WinningsVault = bigInt4

	_data5 := _str[256:320]
	bigInt5 := new(big.Int)
	bigInt5.SetString(_data5, 16)
	_fomoPlayerInfo.GeneralVault = bigInt5

	_data6 := _str[320:384]
	bigInt6 := new(big.Int)
	bigInt6.SetString(_data6, 16)
	_fomoPlayerInfo.AffiliateVault = bigInt6

	_data7 := _str[384:448]
	bigInt7 := new(big.Int)
	bigInt7.SetString(_data7, 16)
	_fomoPlayerInfo.RoundEth = bigInt7

	return _fomoPlayerInfo, nil
}
