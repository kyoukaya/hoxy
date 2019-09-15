package defs

type CGunFinishDevelop struct {
	BuildSlot int `json:"build_slot"`
}

type SGunFinishDevelop struct {
	GunWithUserAdd struct {
		GunWithUserID string `json:"gun_with_user_id"`
		GunID         string `json:"gun_id"`
	} `json:"gun_with_user_add"`
}
