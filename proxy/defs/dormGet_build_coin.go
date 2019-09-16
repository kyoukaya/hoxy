package defs

import "hoxy/proxy/defs/types"

// Dorm/get_build_coin
type CDormGet_build_coin struct {
	VUserID string `json:"v_user_id"`
	DormID  int    `json:"dorm_id"`
}

type SDormGet_build_coin struct {
	BuildCoin types.Int `json:"build_coin"`
}
