package defs

// SEquip/developLog
type SEquipDevelopLog struct {
	LogList []struct {
		ID           string `json:"id"`
		UserID       string `json:"user_id"`
		DevType      string `json:"dev_type"`
		BuildSlot    string `json:"build_slot"`
		DevUname     string `json:"dev_uname"`
		DevLv        string `json:"dev_lv"`
		EquipID      string `json:"equip_id"`
		Mp           string `json:"mp"`
		Ammo         string `json:"ammo"`
		Mre          string `json:"mre"`
		Part         string `json:"part"`
		InputLevel   string `json:"input_level"`
		Core         string `json:"core"`
		ItemNum      string `json:"item_num"`
		DevTime      string `json:"dev_time"`
		FairyID      string `json:"fairy_id"`
		PassiveSkill string `json:"passive_skill"`
		QualityLv    string `json:"quality_lv"`
	} `json:"log_list"`
}
