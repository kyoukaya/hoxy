package defs

// CMission / coinTeam
type CMissionCoinTeam struct {
	TeamChange map[int]struct {
		GunWithUserID int `json:"gun_with_user_id"`
		Position      int `json:"position"`
	} `json:"team_change"`
}

type SMissionCoinTeam int
