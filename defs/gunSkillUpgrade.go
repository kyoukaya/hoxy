package defs

// CGun/skillUpgrade
type CGunSkillUpgrade struct {
	Skill         int `json:"skill"`
	IfQuick       int `json:"if_quick"`
	GunWithUserID int `json:"gun_with_user_id"`
	UpgradeSlot   int `json:"upgrade_slot"`
}

type SGunSkillUpgrade int
