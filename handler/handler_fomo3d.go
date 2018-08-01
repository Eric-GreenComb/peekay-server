package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/peekay-server/bean"
	"github.com/Eric-GreenComb/peekay-server/ethereum"
)

// GetPlayerInfoByAddress GetPlayerInfoByAddress
func GetPlayerInfoByAddress(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conAddr := _formParams.ConAddress
	_addr := _formParams.Address

	ret, err := ethereum.GetPlayerInfoByAddress(_conAddr, _addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, ret)

	_fomoPlayerInfo, err := ethereum.ParseFomoPlayerInfo(ret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _fomoPlayerInfo)
}

// GetTimeLeft GetTimeLeft
func GetTimeLeft(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_conAddr := _formParams.ConAddress

	ret, err := ethereum.GetTimeLeft(_conAddr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_base16 := strings.TrimLeft(strings.TrimPrefix(ret, "0x"), "0")

	_base10, err := strconv.ParseUint(_base16, 16, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _base10})
}
