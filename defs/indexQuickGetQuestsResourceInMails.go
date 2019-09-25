package defs

// Index/QuickGetQuestsResourceInMails
type CIndexQuickGetQuestsResourceInMails struct {
	Type int `json:"type"`
}

type SIndexQuickGetQuestsResourceInMails struct {
	UserExp   int           `json:"user_exp"`
	Mp        int           `json:"mp"`
	Ammo      int           `json:"ammo"`
	Mre       int           `json:"mre"`
	Part      int           `json:"part"`
	Core      int           `json:"core"`
	Gem       int           `json:"gem"`
	GunID     []interface{} `json:"gun_id"`
	ItemIds   string        `json:"item_ids"`
	EquipIds  []interface{} `json:"equip_ids"`
	Furniture []interface{} `json:"furniture"`
	Gift      string        `json:"gift"`
	Skin      []interface{} `json:"skin"`
	BpPay     int           `json:"bp_pay"`
	Coins     string        `json:"coins"`
	FairyIds  string        `json:"fairy_ids"`
}
