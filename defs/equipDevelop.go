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
	Type int `json:"type"`
	// Fairy info
	FairyID      int `json:"fairy_id"`
	PassiveSkill int `json:"passive_skill"`
	QualityLv    int `json:"quality_lv"`
	// Equip info
	EquipID int `json:"equip_id"`
}
