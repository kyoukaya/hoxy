package defs

// Gun/eatGun
type CGunEatGun struct {
	GunWithUserID int   `json:"gun_with_user_id"`
	Item9Num      int   `json:"item9_num"`
	Food          []int `json:"food"`
}

type SGunEatGun struct {
	Pow   int `json:"pow"`
	Dodge int `json:"dodge"`
	Hit   int `json:"hit"`
	Rate  int `json:"rate"`
}
