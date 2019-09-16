package defs

import "hoxy/proxy/defs/types"

// SEquip/developCollectList
type SEquipDevelopCollectList struct {
	LogList []struct {
		ID           string      `json:"id"`
		LogID        string      `json:"log_id"`
		UserID       string      `json:"user_id"`
		DevUID       interface{} `json:"dev_uid"`
		DevUname     types.Str   `json:"dev_uname"`
		DevLv        string      `json:"dev_lv"`
		EquipID      string      `json:"equip_id"`
		Mp           string      `json:"mp"`
		Ammo         string      `json:"ammo"`
		Mre          string      `json:"mre"`
		Part         string      `json:"part"`
		InputLevel   string      `json:"input_level"`
		Core         string      `json:"core"`
		ItemNum      string      `json:"item_num"`
		DevTime      string      `json:"dev_time"`
		CollectTime  string      `json:"collect_time"`
		FairyID      string      `json:"fairy_id"`
		PassiveSkill string      `json:"passive_skill"`
		QualityLv    string      `json:"quality_lv"`
		DevType      string      `json:"dev_type"`
	} `json:"log_list"`
}
