package proxy

import (
	"hoxy/proxy/defs"
	"reflect"
	"strings"
	"sync"
)

// DefMapLock exposes a mutex to allow for safe concurrent access of DefMap.
var DefMapLock = &sync.Mutex{}

// DefMap provides a mapping of an op string to its struct type where possible.
// TODO: defer init
var DefMap = func() map[string]reflect.Type {
	literals := []interface{}{
		defs.CDormGet_build_coin{},
		defs.CDormReceive_favor{},
		defs.CEquipDevelopMulti{},
		defs.CEquipFinishDevelop{},
		defs.CFriendPraise{},
		defs.CFriendVisit{},
		defs.CGunChangeLock{},
		defs.CGunDevelopGun{},
		defs.CGunDevelopLog{},
		defs.CGunDevelopLogCollect{},
		defs.CGunDevelopMultiGun{},
		defs.CGunEatGun{},
		defs.CGunFinishAllDevelop{},
		defs.CGunFinishDevelop{},
		defs.CGunFixFinish{},
		defs.CGunFixGuns{},
		defs.CGunQuickDevelop{},
		defs.CGunRetireGun{},
		defs.CGunTeamGun{},
		defs.CIndexGetMailList{},
		defs.CIndexGetResourceInMail{},
		defs.CIndexHome{},
		defs.CIndexIndex{},
		defs.CIndexQuickGetQuestsResourceInMails{}, // TODO: fix these 2 defs. Their op strings are weird: Index/QuickGetQuestsResourceInMails
		defs.CMissionBattleFinish{},
		defs.CMissionCoinBattleFinish{},
		defs.CMissionFriendTeamAi{},
		defs.CMissionReinforceFriendTeam{},
		defs.CMissionStartMission{},
		defs.CMissionSupplyTeam{},
		defs.CMissionTeamMove{},
		defs.COperationFinishOperation{},
		defs.COperationStartOperation{},
		defs.COuthouseEstablish_build{},
		defs.SDormGet_build_coin{},
		defs.SDormKalinaFavor(0),
		defs.SDormReceive_favor{},
		defs.SDormShare{},
		defs.SEquipDevelopCollectList{},
		defs.SEquipDevelopLog{},
		defs.SEquipDevelopMulti{},
		defs.SEquipFinishDevelop{},
		defs.SFriendApplylist{},
		defs.SFriendDormInfo{},
		defs.SFriendList{},
		defs.SFriendMessagelist{},
		defs.SFriendPraise{},
		defs.SFriendRandomList{},
		defs.SFriendTeamGuns{},
		defs.SFriendVisit{},
		defs.SGunChangeLock(0),
		defs.SGunDevelopCollectList{},
		defs.SGunDevelopGun{},
		defs.SGunDevelopLog{},
		defs.SGunDevelopLogCollect(0),
		defs.SGunDevelopMultiGun{},
		defs.SGunEatGun{},
		defs.SGunFinishAllDevelop{},
		defs.SGunFinishDevelop{},
		defs.SGunFixFinish(0),
		defs.SGunFixGuns(0),
		defs.SGunTeamGun(0),
		defs.SIndexAttendance{},
		defs.SIndexGetMailList{},
		defs.SIndexGetResourceInMail{},
		defs.SIndexGetUidEnMicaQueue{},
		defs.SIndexHome{},
		defs.SIndexIndex{},
		defs.SIndexQuest{},
		defs.SIndexQuickGetQuestsResourceInMails{},
		defs.SIndexQuickGetResourceInMails{},
		defs.SIndexRecoverBp{},
		defs.SIndexVersion{},
		defs.SMissionBattleFinish{},
		defs.SMissionDrawEvent{},
		defs.SMissionEndEnemyTurn{},
		defs.SMissionEndTurn{},
		defs.SMissionFriendTeamMove{},
		defs.SMissionReinforceFriendTeam{},
		defs.SMissionStartMission{},
		defs.SMissionStartTurn{},
		defs.SMissionSupplyTeam(0),
		defs.SMissionTeamMove{},
		defs.SOperationFinishOperation{},
		defs.SOperationStartOperation(0),
		defs.SOuthouseEstablish_build{},
	}
	ret := make(map[string]reflect.Type)
	for _, def := range literals {
		t := reflect.TypeOf(def)
		op := defTypeToOpString(t)
		op = strings.ToLower(op)
		ret[op] = t
	}
	return ret
}()
