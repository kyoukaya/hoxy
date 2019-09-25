package defs

// CFairy/eatFairy
type CFairyEatFairy struct {
	FairyWithUserID int   `json:"fairy_with_user_id"`
	Food            []int `json:"food"`
}

type SFairyEatFairy struct {
	AddFairyExp int `json:"add_fairy_exp"`
}
