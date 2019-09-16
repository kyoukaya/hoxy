package defs

import "hoxy/proxy/defs/types"

// CIndexQuest has no request body

// SIndexQuest Server's Index/Quest response
type SIndexQuest struct {
	Daily struct {
		Operation           types.Int `json:"operation"`
		Fix                 types.Int `json:"fix"`
		Mission             types.Int `json:"mission,string"`
		DevelopGun          types.Int `json:"develop_gun"`
		ID                  types.Int `json:"id"`
		UserID              types.Int `json:"user_id"`
		EndTime             types.Int `json:"end_time"`
		Eat                 types.Int `json:"eat"`
		Upgrade             types.Int `json:"upgrade"`
		CoinMission         types.Int `json:"coin_mission"`
		DevelopEquip        types.Int `json:"develop_equip"`
		WinRobot            types.Int `json:"win_robot"`
		WinPerson           types.Int `json:"win_person"`
		WinBoss             types.Int `json:"win_boss"`
		WinArmorrobot       types.Int `json:"win_armorrobot"`
		WinArmorperson      types.Int `json:"win_armorperson"`
		EatEquip            types.Int `json:"eat_equip"`
		FromFriendBuildCoin types.Int `json:"from_friend_build_coin"`
		BorrowFriendTeam    types.Int `json:"borrow_friend_team"`
		SquadDataAnalyse    types.Int `json:"squad_data_analyse"`
	} `json:"daily"`
	Weekly struct {
		ID                  types.Int `json:"id"`
		UserID              types.Int `json:"user_id"`
		EndTime             types.Int `json:"end_time"`
		Fix                 types.Int `json:"fix"`
		WinRobot            types.Int `json:"win_robot"`
		WinPerson           types.Int `json:"win_person"`
		WinBoss             types.Int `json:"win_boss"`
		WinArmorrobot       types.Int `json:"win_armorrobot"`
		WinArmorperson      types.Int `json:"win_armorperson"`
		Operation           types.Int `json:"operation"`
		SWin                types.Int `json:"s_win"`
		Eat                 types.Int `json:"eat"`
		DevelopGun          types.Int `json:"develop_gun"`
		Upgrade             types.Int `json:"upgrade"`
		CoinMission         types.Int `json:"coin_mission"`
		DevelopEquip        types.Int `json:"develop_equip"`
		SpecialDevelopGun   types.Int `json:"special_develop_gun"`
		AdjustEquip         types.Int `json:"adjust_equip"`
		EatEquip            types.Int `json:"eat_equip"`
		SpecialDevelopEquip types.Int `json:"special_develop_equip"`
		AdjustFairy         types.Int `json:"adjust_fairy"`
		EatFairy            types.Int `json:"eat_fairy"`
		SquadDataAnalyse    types.Int `json:"squad_data_analyse"`
		SquadEatChip        types.Int `json:"squad_eat_chip"`
	} `json:"weekly"`
	Career struct {
		DevelopGun          types.Int `json:"develop_gun"`
		AutoMission1        types.Int `json:"auto_mission_1"`
		Gun5IntoTeam        types.Int `json:"gun5_into_team"`
		CombineGun          types.Int `json:"combine_gun"`
		GashaCount          types.Int `json:"gasha_count"`
		OpenGift            types.Int `json:"open_gift"`
		DormChange          types.Int `json:"dorm_change"`
		FriendVisit         types.Int `json:"friend_visit"`
		EatGun              types.Int `json:"eat_gun"`
		SunFriendTeamInto   types.Int `json:"sun_friend_team_into"`
		NightFriendTeamInto types.Int `json:"night_friend_team_into"`
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
