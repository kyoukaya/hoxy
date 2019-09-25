package defs

// COuthouse/establish_build
type COuthouseEstablish_build struct {
	EstablishType int    `json:"establish_type"`
	Num           int    `json:"num"`
	Payway        string `json:"payway"`
}

type SOuthouseEstablish_build struct {
	BuildCoin  int `json:"build_coin"`
	GiftItemID int `json:"gift_item_id"`
	Exp        int `json:"exp"`
	BuildNum   int `json:"build_num"`
	// BuildTmpData is a mixed type array, e.g., [80,200001,240000,"build_coin",240]
	BuildTmpData []string `json:"build_tmp_data"`
}
