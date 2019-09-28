package defs

// CMission/fairySkillPerform
type CMissionFairySkillPerform struct {
	FairyTeamID int   `json:"fairy_team_id"`
	FairySpot   int   `json:"fairy_spot"`
	SpotID      []int `json:"spot_id"`
}

type SMissionFairySkillPerform struct {
	FairySkillPerform []struct {
		FairyTeamID          int `json:"fairy_team_id"`
		BuildingSpotID       int `json:"building_spot_id"`
		SquadInstanceID      int `json:"squad_instance_id"`
		NextSkillCdTurn      int `json:"next_skill_cd_turn"`
		MissionSkillConfigID int `json:"mission_skill_config_id"`
		PerformSpotID        int `json:"perform_spot_id"`
	} `json:"fairy_skill_perform"`
	FairySkillOnSpot []interface{} `json:"fairy_skill_on_spot"`
	FairySkillOnTeam map[int]map[int]struct {
		TeamID               int    `json:"team_id"`
		BuffID               string `json:"buff_id"`
		FairyTeamID          int    `json:"fairy_team_id"`
		BuildingSpotID       int    `json:"building_spot_id"`
		SquadInstanceID      int    `json:"squad_instance_id"`
		StartTurn            string `json:"start_turn"`
		BattleCount          int    `json:"battle_count"`
		MissionSkillConfigID string `json:"mission_skill_config_id"`
	} `json:"fairy_skill_on_team"`
	FairySkillOnEnemy []interface{} `json:"fairy_skill_on_enemy"`
	FairySkillOnSquad []interface{} `json:"fairy_skill_on_squad"`
	FairySkillReturn  struct {
		FairyEffect map[int]struct {
			OriginKey   int    `json:"origin_key"`
			OriginSpot  int    `json:"origin_spot"`
			OriginType  int    `json:"origin_type"`
			OriginStcID string `json:"origin_stc_id"`
			AimSpot     int    `json:"aim_spot"`
			AimType     int    `json:"aim_type"`
			AimStcID    string `json:"aim_stc_id"`
		} `json:"fairy_effect"`
	} `json:"fairy_skill_return"`
}
