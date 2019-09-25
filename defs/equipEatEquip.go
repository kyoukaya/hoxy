package defs

// CEquip/eatEquip
type CEquipEatEquip struct {
	EquipWithUserID int   `json:"equip_with_user_id"`
	Food            []int `json:"food"`
}

type SEquipEatEquip struct {
	EquipAddExp int `json:"equip_add_exp"`
}
