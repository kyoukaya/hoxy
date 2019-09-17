package proxy_test

import (
	"bytes"
	"fmt"
	"go/importer"
	. "hoxy/proxy"
	"hoxy/proxy/defs"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func compareOutput(orig, marshalled []byte) error {
	if !bytes.Equal(marshalled, orig) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(marshalled), string(orig), false)

		diff := dmp.PatchToText(dmp.PatchMake(diffs))
		return fmt.Errorf("marshalled JSON is not the same as original\n"+
			"Output: %s\nOrig:   %s\n%s", marshalled, orig, diff)
	}
	return nil
}

func testUnmarshal(t *testing.T, op string, orig []byte) interface{} {
	ret, err := UnMarshal(op, orig)
	if err != nil {
		t.Error(err)
		t.Errorf("%#v\n", ret)
	}
	return ret
}

func testMarshalAndCompare(t *testing.T, op string, pkt interface{}, orig []byte) {
	marshalled, err := Marshal(op, pkt)
	if err != nil {
		t.Error(err)
	}
	if err := compareOutput(orig, marshalled); err != nil {
		t.Error(err)
	}
}

func nameToOpString(typeName string) string {
	upCount := 0
	nameBuf := bytes.Buffer{}
	for _, v := range typeName {
		if unicode.IsUpper(v) {
			upCount++
			if upCount == 3 {
				nameBuf.WriteString("/")
				v = unicode.ToLower(v)
			}
		}
		nameBuf.WriteRune(v)
	}
	return nameBuf.String()
}

// Test that all defs declared in hoxy/proxy/defs are initialized in proxy.DefMap.
// Requires source.
func TestDefExists(t *testing.T) {
	pkg, err := importer.For("source", nil).Import("hoxy/proxy/defs")
	if err != nil {
		t.Errorf("error: %s\n", err.Error())
		return
	}
	exportedNames := make(map[string]bool)
	for _, name := range pkg.Scope().Names() {
		exportedNames[strings.ToLower(nameToOpString(name))] = true
	}
	// All exported structs should be in DefMap
	for name, _ := range exportedNames {
		_, ok := DefMap[name]
		if !ok {
			t.Errorf("%s in defs but not in DefMap", name)
		}
	}
	// All structs in DefMap should be exported by defs
	for name, _ := range DefMap {
		if !exportedNames[name] {
			t.Errorf("%s in DefMap but not in defs", name)
		}
	}
}

func loadFile(t *testing.T, fileName string) []byte {
	f, err := os.Open("testdata/" + fileName + ".json")
	if err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	return b
}

func TestGetUidEnMicaQueue(t *testing.T) {
	op := "SIndex/getUidEnMicaQueue"
	orig := loadFile(t, "SMicaQueue")
	ret := testUnmarshal(t, op, orig)
	micaQueue, ok := ret.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, micaQueue, orig)
}

func TestSIndexHome(t *testing.T) {
	op := "SIndex/home"
	orig := loadFile(t, "SIndexHome")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SIndexHome)
	if !ok {
		t.Error("Failed to cast")
	}
	testMarshalAndCompare(t, op, casted, orig)

	orig = loadFile(t, "SIndexHome.1")
	ret = testUnmarshal(t, op, orig)
	casted, ok = ret.(*defs.SIndexHome)
	if !ok {
		t.Error("Failed to cast")
	}
	testMarshalAndCompare(t, op, casted, orig)
}
func TestCMissionBattleFinish(t *testing.T) {
	op := "CMission/battleFinish"
	orig := loadFile(t, "CMissionBattleFinish")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.CMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSMissionBattleFinish(t *testing.T) {
	op := "SMission/battleFinish"
	orig := loadFile(t, "SMissionBattleFinish")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSMissionEndTurn(t *testing.T) {
	t.Log("Testing w/o MissionWinResult,ChangeBelong1,BuildingChangeBelong1")
	op := "SMission/endTurn"
	orig := loadFile(t, "SMissionEndTurn.1")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, casted, orig)
	t.Log("Testing w/ MissionWinResult,ChangeBelong1,BuildingChangeBelong1")
	// Test with missionwinresult
	orig = loadFile(t, "SMissionEndTurn.2")
	ret = testUnmarshal(t, op, orig)
	casted, ok = ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSFriendDormInfo(t *testing.T) {
	op := "SFriend/dormInfo"
	orig := loadFile(t, "SFriendDormInfo")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SFriendDormInfo)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSStartMission(t *testing.T) {
	// TODO: struct order sometimes changes
	// These 2 spot_ids have a different ordering on their structs, extracted from the
	// test json.
	// {"spot_id":"12","team_id":"1","belong":"1","if_random":"0","reinforce_count":"1","seed":2068,"enemy_team_id":"0","boss_hp":"0","enemy_hp_percent":"1","enemy_instance_id":"0","enemy_ai":0,"enemy_ai_para":"","ally_instance_ids":[],"squad_instance_ids":[],"hostage_id":"0","hostage_hp":"0","hostage_max_hp":"0","enemy_birth_turn":"999","supply_count":"0"},
	// {"spot_id":"9","team_id":"2","belong":"1","if_random":"0","reinforce_count":"1","seed":6348,"enemy_team_id":"0","boss_hp":"0","enemy_hp_percent":"1","enemy_instance_id":"0","enemy_ai":0,"enemy_ai_para":"","ally_instance_ids":[],"squad_instance_ids":[],"hostage_id":"0","hostage_hp":"0","hostage_max_hp":"0","enemy_birth_turn":"999","supply_count":"0"}
	op := "SMission/startMission"
	orig := loadFile(t, "SMissionStartMission")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionStartMission)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSIndexQuest(t *testing.T) {
	op := "SIndex/Quest"
	orig := loadFile(t, "SIndexQuest")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SIndexQuest)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSGunDevelopCollectList(t *testing.T) {
	op := "SGun/developCollectList"
	orig := loadFile(t, "SGunDevelopCollectList")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SGunDevelopCollectList)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSGunDevelopLog(t *testing.T) {
	op := "SGun/developLog"
	orig := loadFile(t, "SGunDevelopLog")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SGunDevelopLog)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSMissionReinforceFriendTeam(t *testing.T) {
	op := "SMission/reinforceFriendTeam"
	orig := loadFile(t, "SMissionReinforceFriendTeam")
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionReinforceFriendTeam)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}
