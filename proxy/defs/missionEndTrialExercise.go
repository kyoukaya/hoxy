package defs

import (
	"hoxy/proxy/defs/types"

	"github.com/iancoleman/orderedmap"
)

// CMission/endTrialExercise
type CMissionEndTrialExercise struct {
	IfWin int `json:"if_win"`
	// Ordered map of gun_with_user_id -> struct {
	// 		Life int `json:"life"`
	// 		Dps  int `json:"dps"`
	// 	}
	BattleGuns orderedmap.OrderedMap `json:"battle_guns"`
	UserRec    string                `json:"user_rec"`
	Num1000    struct {
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
	} `json:"1000"`
	Num1001 struct {
	} `json:"1001"`
	// gun_with_user_id -> struct {
	// 	Num47 int `json:"47"`
	// }
	Num1002 orderedmap.OrderedMap `json:"1002"`
	Num1003 struct {
	} `json:"1003"`
	Num1005 struct {
	} `json:"1005"`
	BattleDamage struct {
	} `json:"battle_damage"`
}

type SMissionEndTrialExercise struct {
	// gun_with_user_id -> int
	GunsDps       orderedmap.OrderedMap `json:"guns_dps"`
	GunsRealLife  orderedmap.OrderedMap `json:"guns_real_life"`
	RewardVoucher types.Int             `json:"reward_voucher"`
}
