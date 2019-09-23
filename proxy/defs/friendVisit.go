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
		Lv           string `json:"lv"`
		FUserid      int    `json:"f_userid"`
		Name         string `json:"name"`
		IsPraise     int    `json:"is_praise"`
		HeadpicID    int    `json:"headpic_id"`
		HomepageTime int    `json:"homepage_time"`
	} `json:"info"`
	AdjutantList []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int         `json:"f_userid"`
		Name         string      `json:"name"`
		Lv           string      `json:"lv"`
		HeadpicID    int         `json:"headpic_id"`
		HomepageTime int         `json:"homepage_time"`
		Comment      interface{} `json:"comment"`
	} `json:"adjutant_list"`
	Notice struct {
		IsViewNotice string `json:"is_view_notice"`
		Notice       string `json:"notice"`
	} `json:"notice"`
	GunWithUserList []struct {
		GunID string `json:"gun_id"`
		Skin  string `json:"skin"`
	} `json:"gun_with_user_list"`
	FurnitureList []struct {
		ID          string `json:"id"`
		UserID      string `json:"user_id"`
		Dorm        string `json:"dorm"`
		Type        string `json:"type"`
		FurnitureID string `json:"furniture_id"`
		X           string `json:"x"`
		Y           string `json:"y"`
		ClassesID   string `json:"classes_id"`
	} `json:"furniture_list"`
	BuildCoinFlag int `json:"build_coin_flag"`
}
