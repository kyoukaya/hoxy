package defs

// CFriend/usercard
type CFriendUsercard struct {
	FUserid int `json:"f_userid"`
}

type SFriendUsercard struct {
	FUserid int    `json:"f_userid"`
	Name    string `json:"name"`
	Lv      string `json:"lv"`
	Medal   []struct {
		ID    string `json:"id"`
		Level string `json:"level"`
		Num   string `json:"num"`
	} `json:"medal"`
	CardID    string `json:"card_id"`
	HeadpicID string `json:"headpic_id"`
	HomebgID  string `json:"homebg_id"`
	Intro     string `json:"intro"`
	TeamInfo  []struct {
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
	} `json:"team_info"`
	BorrowTeamToday int `json:"borrow_team_today"`
	FairyInfo       []struct {
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
	} `json:"fairy_info"`
	IsReturnUser int `json:"is_return_user"`
}
