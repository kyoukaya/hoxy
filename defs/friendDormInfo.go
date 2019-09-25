package defs

// SFriend/dormInfo
type SFriendDormInfo struct {
	Info struct {
		// PraiseNum and VisitNum can sometimes be swapped.
		PraiseNum int    `json:"praise_num"`
		VisitNum  int    `json:"visit_num"`
		UserID    int    `json:"user_id"`
		DormID    string `json:"dorm_id"`
	} `json:"info"`
	InMydormList []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int    `json:"f_userid"`
		Name         string `json:"name"`
		Lv           string `json:"lv"`
		HeadpicID    int    `json:"headpic_id"`
		HomepageTime int    `json:"homepage_time"`
		Comment      string `json:"comment"`
	} `json:"in_mydorm_list"`
	MyVisitList []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int    `json:"f_userid"`
		Name         string `json:"name"`
		Lv           string `json:"lv"`
		HeadpicID    int    `json:"headpic_id"`
		HomepageTime int    `json:"homepage_time"`
		CreateTime   int    `json:"create_time"`
	} `json:"my_visit_list"`
	BuildCoinFlag        int    `json:"build_coin_flag"`
	CurrentBuildCoin     string `json:"current_build_coin"`
	EstablishBuildResult []struct {
		FurnitureID   int    `json:"furniture_id"`
		BuildCoin     string `json:"build_coin"`
		Gem           string `json:"gem"`
		EstablishType string `json:"establish_type"`
		PetName       string `json:"pet_name"`
		BuildTmpData  []int  `json:"build_tmp_data"`
	} `json:"establish_build_result"`
	ResetBuildCoinData *struct {
		MaxBuildCoin int `json:"max_build_coin"`
	} `json:"reset_build_coin_data"`
	Notice struct {
		IsViewNotice string `json:"is_view_notice"`
		Notice       string `json:"notice"`
	} `json:"notice"`
}
