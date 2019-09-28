package defs

// CDorm/giftToGun
type CDormGiftToGun struct {
	ItemID        int `json:"item_id"`
	GunWithUserID int `json:"gun_with_user_id"`
	Num           int `json:"num"`
	New           int `json:"new"`
}

type SDormGiftToGun struct {
	AllFavorupGun        []interface{} `json:"all_favorup_gun"`
	LastFavorRecoverTime int           `json:"last_favor_recover_time"`
	SkinAddFavor         int           `json:"skin_add_favor"`
	GunWithUserID        int           `json:"gun_with_user_id"`
	GunExp               int           `json:"gun_exp"`
}
