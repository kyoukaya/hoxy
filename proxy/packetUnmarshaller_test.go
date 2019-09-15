package proxy_test

import (
	"bytes"
	"fmt"
	"go/importer"
	. "hoxy/proxy"
	"hoxy/proxy/defs"
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
	t.Logf("%#v", pkt)
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

func TestGetUidEnMicaQueue(t *testing.T) {
	op := "SIndex/getUidEnMicaQueue"
	orig := []byte(`{"uid":"000","sign":"acb375d3c65ce5","is_username_exist":true,"app_guard_id":"9d644abc6c426c5691f71e:819:BA==","real_name":0,"authentication_url":"http:\/\/realauth.ucenter.ppgame.com\/authoriz.html?appid={0}&openid={1}&accounttype=1&language=zh"}`)
	ret := testUnmarshal(t, op, orig)
	micaQueue, ok := ret.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, micaQueue, orig)
}

func TestSIndexHome(t *testing.T) {
	op := "SIndex/home"
	orig := []byte(`{"recover_mp":0,"recover_ammo":0,"recover_mre":0,"recover_part":0,"gem":170,"all_favorup_gun":[{"gun_with_user_id":"1","favor_afteradd":567005},{"gun_with_user_id":"2","favor_afteradd":567572},{"gun_with_user_id":"3","favor_afteradd":567446}],"last_favor_recover_time":4,"recover_ssoc":0,"last_ssoc_change_time":5,"kick":0,"friend_messagelist":[],"friend_applylist":[{"adjutant":{"gun_id":"16","skin":"0","mod":"0","ai":"16"},"f_userid":7,"name":"H","lv":"10","headpic_id":"0","homepage_time":8,"is_return_user":0}],"index_getmaillist":[{"id":"47683600","user_id":"777000","type":"1","sub_id":"0","user_exp":"0","mp":"100","ammo":"100","mre":"100","part":"100","core":"0","gem":"0","gun_id":"0","fairy_ids":"","item_ids":"","equip_ids":"","furniture":"","gift":"","coins":"","skin":"0","bp_pay":"0","chip":"","title":"prize-10000301","content":"prize-20000301","code":"","start_time":"1567944000","end_time":"1568030400","if_read":"0"},{"id":"46403959","user_id":"0","type":"6","sub_id":"0","user_exp":"0","mp":"0","ammo":"0","mre":"0","part":"0","core":"0","gem":"0","gun_id":"0","fairy_ids":"","item_ids":"100023-60","equip_ids":"","furniture":"","gift":"","coins":"","skin":"0","bp_pay":"0","chip":"","title":"Extra 60 White Knight Armor Fragments","content":"Dear Commander,\nHere are the extra 60 White Knight Armor Fragments that can be used to exchange collaboration rewards from the event shop.\nPlease note that the event shop will be closed on Sep. 9th.","code":"","start_time":"1566894443","end_time":"1568102399","if_read":"0"},{"id":"47457078","user_id":"0","type":"6","sub_id":"0","user_exp":"0","mp":"0","ammo":"0","mre":"0","part":"0","core":"0","gem":"0","gun_id":"0","fairy_ids":"","item_ids":"","equip_ids":"","furniture":"","gift":"","coins":"","skin":"0","bp_pay":"0","chip":"","title":"New T-Dolls Rate-Up","content":"Dear Commander,\n\nThe pull rates of the 5 new T-Dolls: 5-star SMG T-Doll <color=#bf8f00>P90<\/color>, 5-star HG T-Doll <color=#bf8f00>Px4 Storm<\/color>, 5-star MG T-Doll <color=#bf8f00>QJY-88<\/color>, 4-star AR T-Doll <color=#bf8f00>SAR-21<\/color> and 4-star RF T-Doll <color=#bf8f00>K31<\/color> in the standard production pool are increased temporarily.\nTime: <color=#ff0000> Sep. 7th 00:00 – Sep. 8th 23:59 UTC-8 <\/color>","code":"245,0,0","start_time":"1567843200","end_time":"1568015999","if_read":"0"}],"squad_data_daily":[{"user_id":777000,"squad_id":"1","type":"mission:optional","last_finish_time":"2019-09-08","count":0,"receive":0}],"build_coin_flag":-4,"is_bind":"0"}`)
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SIndexHome)
	if !ok {
		t.Error("Failed to cast")
	}
	testMarshalAndCompare(t, op, casted, orig)
}
func TestCMissionBattleFinish(t *testing.T) {
	op := "CMission/battleFinish"
	orig := []byte(`{"spot_id":94,"if_enemy_die":true,"current_time":1,"boss_hp":0,"mvp":2,"last_battle_info":"","use_skill_squads":[],"guns":[{"id":3,"life":39},{"id":4,"life":37},{"id":2,"life":35}],"user_rec":"{\"seed\":1788716,\"record\":[]}","1000":{"10":430,"11":430,"12":430,"13":430,"15":132,"16":0,"17":265,"33":20001,"40":15,"18":0,"19":0,"20":0,"21":0,"22":0,"23":0,"24":189,"25":0,"26":189,"27":12,"34":1,"35":1,"41":12,"42":0,"43":0,"44":0,"92":0},"1001":{},"1002":{"2":{"47":0},"3":{"47":1},"4":{"47":1}},"1003":{},"1005":{},"battle_damage":{}}`)
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.CMissionBattleFinish)
	if !ok {
		t.Error("Failed to cast.")
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSMissionBattleFinish(t *testing.T) {
	op := "SMission/battleFinish"
	orig := []byte(`{"battle_get_gun":{"gun_with_user_id":"4","gun_id":"5"},"user_exp":"15","gun_exp":[{"gun_with_user_id":"1","exp":"150"},{"gun_with_user_id":"2","exp":"180"},{"gun_with_user_id":"3","exp":"195"}],"gun_life":[{"gun_with_user_id":"2","life":"38"}],"squad_exp":[],"battle_rank":"5","free_exp":17,"change_belong":[],"building_defender_change":[],"mission_win_result":[],"seed":3054,"favor_change":{"1":9,"2":30,"3":30},"type5_score":"0","ally_instance_transform":[],"ally_instance_betray":[],"mission_control":{"1":1,"2":2,"3":3}}`)
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionBattleFinish)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSMissionEndTurn(t *testing.T) {
	op := "SMission/endTurn"
	orig := []byte(`{"change_belong":{"95":2,"96":2,"98":2},"building_change_belong":[],"building_defender_change":[],"mission_lose_result":[],"type5_score":"0","fairy_skill_return":[],"fairy_skill_perform":[],"fairy_skill_on_spot":[],"fairy_skill_on_team":[],"fairy_skill_on_enemy":[],"fairy_skill_on_squad":[],"ally_instance_transform":[],"ally_instance_betray":[],"mission_control":{"1":1,"2":2,"3":3}}`)
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SMissionEndTurn)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}

