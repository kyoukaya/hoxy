package defs

// CEquip/finishDevelop
type CEquipFinishDevelop struct {
	BuildSlot int `json:"build_slot"`
}

// SEquip/finishDevelop
type SEquipFinishDevelop struct {
	EquipWithUser struct {
		ID               string `json:"id"`
		UserID           string `json:"user_id"`
		GunWithUserID    string `json:"gun_with_user_id"`
		EquipID          string `json:"equip_id"`
		EquipExp         string `json:"equip_exp"`
		EquipLevel       string `json:"equip_level"`
		Pow              string `json:"pow"`
		Hit              string `json:"hit"`
		Dodge            string `json:"dodge"`
		Speed            string `json:"speed"`
		Rate             string `json:"rate"`
		CriticalHarmRate string `json:"critical_harm_rate"`
		CriticalPercent  string `json:"critical_percent"`
		ArmorPiercing    string `json:"armor_piercing"`
		Armor            string `json:"armor"`
		Shield           string `json:"shield"`
		DamageAmplify    string `json:"damage_amplify"`
		DamageReduction  string `json:"damage_reduction"`
		NightViewPercent string `json:"night_view_percent"`
		BulletNumberUp   string `json:"bullet_number_up"`
		AdjustCount      string `json:"adjust_count"`
		IsLocked         string `json:"is_locked"`
		LastAdjust       string `json:"last_adjust"`
	} `json:"equip_with_user"`
	FairyWithUser int `json:"fairy_with_user"`
}
