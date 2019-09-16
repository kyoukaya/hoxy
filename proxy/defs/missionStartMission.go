package defs

import "hoxy/proxy/defs/types"

// Mission/startMission
type CMissionStartMission struct {
	MissionID int `json:"mission_id"`
	Spots     []struct {
		SpotID int `json:"spot_id"`
		TeamID int `json:"team_id"`
	} `json:"spots"`
	SquadSpots []interface{} `json:"squad_spots"`
	AllyID     int           `json:"ally_id"`
}

type SMissionStartMission struct {
	BuildingInfo []interface{} `json:"building_info"`
	SquadInfo    []interface{} `json:"squad_info"`
	Ap           int           `json:"ap"`
	// Can throw marshal mismatch errors because it randomly changes ordering sometimes.
	SpotActInfo []struct {
		SpotID           string        `json:"spot_id"`
		EnemyTeamID      string        `json:"enemy_team_id"`
		BossHp           string        `json:"boss_hp"`
		EnemyHpPercent   string        `json:"enemy_hp_percent"`
		EnemyInstanceID  string        `json:"enemy_instance_id"`
		EnemyBirthTurn   string        `json:"enemy_birth_turn"`
		EnemyAi          types.Int     `json:"enemy_ai"`
		EnemyAiPara      types.Int     `json:"enemy_ai_para"`
		Belong           string        `json:"belong"`
		IfRandom         string        `json:"if_random"`
		Seed             int           `json:"seed"`
		TeamID           string        `json:"team_id"`
		AllyInstanceIds  []interface{} `json:"ally_instance_ids"`
		SquadInstanceIds []interface{} `json:"squad_instance_ids"`
		HostageID        string        `json:"hostage_id"`
		HostageHp        string        `json:"hostage_hp"`
		HostageMaxHp     string        `json:"hostage_max_hp"`
		ReinforceCount   string        `json:"reinforce_count"`
		SupplyCount      string        `json:"supply_count"`
	} `json:"spot_act_info"`
	AllyInstanceInfo  []interface{} `json:"ally_instance_info"`
	FairySkillReturn  []interface{} `json:"fairy_skill_return"`
	FairySkillPerform []interface{} `json:"fairy_skill_perform"`
	FairySkillOnSpot  []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam  []interface{} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad []interface{} `json:"fairy_skill_on_squad"`
}
