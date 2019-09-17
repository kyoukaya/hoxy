package defs

// CFairy/skillUpgrade
type CFairySkillUpgrade struct {
	FairyWithUserID int  `json:"fairy_with_user_id"`
	UpgradeSlot     int  `json:"upgrade_slot"`
	IfQuick         bool `json:"if_quick"`
}

type SFairySkillUpgrade int
