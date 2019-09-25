package defs

// CGun/developLog
type CGunDevelopLog struct {
	DevType int `json:"dev_type"`
}

// SGun/developLog
type SGunDevelopLog struct {
	LogList []struct {
		ID         string `json:"id"`
		DevType    string `json:"dev_type"`
		UserID     string `json:"user_id"`
		BuildSlot  string `json:"build_slot"`
		DevUname   string `json:"dev_uname"`
		DevLv      string `json:"dev_lv"`
		GunID      string `json:"gun_id"`
		Mp         string `json:"mp"`
		Ammo       string `json:"ammo"`
		Mre        string `json:"mre"`
		Part       string `json:"part"`
		InputLevel string `json:"input_level"`
		Item1Num   string `json:"item1_num"`
		Core       string `json:"core"`
		DevTime    string `json:"dev_time"`
	} `json:"log_list"`
}
