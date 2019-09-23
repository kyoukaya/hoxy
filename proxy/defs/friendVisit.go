package defs

// CFriend/visit
type CFriendVisit struct {
	FUserid int `json:"f_userid"`
}

type SFriendVisit struct {
	Info struct {
		VisitNum     int    `json:"visit_num"`
		PraiseNum    int    `json:"praise_num"`
		UserID       int    `json:"user_id"`
		Lv           int    `json:"lv"`
		FUserid      int    `json:"f_userid"`
		Name         string `json:"name"`
		IsPraise     int    `json:"is_praise"`
		HeadpicID    int    `json:"headpic_id"`
		HomepageTime int    `json:"homepage_time"`
	} `json:"info"`
	AdjutantList []struct {
		Adjutant *struct {
			GunID int `json:"gun_id"`
			Skin  int `json:"skin"`
			Mod   int `json:"mod"`
			Ai    int `json:"ai"`
		} `json:"adjutant"`
		FUserid       int         `json:"f_userid"`
		Name          string      `json:"name"`
		Lv            int         `json:"lv"`
		HeadpicID     int         `json:"headpic_id"`
		HomepageTime  int         `json:"homepage_time"`
		Comment       interface{} `json:"comment"`
		AdjutantFairy *struct {
			FairyID string `json:"fairy_id"`
			Skin    string `json:"skin"`
		} `json:"adjutant_fairy"`
	} `json:"adjutant_list"`
	Notice struct {
		IsViewNotice string `json:"is_view_notice"`
		Notice       string `json:"notice"`
	} `json:"notice"`
	GunWithUserList []struct {
		GunID int `json:"gun_id"`
		Skin  int `json:"skin"`
	} `json:"gun_with_user_list"`
	FurnitureList []struct {
		ID          int `json:"id"`
		UserID      int `json:"user_id"`
		Dorm        int `json:"dorm"`
		Type        int `json:"type"`
		FurnitureID int `json:"furniture_id"`
		X           int `json:"x"`
		Y           int `json:"y"`
		ClassesID   int `json:"classes_id"`
	} `json:"furniture_list"`
	BuildCoinFlag int `json:"build_coin_flag"`
}
