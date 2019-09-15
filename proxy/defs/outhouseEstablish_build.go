package defs

// COuthouse/establish_build
type COuthouseEstablish_build struct {
	EstablishType int    `json:"establish_type"`
	Num           int    `json:"num"`
	Payway        string `json:"payway"`
}

type SOuthouseEstablish_build struct {
	BuildCoin    int           `json:"build_coin"`
	GiftItemID   int           `json:"gift_item_id"`
	Exp          int           `json:"exp"`
	BuildNum     int           `json:"build_num"`
	BuildTmpData []interface{} `json:"build_tmp_data"`
}
