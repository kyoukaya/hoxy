package defs

// Gun/developLogCollect
type CGunDevelopLogCollect struct {
	DevUID  int `json:"dev_uid"`
	LogID   int `json:"log_id"`
	LogInfo struct {
		DevType    int    `json:"dev_type"`
		Item1Num   int    `json:"item1_num"`
		Core       int    `json:"core"`
		InputLevel int    `json:"input_level"`
		DevUname   string `json:"dev_uname"`
		DevLv      int    `json:"dev_lv"`
		GunID      int    `json:"gun_id"`
		Mp         int    `json:"mp"`
		Ammo       int    `json:"ammo"`
		Mre        int    `json:"mre"`
		Part       int    `json:"part"`
		DevTime    int    `json:"dev_time"`
	} `json:"log_info"`
}

type SGunDevelopLogCollect int
