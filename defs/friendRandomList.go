package defs


// SFriend/randomList
type SFriendRandomList struct {
	List []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int    `json:"f_userid"`
		Name         string `json:"name"`
		Lv           string `json:"lv"`
		HeadpicID    string `json:"headpic_id"`
		HomepageTime int    `json:"homepage_time"`
		IsReturnUser int    `json:"is_return_user"`
	} `json:"list"`
}
