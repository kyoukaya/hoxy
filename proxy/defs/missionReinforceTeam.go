package defs

// CMission/reinforceTeam
type CMissionReinforceTeam struct {
	SpotID int `json:"spot_id"`
	TeamID int `json:"team_id"`
}

type SMissionReinforceTeam struct {
	Ap                int           `json:"ap"`
	FairySkillReturn  []interface{} `json:"fairy_skill_return"`
	FairySkillPerform []interface{} `json:"fairy_skill_perform"`
	FairySkillOnSpot  []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam  []interface{} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad []interface{} `json:"fairy_skill_on_squad"`
}
