package defs

// CIndexHome Client's request for Index/home
type CIndexHome struct {
	DataVersion string `json:"data_version"`
	AbVersion   string `json:"ab_version"`
	StartID     int    `json:"start_id"`
	IgnoreTime  int    `json:"ignore_time"`
}

// SIndexHome Server's response for Index/home
type SIndexHome struct {
	RecoverMp     int `json:"recover_mp"`
	RecoverAmmo   int `json:"recover_ammo"`
	RecoverMre    int `json:"recover_mre"`
	RecoverPart   int `json:"recover_part"`
	Gem           int `json:"gem"`
	AllFavorupGun []struct {
		GunWithUserID string      `json:"gun_with_user_id"`
		FavorAfteradd interface{} `json:"favor_afteradd"` // Can be a string or int
	} `json:"all_favorup_gun"`
	LastFavorRecoverTime interface{}   `json:"last_favor_recover_time"` // Can be a string or int
	RecoverSsoc          int           `json:"recover_ssoc"`
	LastSsocChangeTime   int           `json:"last_ssoc_change_time"`
	Kick                 int           `json:"kick"`
	FriendMessagelist    []interface{} `json:"friend_messagelist"`
	FriendApplylist      []struct {
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
	} `json:"friend_applylist"`
	IndexGetmaillist []struct {
		ID        string `json:"id"`
		UserID    string `json:"user_id"`
		Type      string `json:"type"`
		SubID     string `json:"sub_id"`
		UserExp   string `json:"user_exp"`
		Mp        string `json:"mp"`
		Ammo      string `json:"ammo"`
		Mre       string `json:"mre"`
		Part      string `json:"part"`
		Core      string `json:"core"`
		Gem       string `json:"gem"`
		GunID     string `json:"gun_id"`
		FairyIds  string `json:"fairy_ids"`
		ItemIds   string `json:"item_ids"`
		EquipIds  string `json:"equip_ids"`
		Furniture string `json:"furniture"`
		Gift      string `json:"gift"`
		Coins     string `json:"coins"`
		Skin      string `json:"skin"`
		BpPay     string `json:"bp_pay"`
		Chip      string `json:"chip"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Code      string `json:"code"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		IfRead    string `json:"if_read"`
	} `json:"index_getmaillist"`
	SquadDataDaily []struct {
		UserID         int    `json:"user_id"`
		SquadID        string `json:"squad_id"`
		Type           string `json:"type"`
		LastFinishTime string `json:"last_finish_time"`
		Count          int    `json:"count"`
		Receive        int    `json:"receive"`
	} `json:"squad_data_daily"`
	BuildCoinFlag int    `json:"build_coin_flag"`
	IsBind        string `json:"is_bind"`
}
