package bean

import (
	"math/big"
)

// FomoPlayerInfo FomoPlayerInfo
type FomoPlayerInfo struct {
	PlayerID       *big.Int `json:"id"`
	PlayerName     string   `json:"name"`
	KeysOwned      *big.Int `json:"keys"`
	WinningsVault  *big.Int `json:"win"`
	GeneralVault   *big.Int `json:"vault"`
	AffiliateVault *big.Int `json:"affi"`
	RoundEth       *big.Int `json:"eth"`
}
