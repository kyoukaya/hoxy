package defs

// SFriend/list
type SFriendList struct {
	List []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant,omitempty"`
		FUserid       int         `json:"f_userid"`
		EndTime       int         `json:"end_time"`
		Name          string      `json:"name"`
		Lv            string      `json:"lv"`
		HeadpicID     string      `json:"headpic_id"`
		HomepageTime  interface{} `json:"homepage_time"` // string or int
		IsReturnUser  int         `json:"is_return_user"`
		AdjutantFairy struct {
			FairyID string `json:"fairy_id"`
			Skin    string `json:"skin"`
		} `json:"adjutant_fairy,omitempty"`
	} `json:"list"`
}
