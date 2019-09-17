package defs

import "github.com/iancoleman/orderedmap"

// TODO: document what all these cryptic fields mean.

// Mission/battleFinish
type CMissionBattleFinish struct {
	SpotID         int           `json:"spot_id"`
	IfEnemyDie     bool          `json:"if_enemy_die"`
	CurrentTime    int           `json:"current_time"`
	BossHp         int           `json:"boss_hp"`
	Mvp            int           `json:"mvp"`
	LastBattleInfo string        `json:"last_battle_info"`
	UseSkillSquads []interface{} `json:"use_skill_squads"`
	Guns           []struct {
		ID   int `json:"id"`
		Life int `json:"life"`
	} `json:"guns"`
	UserRec string `json:"user_rec"` // contains seed, and a record of movements
	Num1000 struct {
		Num10 int `json:"10"`
		Num11 int `json:"11"`
		Num12 int `json:"12"`
		Num13 int `json:"13"`
		Num15 int `json:"15"`
		Num16 int `json:"16"`
		Num17 int `json:"17"`
		Num33 int `json:"33"`
		Num40 int `json:"40"`
		Num18 int `json:"18"`
		Num19 int `json:"19"`
		Num20 int `json:"20"`
		Num21 int `json:"21"`
		Num22 int `json:"22"`
		Num23 int `json:"23"`
		Num24 int `json:"24"`
		Num25 int `json:"25"`
		Num26 int `json:"26"`
		Num27 int `json:"27"`
		Num34 int `json:"34"`
		Num35 int `json:"35"`
		Num41 int `json:"41"`
		Num42 int `json:"42"`
		Num43 int `json:"43"`
		Num44 int `json:"44"`
		Num92 int `json:"92"`
	} `json:"1000"`
	Num1001 struct {
	} `json:"1001"`
	// Value is an ordered map like so, [1,2,3,4,5] are gun_with_user_id
	// {"gun_with_user_id":{"47":01},"gun_with_user_id":{"47":10}
	// "gun_with_user_id":{"47":0},"gun_with_user_id":{"47":0},"gun_with_user_id":{"47":0}}
	Num1002 orderedmap.OrderedMap `json:"1002"`
	// Seems to be a map of a fairy id to some unknown object
	// {"fairy_id":{"9":0,"68":0}}
	Num1003      orderedmap.OrderedMap `json:"1003"`
	Num1005      interface{}           `json:"1005"`
	BattleDamage struct {
	} `json:"battle_damage"`
}

type SMissionBattleFinish struct {
	BattleGetGun struct {
		GunWithUserID string `json:"gun_with_user_id"`
		GunID         string `json:"gun_id"`
	} `json:"battle_get_gun"`
	UserExp string `json:"user_exp"`
	GunExp  []struct {
		GunWithUserID string `json:"gun_with_user_id"`
		Exp           string `json:"exp"`
	} `json:"gun_exp"`
	FairyExp               *int                  `json:"fairy_exp,omitempty"`
	GunLife                []interface{}         `json:"gun_life"`
	SquadExp               []interface{}         `json:"squad_exp"`
	BattleRank             string                `json:"battle_rank"`
	FreeExp                int                   `json:"free_exp"`
	ChangeBelong           []interface{}         `json:"change_belong"`
	BuildingDefenderChange []interface{}         `json:"building_defender_change"`
	MissionWinResult       []interface{}         `json:"mission_win_result"`
	Seed                   int                   `json:"seed"`
	FavorChange            orderedmap.OrderedMap `json:"favor_change"` // map[string]Int
	Type5Score             string                `json:"type5_score"`
	AllyInstanceTransform  []interface{}         `json:"ally_instance_transform"`
	AllyInstanceBetray     []interface{}         `json:"ally_instance_betray"`
	MissionControl         struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
	} `json:"mission_control"`
}
