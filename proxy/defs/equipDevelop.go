package defs

// CEquip/develop
type CEquipDevelop struct {
	Mp         int `json:"mp"`
	Ammo       int `json:"ammo"`
	Mre        int `json:"mre"`
	Part       int `json:"part"`
	BuildSlot  int `json:"build_slot"`
	InputLevel int `json:"input_level"`
}

type SEquipDevelop struct {
	Type    int `json:"type"`
	EquipID int `json:"equip_id"`
}
