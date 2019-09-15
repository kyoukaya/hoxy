package defs

// CEquip/developMulti
type CEquipDevelopMulti struct {
	Mp         int `json:"mp"`
	Ammo       int `json:"ammo"`
	Mre        int `json:"mre"`
	Part       int `json:"part"`
	InputLevel int `json:"input_level"`
	BuildQuick int `json:"build_quick"`
	BuildMulti int `json:"build_multi"`
	BuildHeavy int `json:"build_heavy"`
}

// SEquip/developMulti
type SEquipDevelopMulti struct {
	EquipIds []struct {
		Info struct {
			Type    int `json:"type"`
			EquipID int `json:"equip_id"`
		} `json:"info"`
		Slot int `json:"slot"`
	} `json:"equip_ids"`
}
