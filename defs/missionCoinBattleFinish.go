package defs

// CMission/coinBattleFinish
type CMissionCoinBattleFinish struct {
	Num1000 struct {
		Num10 int `json:"10"`
		Num11 int `json:"11"`
		Num12 int `json:"12"`
		Num13 int `json:"13"`
		Num15 int `json:"15"`
		Num16 int `json:"16"`
		Num17 int `json:"17"`
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
		Num33 int `json:"33"`
		Num34 int `json:"34"`
		Num35 int `json:"35"`
		Num40 int `json:"40"`
		Num41 int `json:"41"`
		Num42 int `json:"42"`
		Num43 int `json:"43"`
		Num44 int `json:"44"`
	} `json:"1000"`
	Num1001 struct {
	} `json:"1001"`
	Num1002 map[int]struct {
		Num47 int `json:"47"`
	} `json:"1002"`
	Num1003 struct {
	} `json:"1003"`
	Num1005 struct {
	} `json:"1005"`
	MissionID  int     `json:"mission_id"`
	BossHp     int     `json:"boss_hp"`
	Duration   float64 `json:"duration"`
	BattleTime struct {
	} `json:"battle_time"`
}

type SMissionCoinBattleFinish struct {
	CoinNum  string `json:"coin_num"`
	CoinType string `json:"coin_type"`
}
