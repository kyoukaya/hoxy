package defs

// CFriend/praise
type CFriendPraise struct {
	FUserid int `json:"f_userid"`
}

type SFriendPraise struct {
	Package struct {
		ID            string        `json:"id"`
		Name          string        `json:"name"`
		UserExp       string        `json:"user_exp"`
		Mp            string        `json:"mp"`
		Ammo          string        `json:"ammo"`
		Mre           string        `json:"mre"`
		Part          string        `json:"part"`
		Core          string        `json:"core"`
		Gem           string        `json:"gem"`
		GunID         string        `json:"gun_id"`
		ItemIds       string        `json:"item_ids"`
		Furniture     []interface{} `json:"furniture"`
		Gift          string        `json:"gift"`
		EquipIds      string        `json:"equip_ids"`
		Coins         string        `json:"coins"`
		Skin          string        `json:"skin"`
		Content       string        `json:"content"`
		SendLimit     string        `json:"send_limit"`
		Icon          string        `json:"icon"`
		BpPay         string        `json:"bp_pay"`
		FairyIds      string        `json:"fairy_ids"`
		Chip          []interface{} `json:"chip"`
		GunWithUserID int           `json:"gun_with_user_id"`
		Fairys        []interface{} `json:"fairys"`
		Equips        []interface{} `json:"equips"`
	} `json:"package"`
}
