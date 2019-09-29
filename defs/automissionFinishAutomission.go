package defs

// CAutomission/finishAutomission
type CAutomissionFinishAutomission struct {
	AutoMissionID int `json:"auto_mission_id"`
}

type SAutomissionFinishAutomission struct {
	AddGunExp []struct {
		GunWithUserID int `json:"gun_with_user_id"`
		GunExp        int `json:"gun_exp"`
		GunLife       int `json:"gun_life"`
	} `json:"add_gun_exp"`
	AddFairyExp []struct {
		FairyWithUserID int `json:"fairy_with_user_id"`
		FairyExp        int `json:"fairy_exp"`
	} `json:"add_fairy_exp"`
	FreeExp       int           `json:"free_exp"`
	AddUserExp    int           `json:"add_user_exp"`
	AddPrize      []interface{} `json:"add_prize"`
	FavorChange   map[int]int   `json:"favor_change"`
	SuccessNumber int           `json:"success_number"`
}
