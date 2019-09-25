package defs

// Mission/teamMove
type CMissionTeamMove struct {
	TeamID     int `json:"team_id"`
	FromSpotID int `json:"from_spot_id"`
	ToSpotID   int `json:"to_spot_id"`
	MoveType   int `json:"move_type"`
}

type SMissionTeamMove struct {
	// Fields that only appear when resources are gained when a random node is triggered.
	Mp   int `json:"mp"`
	Ammo int `json:"ammo"`
	Mre  int `json:"mre"`
	Part int `json:"part"`

	BuildingDefenderChange []interface{} `json:"building_defender_change"`
	Ap                     int           `json:"ap"`
	Type5Score             string        `json:"type5_score"`
	FairySkillReturn       []interface{} `json:"fairy_skill_return"`
	FairySkillPerform      []interface{} `json:"fairy_skill_perform"`
	FairySkillOnSpot       []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam       []interface{} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy      []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad      []interface{} `json:"fairy_skill_on_squad"`
	AllyInstanceTransform  []interface{} `json:"ally_instance_transform"`
	AllyInstanceBetray     []interface{} `json:"ally_instance_betray"`
	MissionControl         struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
	} `json:"mission_control"`
	ChangeBelong []interface{} `json:"change_belong"`
}
