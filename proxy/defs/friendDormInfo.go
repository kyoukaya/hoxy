package defs

import (
	"hoxy/proxy/defs/types"
	"hoxy/proxy/defs/xdefs"
)

// SFriend/dormInfo
type SFriendDormInfo struct {
	Info struct {
		PraiseNum int       `json:"praise_num"`
		VisitNum  int       `json:"visit_num"`
		UserID    types.Int `json:"user_id"`
		DormID    string    `json:"dorm_id"`
	} `json:"info"`
	InMydormList []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int         `json:"f_userid"`
		Name         string      `json:"name"`
		Lv           string      `json:"lv"`
		HeadpicID    types.Int   `json:"headpic_id"`
		HomepageTime int         `json:"homepage_time"`
		Comment      interface{} `json:"comment"` // string or null
	} `json:"in_mydorm_list"`
	MyVisitList          []interface{} `json:"my_visit_list"`
	BuildCoinFlag        types.Int     `json:"build_coin_flag"`
	CurrentBuildCoin     string        `json:"current_build_coin"`
	EstablishBuildResult []struct {
		FurnitureID   types.Int   `json:"furniture_id"`
		BuildCoin     string      `json:"build_coin"`
		Gem           string      `json:"gem"`
		EstablishType string      `json:"establish_type"`
		PetName       string      `json:"pet_name"`
		BuildTmpData  []types.Int `json:"build_tmp_data"`
	} `json:"establish_build_result"`
	ResetBuildCoinData xdefs.ResetBuildCoinData `json:"reset_build_coin_data"`
	Notice             struct {
		IsViewNotice string `json:"is_view_notice"`
		Notice       string `json:"notice"`
	} `json:"notice"`
}