func TestSIndexQuest(t *testing.T) {
	op := "SIndex/Quest"
	orig := []byte(`{"daily":{"operation":2,"fix":5,"mission":1,"borrow_friend_team":2,"eat_equip":"0","develop_equip":"0","develop_gun":"0","eat":"0","coin_mission":"0","from_friend_build_coin":"0","id":"9721","user_id":"102304","end_time":"1568534400","upgrade":"0","win_robot":"0","win_person":"0","win_boss":"0","win_armorrobot":"0","win_armorperson":"0","squad_data_analyse":"0"},"weekly":{"id":"9721","user_id":"102304","end_time":1568620800,"fix":"103","win_robot":"25","win_person":"203","win_boss":"10","win_armorrobot":0,"win_armorperson":0,"operation":"50","s_win":"100","eat":"20","develop_gun":"20","upgrade":"5","coin_mission":"19","develop_equip":"20","special_develop_gun":"1","adjust_equip":"5","eat_equip":"20","special_develop_equip":"3","adjust_fairy":"3","eat_fairy":0,"squad_data_analyse":0,"squad_eat_chip":0},"career":{"gun5_into_team":9438,"eat_equip":1620,"develop_equip":2924,"develop_gun":2429,"eat_gun":2645,"mission_duplicate_type3":887,"friend_visit":7393,"friend_praise":1732,"upgrade":482,"retire_gun":25214,"upgrade_fairy":55,"mission_duplicate_type2":279,"special_develop_gun":107,"develop_fairy":102,"trial_mission":137,"mission_duplicate_type1":290,"adjust_fairy":123,"adjust_equip":561,"adjust_equip_attr_100":309,"combine_gun_type_1":79,"combine_gun":355,"open_gift":11458,"eat_fairy":49,"combine_gun_type_5":32,"combine_gun_type_4":98,"auto_mission_1":"26","operation_5":"4","mission_emergency_1_1_1":"2","mission_emergency_3_1_1":"1","mission_emergency_3_2_1":"1","mission_emergency_3_3_1":"1","mission_emergency_3_4_1":"1","mission_emergency_3_5_1":"1","gun_equip_type_4":"47","gun5_into_team2":"314","retire_equip":"5","combine_gun_type_2":"85","combine_gun_type_3":"45","gasha_count":"453","dorm_change":"248","friend_card_change":"4","friend_headpic_change":"8","reinforce_team2":"1050","gun_equip_type_5":"62","team_from_1":"422","team_from_2":"55","team_from_3":"2","team_from_4":"3","team_from_5":"2","mission_emergency_1_2_1":"49","mission_emergency_1_3_1":"1","mission_emergency_1_4_1":"1","mission_emergency_1_5_1":"1","mission_emergency_1_6_1":"1","mission_emergency_1_7_1":"1","mission_emergency_1_8_1":"1","retire_fairy":"5","establish_up_type_301":"10","establish_up_type_302":"10","establish_up_type_303":"10","establish_up_type_304":"10","friend_apply":"139","sun_friend_team_into":"35","night_friend_team_into":"4"},"static_career_quest":[{"id":"1","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"auto_mission_1","count":"1","prize_id":"701","title":"career_quest-10000001","content":"career_quest-20000001","unlock_course":""},{"id":"2","unlock_lv":"","unlock_ids":"auto_mission_1:1","unlock_label":"","type":"operation_5","count":"1","prize_id":"702","title":"career_quest-10000002","content":"career_quest-20000002","unlock_course":""},{"id":"3","unlock_lv":"","unlock_ids":"operation_5:1","unlock_label":"","type":"reinforce_team2","count":"1","prize_id":"703","title":"career_quest-10000003","content":"career_quest-20000003","unlock_course":""},{"id":"4","unlock_lv":"","unlock_ids":"","unlock_label":"mission:15","type":"mission_emergency_1_1_1","count":"1","prize_id":"704","title":"career_quest-10000004","content":"career_quest-20000004","unlock_course":""},{"id":"5","unlock_lv":"","unlock_ids":"mission_emergency_1_1_1:1","unlock_label":"","type":"mission_emergency_1_2_1","count":"1","prize_id":"705","title":"career_quest-10000005","content":"career_quest-20000005","unlock_course":""},{"id":"6","unlock_lv":"","unlock_ids":"mission_emergency_1_2_1:1","unlock_label":"","type":"mission_emergency_1_3_1","count":"1","prize_id":"706","title":"career_quest-10000006","content":"career_quest-20000006","unlock_course":""},{"id":"7","unlock_lv":"","unlock_ids":"mission_emergency_1_3_1:1","unlock_label":"","type":"mission_emergency_1_4_1","count":"1","prize_id":"707","title":"career_quest-10000007","content":"career_quest-20000007","unlock_course":""},{"id":"8","unlock_lv":"","unlock_ids":"mission_emergency_1_4_1:1","unlock_label":"","type":"mission_emergency_1_5_1","count":"1","prize_id":"708","title":"career_quest-10000008","content":"career_quest-20000008","unlock_course":""},{"id":"9","unlock_lv":"","unlock_ids":"mission_emergency_1_5_1:1","unlock_label":"","type":"mission_emergency_1_6_1","count":"1","prize_id":"709","title":"career_quest-10000009","content":"career_quest-20000009","unlock_course":""},{"id":"10","unlock_lv":"","unlock_ids":"mission_emergency_1_6_1:1","unlock_label":"","type":"mission_emergency_1_7_1","count":"1","prize_id":"710","title":"career_quest-10000010","content":"career_quest-20000010","unlock_course":""},{"id":"11","unlock_lv":"","unlock_ids":"mission_emergency_1_7_1:1","unlock_label":"","type":"mission_emergency_1_8_1","count":"1","prize_id":"711","title":"career_quest-10000011","content":"career_quest-20000011","unlock_course":""},{"id":"12","unlock_lv":"","unlock_ids":"","unlock_label":"mission:90001","type":"mission_emergency_3_1_1","count":"1","prize_id":"712","title":"career_quest-10000012","content":"career_quest-20000012","unlock_course":""},{"id":"13","unlock_lv":"","unlock_ids":"mission_emergency_3_1_1:1","unlock_label":"","type":"mission_emergency_3_2_1","count":"1","prize_id":"713","title":"career_quest-10000013","content":"career_quest-20000013","unlock_course":""},{"id":"14","unlock_lv":"","unlock_ids":"mission_emergency_3_2_1:1","unlock_label":"","type":"mission_emergency_3_3_1","count":"1","prize_id":"714","title":"career_quest-10000014","content":"career_quest-20000014","unlock_course":""},{"id":"15","unlock_lv":"","unlock_ids":"mission_emergency_3_3_1:1","unlock_label":"","type":"mission_emergency_3_4_1","count":"1","prize_id":"715","title":"career_quest-10000015","content":"career_quest-20000015","unlock_course":""},{"id":"16","unlock_lv":"","unlock_ids":"mission_emergency_3_4_1:1","unlock_label":"","type":"mission_emergency_3_5_1","count":"1","prize_id":"716","title":"career_quest-10000016","content":"career_quest-20000016","unlock_course":""},{"id":"17","unlock_lv":"12","unlock_ids":"","unlock_label":"","type":"mission_duplicate_type1","count":"1","prize_id":"717","title":"career_quest-10000017","content":"career_quest-20000017","unlock_course":""},{"id":"18","unlock_lv":"12","unlock_ids":"","unlock_label":"","type":"mission_duplicate_type2","count":"1","prize_id":"718","title":"career_quest-10000018","content":"career_quest-20000018","unlock_course":""},{"id":"19","unlock_lv":"12","unlock_ids":"","unlock_label":"","type":"mission_duplicate_type3","count":"1","prize_id":"719","title":"career_quest-10000019","content":"career_quest-20000019","unlock_course":""},{"id":"20","unlock_lv":"","unlock_ids":"","unlock_label":"mission:90008","type":"trial_mission","count":"1","prize_id":"720","title":"career_quest-10000020","content":"career_quest-20000020","unlock_course":""},{"id":"21","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"gun5_into_team","count":"1","prize_id":"721","title":"career_quest-10000021","content":"career_quest-20000021","unlock_course":""},{"id":"22","unlock_lv":"","unlock_ids":"mission_emergency_3_1_1:1","unlock_label":"","type":"gun_equip_type_4","count":"1","prize_id":"722","title":"career_quest-10000022","content":"career_quest-20000022","unlock_course":""},{"id":"23","unlock_lv":"","unlock_ids":"mission_emergency_3_1_1:1","unlock_label":"","type":"gun_equip_type_5","count":"1","prize_id":"723","title":"career_quest-10000023","content":"career_quest-20000023","unlock_course":""},{"id":"24","unlock_lv":"","unlock_ids":"gun5_into_team:1","unlock_label":"","type":"gun5_into_team2","count":"1","prize_id":"724","title":"career_quest-10000024","content":"career_quest-20000024","unlock_course":""},{"id":"25","unlock_lv":"","unlock_ids":"gun5_into_team2:1","unlock_label":"","type":"team_from_1","count":"1","prize_id":"725","title":"career_quest-10000025","content":"career_quest-20000025","unlock_course":""},{"id":"26","unlock_lv":"","unlock_ids":"team_from_1:1","unlock_label":"","type":"team_from_2","count":"1","prize_id":"726","title":"career_quest-10000026","content":"career_quest-20000026","unlock_course":""},{"id":"27","unlock_lv":"","unlock_ids":"team_from_2:1","unlock_label":"","type":"team_from_3","count":"1","prize_id":"727","title":"career_quest-10000027","content":"career_quest-20000027","unlock_course":""},{"id":"28","unlock_lv":"","unlock_ids":"team_from_3:1","unlock_label":"","type":"team_from_4","count":"1","prize_id":"728","title":"career_quest-10000028","content":"career_quest-20000028","unlock_course":""},{"id":"29","unlock_lv":"","unlock_ids":"team_from_4:1","unlock_label":"","type":"team_from_5","count":"1","prize_id":"729","title":"career_quest-10000029","content":"career_quest-20000029","unlock_course":""},{"id":"50","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"develop_gun","count":"5","prize_id":"730","title":"career_quest-10000050","content":"career_quest-20000050","unlock_course":""},{"id":"51","unlock_lv":"","unlock_ids":"develop_gun:1","unlock_label":"","type":"combine_gun","count":"1","prize_id":"731","title":"career_quest-10000051","content":"career_quest-20000051","unlock_course":""},{"id":"52","unlock_lv":"","unlock_ids":"develop_gun:1","unlock_label":"","type":"eat_gun","count":"5","prize_id":"732","title":"career_quest-10000052","content":"career_quest-20000052","unlock_course":""},{"id":"53","unlock_lv":"","unlock_ids":"eat_gun:1","unlock_label":"","type":"retire_gun","count":"5","prize_id":"733","title":"career_quest-10000053","content":"career_quest-20000053","unlock_course":""},{"id":"54","unlock_lv":"","unlock_ids":"","unlock_label":"mission:20","type":"develop_equip","count":"5","prize_id":"734","title":"career_quest-10000054","content":"career_quest-20000054","unlock_course":""},{"id":"55","unlock_lv":"","unlock_ids":"","unlock_label":"mission:20","type":"retire_equip","count":"5","prize_id":"735","title":"career_quest-10000055","content":"career_quest-20000055","unlock_course":""},{"id":"56","unlock_lv":"","unlock_ids":"combine_gun:1","unlock_label":"","type":"combine_gun_type_4","count":"1","prize_id":"736","title":"career_quest-10000056","content":"career_quest-20000056","unlock_course":""},{"id":"57","unlock_lv":"","unlock_ids":"combine_gun:1","unlock_label":"","type":"combine_gun_type_2","count":"1","prize_id":"737","title":"career_quest-10000057","content":"career_quest-20000057","unlock_course":""},{"id":"58","unlock_lv":"","unlock_ids":"combine_gun:1","unlock_label":"","type":"combine_gun_type_5","count":"1","prize_id":"738","title":"career_quest-10000058","content":"career_quest-20000058","unlock_course":""},{"id":"59","unlock_lv":"","unlock_ids":"combine_gun:1","unlock_label":"","type":"combine_gun_type_3","count":"1","prize_id":"739","title":"career_quest-10000059","content":"career_quest-20000059","unlock_course":""},{"id":"60","unlock_lv":"","unlock_ids":"combine_gun:1","unlock_label":"","type":"combine_gun_type_1","count":"1","prize_id":"740","title":"career_quest-10000060","content":"career_quest-20000060","unlock_course":""},{"id":"61","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:30","type":"special_develop_gun","count":"1","prize_id":"741","title":"career_quest-10000061","content":"career_quest-20000061","unlock_course":""},{"id":"100","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"gasha_count","count":"1","prize_id":"742","title":"career_quest-10000100","content":"career_quest-20000100","unlock_course":""},{"id":"101","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"open_gift","count":"5","prize_id":"743","title":"career_quest-10000101","content":"career_quest-20000101","unlock_course":""},{"id":"102","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"dorm_change","count":"1","prize_id":"744","title":"career_quest-10000102","content":"career_quest-20000102","unlock_course":""},{"id":"103","unlock_lv":"3","unlock_ids":"","unlock_label":"","type":"friend_visit","count":"5","prize_id":"745","title":"career_quest-10000103","content":"career_quest-20000103","unlock_course":""},{"id":"104","unlock_lv":"3","unlock_ids":"friend_visit:1","unlock_label":"","type":"friend_praise","count":"5","prize_id":"746","title":"career_quest-10000104","content":"career_quest-20000104","unlock_course":""},{"id":"150","unlock_lv":"","unlock_ids":"","unlock_label":"mission:7","type":"friend_apply","count":"1","prize_id":"758","title":"career_quest-10000150","content":"career_quest-20000150","unlock_course":""},{"id":"151","unlock_lv":"","unlock_ids":"","unlock_label":"mission:7","type":"friend_card_change","count":"1","prize_id":"759","title":"career_quest-10000151","content":"career_quest-20000151","unlock_course":""},{"id":"152","unlock_lv":"","unlock_ids":"","unlock_label":"mission:7","type":"friend_headpic_change","count":"1","prize_id":"760","title":"career_quest-10000152","content":"career_quest-20000152","unlock_course":""},{"id":"153","unlock_lv":"12","unlock_ids":"","unlock_label":"","type":"upgrade","count":"5","prize_id":"761","title":"career_quest-10000153","content":"career_quest-20000153","unlock_course":""},{"id":"154","unlock_lv":"","unlock_ids":"","unlock_label":"mission:90004","type":"eat_equip","count":"5","prize_id":"762","title":"career_quest-10000154","content":"career_quest-20000154","unlock_course":""},{"id":"155","unlock_lv":"","unlock_ids":"","unlock_label":"mission:90008","type":"adjust_equip","count":"5","prize_id":"763","title":"career_quest-10000155","content":"career_quest-20000155","unlock_course":""},{"id":"156","unlock_lv":"","unlock_ids":"adjust_equip:1","unlock_label":"","type":"adjust_equip_attr_100","count":"1","prize_id":"764","title":"career_quest-10000156","content":"career_quest-20000156","unlock_course":""},{"id":"157","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"develop_fairy","count":"1","prize_id":"765","title":"career_quest-10000157","content":"career_quest-20000157","unlock_course":"80"},{"id":"158","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"eat_fairy","count":"5","prize_id":"766","title":"career_quest-10000158","content":"career_quest-20000158","unlock_course":"86,87,88"},{"id":"159","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"upgrade_fairy","count":"5","prize_id":"767","title":"career_quest-10000159","content":"career_quest-20000159","unlock_course":"86,87,88"},{"id":"160","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"adjust_fairy","count":"5","prize_id":"768","title":"career_quest-10000160","content":"career_quest-20000160","unlock_course":"86,87,88"},{"id":"161","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"retire_fairy","count":"5","prize_id":"769","title":"career_quest-10000161","content":"career_quest-20000161","unlock_course":"86,87,88"},{"id":"162","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"establish_up_type_301","count":"1","prize_id":"770","title":"career_quest-10000162","content":"career_quest-20000162","unlock_course":""},{"id":"163","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"establish_up_type_302","count":"1","prize_id":"771","title":"career_quest-10000163","content":"career_quest-20000163","unlock_course":""},{"id":"164","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"establish_up_type_303","count":"1","prize_id":"772","title":"career_quest-10000164","content":"career_quest-20000164","unlock_course":""},{"id":"165","unlock_lv":"","unlock_ids":"","unlock_label":"bestrank1:60","type":"establish_up_type_304","count":"1","prize_id":"773","title":"career_quest-10000165","content":"career_quest-20000165","unlock_course":""},{"id":"170","unlock_lv":"","unlock_ids":"","unlock_label":"","type":"sun_friend_team_into","count":"1","prize_id":"778","title":"career_quest-10000170","content":"career_quest-20000170","unlock_course":"100"},{"id":"171","unlock_lv":"","unlock_ids":"","unlock_label":"","type":"night_friend_team_into","count":"1","prize_id":"779","title":"career_quest-10000171","content":"career_quest-20000171","unlock_course":"100"}]}`)
	ret := testUnmarshal(t, op, orig)
	casted, ok := ret.(*defs.SIndexQuest)
	if !ok {
		t.Error(casted, ok)
	}
	testMarshalAndCompare(t, op, casted, orig)
}
