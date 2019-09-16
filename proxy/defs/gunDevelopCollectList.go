package defs

import "hoxy/proxy/defs/types"

// Gun/developCollectList
type SGunDevelopCollectList struct {
	LogList []struct {
		ID          string      `json:"id"`
		DevType     string      `json:"dev_type"`
		LogID       string      `json:"log_id"`
		UserID      string      `json:"user_id"`
		DevUID      interface{} `json:"dev_uid"`
		DevUname    types.Str   `json:"dev_uname"`
		DevLv       string      `json:"dev_lv"`
		GunID       string      `json:"gun_id"`
		Mp          string      `json:"mp"`
		Ammo        string      `json:"ammo"`
		Mre         string      `json:"mre"`
		Part        string      `json:"part"`
		InputLevel  string      `json:"input_level"`
		Item1Num    string      `json:"item1_num"`
		Core        string      `json:"core"`
		DevTime     string      `json:"dev_time"`
		CollectTime string      `json:"collect_time"`
	} `json:"log_list"`
}
