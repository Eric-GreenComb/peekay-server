package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params     string `form:"params" json:"params"`     // params
	ConAddress string `form:"con_addr" json:"con_addr"` // ttl
	Address    string `form:"addr" json:"addr"`         //
}
