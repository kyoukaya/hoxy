package defs

import "github.com/iancoleman/orderedmap"

// CEquip/adjust
type CEquipAdjust struct {
	EquipWithUserID int `json:"equip_with_user_id"`
}

// [] if nil
type SEquipAdjust orderedmap.OrderedMap
