package defs

// SFriend/list
type SFriendList struct {
	List []struct {
		Adjutant *struct {
			GunID int `json:"gun_id"`
			Skin  int `json:"skin"`
			Mod   int `json:"mod"`
			Ai    int `json:"ai"`
		} `json:"adjutant"`
		FUserid       int    `json:"f_userid"`
		EndTime       int    `json:"end_time"`
		Name          string `json:"name"`
		Lv            int    `json:"lv"`
		HeadpicID     int    `json:"headpic_id"`
		HomepageTime  int    `json:"homepage_time"`
		IsReturnUser  int    `json:"is_return_user"`
		AdjutantFairy *struct {
			FairyID int `json:"fairy_id"`
			Skin    int `json:"skin"`
		} `json:"adjutant_fairy"`
	} `json:"list"`
}
