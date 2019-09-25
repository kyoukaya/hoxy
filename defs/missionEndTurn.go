package defs

// Hoxy 22:46:24 INFO >>## [Mission/endTurn]#156794553100029

// Mission/endTurn
type SMissionEndTurn struct {
	ChangeBelong           map[int]int `json:"change_belong"`
	ChangeBelong1          map[int]int `json:"change_belong1"`
	ChangeBelong2          map[int]int `json:"change_belong2"`
	BuildingChangeBelong   map[int]int `json:"building_change_belong"`
	BuildingChangeBelong1  map[int]int `json:"building_change_belong1"`
	BuildingChangeBelong2  map[int]int `json:"building_change_belong2"`
	BuildingDefenderChange map[int]int `json:"building_defender_change"`
	// TODO: define lose event
	MissionLoseResult []interface{} `json:"mission_lose_result"`
	MissionWinResult  *struct {
		Rank      int         `json:"rank"`
		Medal4    int         `json:"medal4"`
		Open      map[int]int `json:"open"`
		UserExp   int         `json:"user_exp"`
		RewardGun struct {
			GunWithUserID int `json:"gun_with_user_id"`
			GunID         int `json:"gun_id"`
		} `json:"reward_gun"`
		MissionInfo struct {
			Turn                    int `json:"turn"`
			EnemydieNum             int `json:"enemydie_num"`
			EnemydieNumKillbyfriend int `json:"enemydie_num_killbyfriend"`
			GundieNum               int `json:"gundie_num"`
		} `json:"mission_info"`
	} `json:"mission_win_result"`
	EnemyMove []struct {
		FromSpotID      int `json:"from_spot_id"`
		ToSpotID        int `json:"to_spot_id"`
		EnemyAi         int `json:"enemy_ai"`
		EnemyAiPara     int `json:"enemy_ai_para"`
		HostageID       int `json:"hostage_id"`
		HostageHp       int `json:"hostage_hp"`
		SquadInstanceID int `json:"squad_instance_id"`
	} `json:"enemy_move"`
	AllyMove              []interface{} `json:"ally_move"`
	AllyBattle            []interface{} `json:"ally_battle"`
	Type5Score            string        `json:"type5_score"`
	FairySkillReturn      []interface{} `json:"fairy_skill_return"`
	FairySkillPerform     []interface{} `json:"fairy_skill_perform"`
	FairySkillOnSpot      []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam      []interface{} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy     []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad     []interface{} `json:"fairy_skill_on_squad"`
	AllyInstanceTransform []interface{} `json:"ally_instance_transform"`
	AllyInstanceBetray    []interface{} `json:"ally_instance_betray"`
	MissionControl        struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
	} `json:"mission_control"`
}
