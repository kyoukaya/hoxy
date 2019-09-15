package defs

// Gun/finishAllDevelop
type CGunFinishAllDevelop struct {
	IsCostItem3 int `json:"is_cost_item3"`
}

type SGunFinishAllDevelop struct {
	GunWithUserAddList []struct {
		BuildSlot     int    `json:"build_slot"`
		GunWithUserID string `json:"gun_with_user_id"`
		GunID         string `json:"gun_id"`
	} `json:"gun_with_user_add_list"`
	CostItem3Num int `json:"cost_item3_num"`
}
