package defs

// CIndexQuest has no request body

// SIndexQuest Server's Index/Quest response
type SIndexQuest struct {
	Daily struct {
		DevelopGun          int         `json:"develop_gun"`
		Operation           interface{} `json:"operation"` // str or int
		ID                  string      `json:"id"`
		UserID              string      `json:"user_id"`
		EndTime             string      `json:"end_time"`
		Mission             string      `json:"mission"`
		Eat                 interface{} `json:"eat"` // str or int
		Fix                 string      `json:"fix"`
		Upgrade             string      `json:"upgrade"`
		CoinMission         string      `json:"coin_mission"`
		DevelopEquip        string      `json:"develop_equip"`
		WinRobot            string      `json:"win_robot"`
		WinPerson           string      `json:"win_person"`
		WinBoss             string      `json:"win_boss"`
		WinArmorrobot       string      `json:"win_armorrobot"`
		WinArmorperson      string      `json:"win_armorperson"`
		EatEquip            string      `json:"eat_equip"`
		FromFriendBuildCoin string      `json:"from_friend_build_coin"`
		BorrowFriendTeam    string      `json:"borrow_friend_team"`
		SquadDataAnalyse    string      `json:"squad_data_analyse"`
	} `json:"daily"`
	Weekly struct {
		ID                  string      `json:"id"`
		UserID              string      `json:"user_id"`
		EndTime             int         `json:"end_time"`
		Fix                 int         `json:"fix"`
		WinRobot            int         `json:"win_robot"`
		WinPerson           int         `json:"win_person"`
		WinBoss             int         `json:"win_boss"`
		WinArmorrobot       int         `json:"win_armorrobot"`
		WinArmorperson      int         `json:"win_armorperson"`
		Operation           interface{} `json:"operation"` // str or int
		SWin                int         `json:"s_win"`
		Eat                 interface{} `json:"eat"` // str or int
		DevelopGun          string      `json:"develop_gun"`
		Upgrade             int         `json:"upgrade"`
		CoinMission         int         `json:"coin_mission"`
		DevelopEquip        int         `json:"develop_equip"`
		SpecialDevelopGun   int         `json:"special_develop_gun"`
		AdjustEquip         int         `json:"adjust_equip"`
		EatEquip            int         `json:"eat_equip"`
		SpecialDevelopEquip int         `json:"special_develop_equip"`
		AdjustFairy         int         `json:"adjust_fairy"`
		EatFairy            int         `json:"eat_fairy"`
		SquadDataAnalyse    int         `json:"squad_data_analyse"`
		SquadEatChip        int         `json:"squad_eat_chip"`
	} `json:"weekly"`
	Career struct {
		DevelopGun          int    `json:"develop_gun"`
		AutoMission1        string `json:"auto_mission_1"`
		Gun5IntoTeam        string `json:"gun5_into_team"`
		CombineGun          string `json:"combine_gun"`
		GashaCount          string `json:"gasha_count"`
		OpenGift            string `json:"open_gift"`
		DormChange          string `json:"dorm_change"`
		FriendVisit         string `json:"friend_visit"`
		EatGun              string `json:"eat_gun"`
		SunFriendTeamInto   string `json:"sun_friend_team_into"`
		NightFriendTeamInto string `json:"night_friend_team_into"`
	} `json:"career"`
	StaticCareerQuest []struct {
		ID           string `json:"id"`
		UnlockLv     string `json:"unlock_lv"`
		UnlockIds    string `json:"unlock_ids"`
		UnlockLabel  string `json:"unlock_label"`
		Type         string `json:"type"`
		Count        string `json:"count"`
		PrizeID      string `json:"prize_id"`
		Title        string `json:"title"`
		Content      string `json:"content"`
		UnlockCourse string `json:"unlock_course"`
	} `json:"static_career_quest"`
}
