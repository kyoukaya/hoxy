package defs

// CMission/reinforceFriendTeam
type CMissionReinforceFriendTeam struct {
	SpotID       int   `json:"spot_id"`
	FriendTeamID int   `json:"friend_team_id"`
	GroupID      int   `json:"group_id"`
	FriendGunids []int `json:"friend_gunids"`
}

// TODO: Fix SMission/reinforceFriendTeam
// SMission/reinforceFriendTeam
type SMissionReinforceFriendTeam struct {
	Guns map[int]struct {
		ID             string `json:"id"`
		UserID         string `json:"user_id"`
		GroupID        string `json:"group_id"`
		GunID          string `json:"gun_id"`
		GunExp         string `json:"gun_exp"`
		GunLevel       string `json:"gun_level"`
		Location       string `json:"location"`
		Position       string `json:"position"`
		Pow            string `json:"pow"`
		Hit            string `json:"hit"`
		Dodge          string `json:"dodge"`
		Rate           string `json:"rate"`
		Skill1         string `json:"skill1"`
		Skill2         string `json:"skill2"`
		Number         string `json:"number"`
		Equip1         string `json:"equip1"`
		Equip2         string `json:"equip2"`
		Equip3         string `json:"equip3"`
		Favor          string `json:"favor"`
		Skin           string `json:"skin"`
		SoulBond       string `json:"soul_bond"`
		IfModification string `json:"if_modification"`
		GunWithUserID  string `json:"gun_with_user_id"`
	} `json:"guns"`
	Equips map[int]struct {
		ID               string `json:"id"`
		UserID           string `json:"user_id"`
		GroupID          string `json:"group_id"`
		GunWithFriendID  string `json:"gun_with_friend_id"`
		EquipID          string `json:"equip_id"`
		EquipExp         string `json:"equip_exp"`
		EquipLevel       string `json:"equip_level"`
		Pow              string `json:"pow"`
		Hit              string `json:"hit"`
		Dodge            string `json:"dodge"`
		Speed            string `json:"speed"`
		Rate             string `json:"rate"`
		CriticalHarmRate string `json:"critical_harm_rate"`
		CriticalPercent  string `json:"critical_percent"`
		ArmorPiercing    string `json:"armor_piercing"`
		Armor            string `json:"armor"`
		Shield           string `json:"shield"`
		DamageAmplify    string `json:"damage_amplify"`
		DamageReduction  string `json:"damage_reduction"`
		NightViewPercent string `json:"night_view_percent"`
		BulletNumberUp   string `json:"bullet_number_up"`
		EquipWithUserID  string `json:"equip_with_user_id"`
	} `json:"equips"`
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
