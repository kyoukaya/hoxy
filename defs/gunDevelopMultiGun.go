package defs

// Gun/developMultiGun
type CGunDevelopMultiGun struct {
	Mp         int `json:"mp"`
	Ammo       int `json:"ammo"`
	Mre        int `json:"mre"`
	Part       int `json:"part"`
	InputLevel int `json:"input_level"`
	BuildQuick int `json:"build_quick"`
	BuildMulti int `json:"build_multi"`
	BuildHeavy int `json:"build_heavy"`
}

type SGunDevelopMultiGun struct {
	GunIds []struct {
		ID   int `json:"id"`
		Slot int `json:"slot"`
	} `json:"gun_ids"`
}
