package defs

// SFriend/list
type SFriendList struct {
	List []struct {
		Adjutant *struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid       int    `json:"f_userid"`
		EndTime       int    `json:"end_time"`
		Name          string `json:"name"`
		Lv            string `json:"lv"`
		HeadpicID     string `json:"headpic_id"`
		HomepageTime  int    `json:"homepage_time`
		IsReturnUser  int    `json:"is_return_user"`
		AdjutantFairy *struct {
			FairyID string `json:"fairy_id"`
			Skin    string `json:"skin"`
		} `json:"adjutant_fairy"`
	} `json:"list"`
}
