package defs

// CIndexGetMailList Client's request for Index/getMailList
type CIndexGetMailList struct {
	StartID    int `json:"start_id"`
	IgnoreTime int `json:"ignore_time"`
}

// SIndexGetMailList Server's response for Index/getMailList
type SIndexGetMailList []struct {
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
}
