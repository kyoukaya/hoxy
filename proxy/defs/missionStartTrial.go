package defs

// CMission/startTrial
type CMissionStartTrial struct {
	TeamIds    string `json:"team_ids"`
	BattleTeam int    `json:"battle_team"`
}

type SMissionStartTrial struct {
	TrialID int `json:"trial_id"`
}
