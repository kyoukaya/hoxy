package defs

// Gun/developGun
type CGunDevelopGun struct {
	Mp         int `json:"mp"`
	Ammo       int `json:"ammo"`
	Mre        int `json:"mre"`
	Part       int `json:"part"`
	BuildSlot  int `json:"build_slot"`
	InputLevel int `json:"input_level"`
}

type SGunDevelopGun struct {
	GunID int `json:"gun_id"`
}
