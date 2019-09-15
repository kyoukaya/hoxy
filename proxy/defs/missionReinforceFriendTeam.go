package defs

// CMission/reinforceFriendTeam
type CMissionReinforceFriendTeam struct {
	SpotID       int   `json:"spot_id"`
	FriendTeamID int   `json:"friend_team_id"`
	GroupID      int   `json:"group_id"`
	FriendGunids []int `json:"friend_gunids"`
}

// SMission/reinforceFriendTeam
type SMissionReinforceFriendTeam struct {
	Guns            interface{} `json:"guns"`
	Equips          interface{} `json:"equips"`
	FriendTeamFairy struct {
		ID              string `json:"id"`
		UserID          string `json:"user_id"`
		GroupID         string `json:"group_id"`
		FairyID         string `json:"fairy_id"`
		FairyLv         string `json:"fairy_lv"`
		FairyExp        string `json:"fairy_exp"`
		QualityLv       string `json:"quality_lv"`
		QualityExp      string `json:"quality_exp"`
		SkillLv         string `json:"skill_lv"`
		PassiveSkill    string `json:"passive_skill"`
		EquipID         string `json:"equip_id"`
		Skin            string `json:"skin"`
		FairyWithUserID string `json:"fairy_with_user_id"`
	} `json:"friend_team_fairy"`
	FriendTeamEndTime int           `json:"friend_team_end_time"`
	Ap                string        `json:"ap"`
	FairySkillReturn  []interface{} `json:"fairy_skill_return"`
	FairySkillPerform []interface{} `json:"fairy_skill_perform"`
	FairySkillOnSpot  []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam  []interface{} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad []interface{} `json:"fairy_skill_on_squad"`
}
