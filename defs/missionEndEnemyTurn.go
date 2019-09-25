package defs

// SMission/endEnemyTurn
type SMissionEndEnemyTurn struct {
	ChangeBelong           map[int]int   `json:"change_belong"`
	BuildingChangeBelong   []interface{} `json:"building_change_belong"`
	BuildingDefenderChange []interface{} `json:"building_defender_change"`
	MissionLoseResult      []interface{} `json:"mission_lose_result"`
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
}
