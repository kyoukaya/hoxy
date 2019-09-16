package defs

// SFriend/teamGuns
type SFriendTeamGuns struct {
	BorrowTeamToday int `json:"borrow_team_today"`
	GunWithFriend   struct {
		GunsWithFriendAvailable []struct {
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
		} `json:"guns_with_friend_available"`
		FriendTeamEffectArr interface{} `json:"friend_team_effect_arr"`
	} `json:"gun_with_friend"`
	FairyWithFriendInfo interface{} `json:"fairy_with_friend_info"`
}