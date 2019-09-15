package defs

// SFriend/dormInfo
type SFriendDormInfo struct {
	Info struct {
		PraiseNum int    `json:"praise_num"`
		VisitNum  int    `json:"visit_num"`
		UserID    string `json:"user_id"`
		DormID    string `json:"dorm_id"`
	} `json:"info"`
	InMydormList []struct {
		Adjutant struct {
			GunID string `json:"gun_id"`
			Skin  string `json:"skin"`
			Mod   string `json:"mod"`
			Ai    string `json:"ai"`
		} `json:"adjutant"`
		FUserid      int         `json:"f_userid"`
		Name         string      `json:"name"`
		Lv           string      `json:"lv"`
		HeadpicID    interface{} `json:"headpic_id"` // Can be string or int
		HomepageTime int         `json:"homepage_time"`
		Comment      interface{} `json:"comment"` // string or null
	} `json:"in_mydorm_list"`
	MyVisitList          []interface{} `json:"my_visit_list"`
	BuildCoinFlag        interface{}   `json:"build_coin_flag"` // string or int
	CurrentBuildCoin     string        `json:"current_build_coin"`
	EstablishBuildResult []struct {
		FurnitureID   interface{} `json:"furniture_id"`
		BuildCoin     string      `json:"build_coin"`
		Gem           string      `json:"gem"`
		EstablishType string      `json:"establish_type"`
		PetName       string      `json:"pet_name"`
		BuildTmpData  []string    `json:"build_tmp_data"`
	} `json:"establish_build_result"`
	ResetBuildCoinData []interface{} `json:"reset_build_coin_data"`
	Notice             struct {
		IsViewNotice string `json:"is_view_notice"`
		Notice       string `json:"notice"`
	} `json:"notice"`
}
