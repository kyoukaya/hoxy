package defs

// CEquip/gunEquip
type CEquipGunEquip struct {
	IfOutfit        int `json:"if_outfit"`
	GunWithUserID   int `json:"gun_with_user_id"`
	EquipWithUserID int `json:"equip_with_user_id"`
	EquipSlot       int `json:"equip_slot"`
}

type SEquipGunEquip int
