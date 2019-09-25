package json_test

import (
	"bytes"
	"fmt"
	"go/importer"
	"hoxy/defs"
	. "hoxy/proxy/json"
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

func testMarshalAndCompare(t *testing.T, marFunc MarshalFunc, data interface{}, orig []byte) {
	mar, err := marFunc(data)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(mar, orig) {
		t.Errorf("Marshalled output different from original\n"+
			"Original: %s\n"+
			"Output:   %s", orig, mar)
	}
}

// Test that all defs declared in hoxy/defs are initialized in proxy.DefMap.
// Requires source.
func TestDefExists(t *testing.T) {
	pkg, err := importer.For("source", nil).Import("hoxy/defs")
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
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexIndex(t *testing.T) {
	op := "SIndex/index"
	orig := loadFile(t, "SIndexIndex")
	ret, marFunc, err := UnMarshal(op, orig)
	if err != nil {
		t.Error(err)
	}
	casted, ok := ret.(*defs.SIndexIndex)
	if !ok {
		t.Error("Failed to cast")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexDownloadSuccess(t *testing.T) {
	op := "SIndex/downloadSuccess"
	orig := []byte(`1`)
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexDownloadSuccess)
	if !ok {
		t.Error("Failed to cast")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexGetMailList(t *testing.T) {
	op := "SIndex/getMailList"
	orig := []byte(`[]`)
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexGetMailList)
	if !ok {
		t.Error("Failed to cast")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexHome(t *testing.T) {
	op := "SIndex/home"
	orig := loadFile(t, "SIndexHome")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexHome)
	if !ok {
		t.Error("Failed to cast")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}
func TestCMissionBattleFinish(t *testing.T) {
	op := "CMission/battleFinish"
	orig := loadFile(t, "CMissionBattleFinish")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.CMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSMissionBattleFinish(t *testing.T) {
	op := "SMission/battleFinish"
	orig := loadFile(t, "SMissionBattleFinish")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
	orig = loadFile(t, "SMissionBattleFinish.1")
	ret, marFunc, err = UnMarshal(op, orig)
	casted, ok = ret.(*defs.SMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSMissionDrawEvent(t *testing.T) {
	op := "SMission/drawEvent"
	orig := loadFile(t, "SMissionDrawEvent")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SMissionDrawEvent)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexAttendance(t *testing.T) {
	op := "SIndex/attendance"
	orig := loadFile(t, "SIndexAttendance")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexAttendance)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSEquipAdjust(t *testing.T) {
	op := "SEquip/adjust"
	orig := []byte(`{"pow":10000,"hit":7500}`)
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SEquipAdjust)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
	orig = []byte(`[]`)
	ret, marFunc, err = UnMarshal(op, orig)
	casted, ok = ret.(*defs.SEquipAdjust)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSMissionEndTurn(t *testing.T) {
	op := "SMission/endTurn"
	orig := loadFile(t, "SMissionEndTurn.1")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
	orig = loadFile(t, "SMissionEndTurn.2")
	ret, marFunc, err = UnMarshal(op, orig)
	casted, ok = ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
	orig = loadFile(t, "SMissionEndTurn.3")
	ret, marFunc, err = UnMarshal(op, orig)
	casted, ok = ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestCGunRetireGun(t *testing.T) {
	op := "CGun/retireGun"
	orig := []byte(`[1,2,3,4]`)
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.CGunRetireGun)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		t.Error(err)
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSOuthouseEstablishBuild(t *testing.T) {
	op := "SOuthouse/establish_build"
	orig := loadFile(t, "SOuthouseEstablish_build")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SOuthouseEstablish_build)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		t.Error(err)
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSFriendList(t *testing.T) {
	op := "SFriend/list"
	orig := loadFile(t, "SFriendList")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SFriendList)
	if !ok {
		t.Error("Failed to cast.")
	}
	if err != nil {
		t.Error(err)
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func Benchmark(b *testing.B) {
	op := "SFriend/teamGuns"
	orig := loadFile(nil, "SFriendTeamGuns")
	UnMarshal(op, orig)
}

func TestSFriendDormInfo(t *testing.T) {
	op := "SFriend/dormInfo"
	orig := loadFile(t, "SFriendDormInfo")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SFriendDormInfo)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
	orig = loadFile(t, "SFriendDormInfo.1")
	ret, marFunc, err = UnMarshal(op, orig)
	casted, ok = ret.(*defs.SFriendDormInfo)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		t.Error(err)
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSStartMission(t *testing.T) {
	op := "SMission/startMission"
	orig := loadFile(t, "SMissionStartMission")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SMissionStartMission)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSIndexQuest(t *testing.T) {
	op := "SIndex/Quest"
	orig := loadFile(t, "SIndexQuest")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SIndexQuest)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSGunDevelopCollectList(t *testing.T) {
	op := "SGun/developCollectList"
	orig := loadFile(t, "SGunDevelopCollectList")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SGunDevelopCollectList)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSGunDevelopLog(t *testing.T) {
	op := "SGun/developLog"
	orig := loadFile(t, "SGunDevelopLog")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SGunDevelopLog)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}

func TestSMissionReinforceFriendTeam(t *testing.T) {
	op := "SMission/reinforceFriendTeam"
	orig := loadFile(t, "SMissionReinforceFriendTeam")
	ret, marFunc, err := UnMarshal(op, orig)
	casted, ok := ret.(*defs.SMissionReinforceFriendTeam)
	if !ok {
		t.Error(casted, ok)
	}
	if err != nil {
		testMarshalAndCompare(t, marFunc, casted, orig)
	}
}
