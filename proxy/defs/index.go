package defs

// CIndexIndex Client's request for Index/index
type CIndexIndex struct {
	Time int `json:"time"`
}

// SIndexIndex Server's response for Index/index
type SIndexIndex struct {
	UserInfo struct {
		Gem                         string `json:"gem"`
		ID                          string `json:"id"`
		UserID                      string `json:"user_id"`
		OpenID                      string `json:"open_id"`
		ChannelID                   string `json:"channel_id"`
		Sign                        string `json:"sign"`
		Name                        string `json:"name"`
		GemPay                      string `json:"gem_pay"`
		PauseTurnChance             string `json:"pause_turn_chance"`
		Experience                  string `json:"experience"`
		Lv                          string `json:"lv"`
		Maxequip                    string `json:"maxequip"`
		Maxgun                      string `json:"maxgun"`
		Maxteam                     string `json:"maxteam"`
		MaxBuildSlot                string `json:"max_build_slot"`
		MaxEquipBuildSlot           string `json:"max_equip_build_slot"`
		MaxFixSlot                  string `json:"max_fix_slot"`
		MaxUpgradeSlot              string `json:"max_upgrade_slot"`
		MaxGunPreset                string `json:"max_gun_preset"`
		MaxFairy                    string `json:"max_fairy"`
		Bp                          string `json:"bp"`
		BpPay                       string `json:"bp_pay"`
		Mp                          string `json:"mp"`
		Ammo                        string `json:"ammo"`
		Mre                         string `json:"mre"`
		Part                        string `json:"part"`
		Core                        string `json:"core"`
		Coin1                       string `json:"coin1"`
		Coin2                       string `json:"coin2"`
		Coin3                       string `json:"coin3"`
		Monthlycard1EndTime         string `json:"monthlycard1_end_time"`
		Monthlycard2EndTime         string `json:"monthlycard2_end_time"`
		Monthlycard3EndTime         string `json:"monthlycard3_end_time"`
		Growthfond                  string `json:"growthfond"`
		LastRecoverTime             string `json:"last_recover_time"`
		LastBpRecoverTime           string `json:"last_bp_recover_time"`
		LastFavorRecoverTime        string `json:"last_favor_recover_time"`
		LastMonthlycard1RecoverTime string `json:"last_monthlycard1_recover_time"`
		LastMonthlycard2RecoverTime string `json:"last_monthlycard2_recover_time"`
		LastLoginTime               string `json:"last_login_time"`
		RegTime                     string `json:"reg_time"`
		GunCollect                  string `json:"gun_collect"`
		Maxdorm                     string `json:"maxdorm"`
		MaxdormSpare                string `json:"maxdorm_spare"`
		LastDormMaterialChangeTime1 string `json:"last_dorm_material_change_time1"`
		LastDormMaterialChangeTime2 string `json:"last_dorm_material_change_time2"`
		MaterialAvailableNum1       string `json:"material_available_num1"`
		MaterialAvailableNum2       string `json:"material_available_num2"`
		DormMaxScore                string `json:"dorm_max_score"`
		LastSsocChangeTime          string `json:"last_ssoc_change_time"`
		AppGuardID                  string `json:"app_guard_id"`
		AppGuardNum                 string `json:"app_guard_num"`
		IsBind                      string `json:"is_bind"`
	} `json:"user_info"`
	UpgradeActInfo []struct {
		UserID          string `json:"user_id"`
		FairyWithUserID string `json:"fairy_with_user_id"`
		GunWithUserID   string `json:"gun_with_user_id"`
		Type            string `json:"type"`
		Skill           string `json:"skill"`
		UpgradeSlot     string `json:"upgrade_slot"`
		EndTime         string `json:"end_time"`
	} `json:"upgrade_act_info"`
	GuideInfo struct {
		ID     string `json:"id"`
		UserID string `json:"user_id"`
		Guide  string `json:"guide"`
	} `json:"guide_info"`
	ItemWithUserInfo []struct {
		ItemID string `json:"item_id"`
		Number string `json:"number"`
		ID     string `json:"id,omitempty"`
	} `json:"item_with_user_info"`
	ItemLimitWithUserType9Info []struct {
		ItemID         string `json:"item_id"`
		DailyGet       string `json:"daily_get"`
		DailyClearTime string `json:"daily_clear_time"`
	} `json:"item_limit_with_user_type9_info"`
	MissionWithUserInfo []struct {
		ID                    string `json:"id"`
		UserID                string `json:"user_id"`
		MissionID             string `json:"mission_id"`
		Medal1                string `json:"medal1"`
		Medal2                string `json:"medal2"`
		Bestrank              string `json:"bestrank"`
		Medal4                string `json:"medal4"`
		Counter               string `json:"counter"`
		WinCounter            string `json:"win_counter"`
		ShortestInCoinmission string `json:"shortest_in_coinmission"`
		Type5Score            string `json:"type5_score"`
		IsOpen                string `json:"is_open"`
		IsDropDrawEvent       string `json:"is_drop_draw_event"`
		IsClose               string `json:"is_close"`
		CycleWinCount         string `json:"cycle_win_count"`
		MappedWinCounter      string `json:"mapped_win_counter"`
	} `json:"mission_with_user_info"`
	OperationActInfo []struct {
		ID          string `json:"id"`
		OperationID string `json:"operation_id"`
		UserID      string `json:"user_id"`
		TeamID      string `json:"team_id"`
		StartTime   string `json:"start_time"`
	} `json:"operation_act_info"`
	DevelopActInfo []interface{} `json:"develop_act_info"`
	FixActInfo     []interface{} `json:"fix_act_info"`
	MissionActInfo interface{}   `json:"mission_act_info"`
	SpotActInfo    interface{}   `json:"spot_act_info"`
	UserRecord     struct {
		UserID                         string      `json:"user_id"`
		MissionCampaign                string      `json:"mission_campaign"`
		SpecialMissionCampaign         string      `json:"special_mission_campaign"`
		AttendanceType1Day             int         `json:"attendance_type1_day"`
		AttendanceType1Time            int         `json:"attendance_type1_time"`
		AttendanceType2Day             string      `json:"attendance_type2_day"`
		AttendanceType2Time            string      `json:"attendance_type2_time"`
		DevelopLowrankCount            int         `json:"develop_lowrank_count"`
		SpecialDevelopLowrankCount     interface{} `json:"special_develop_lowrank_count"`
		GetGiftIds                     string      `json:"get_gift_ids"`
		SpendPoint                     string      `json:"spend_point"`
		GemMallIds                     string      `json:"gem_mall_ids"`
		ProductType21Ids               string      `json:"product_type21_ids"`
		SevenType                      string      `json:"seven_type"`
		SevenStartTime                 string      `json:"seven_start_time"`
		SevenAttendanceDays            string      `json:"seven_attendance_days"`
		SevenSpendPoint                string      `json:"seven_spend_point"`
		LastBpBuyTime                  string      `json:"last_bp_buy_time"`
		BpBuyCount                     string      `json:"bp_buy_count"`
		NewDevelopgunCount             string      `json:"new_developgun_count"`
		Buyitem1DevelopgunCount        float64     `json:"buyitem1_developgun_count"`
		Buyitem1SpecialdevelopgunCount int         `json:"buyitem1_specialdevelopgun_count"`
		Buyitem1Num                    string      `json:"buyitem1_num"`
		LastDevelopgunTime             int         `json:"last_developgun_time"`
		LastSpecialdevelopgunTime      int         `json:"last_specialdevelopgun_time"`
		FurnitureClasses               string      `json:"furniture_classes"`
		Adjutant                       string      `json:"adjutant"`
		AdjutantFairy                  string      `json:"adjutant_fairy"`
		MissionGroupTodayResetNum      int         `json:"mission_group_today_reset_num"`
		MissionGroupLastResetTime      string      `json:"mission_group_last_reset_time"`
		SpendpointAgeID                string      `json:"spendpoint_age_id"`
		SpendPointThismonth            string      `json:"spend_point_thismonth"`
		LastSpendpointTime             string      `json:"last_spendpoint_time"`
		NextCoreRecoverGunTime         string      `json:"next_core_recover_gun_time"`
		LastDatacellChangeTime         string      `json:"last_datacell_change_time"`
		BuympNum                       int         `json:"buymp_num"`
		BuyammoNum                     int         `json:"buyammo_num"`
		BuymreNum                      int         `json:"buymre_num"`
		BuypartNum                     int         `json:"buypart_num"`
		DevelopNonewNum                int         `json:"develop_nonew_num"`
	} `json:"user_record"`
	EquipWithUserInfo     interface{}   `json:"equip_with_user_info"`
	DevelopEquipActInfo   []interface{} `json:"develop_equip_act_info"`
	TeamInCoinMissionInfo struct {
		Num1 struct {
			UserID        string `json:"user_id"`
			Location      string `json:"location"`
			GunWithUserID string `json:"gun_with_user_id"`
			Position      string `json:"position"`
		} `json:"1"`
		Num2 struct {
			UserID        string `json:"user_id"`
			Location      string `json:"location"`
			GunWithUserID string `json:"gun_with_user_id"`
			Position      string `json:"position"`
		} `json:"2"`
		Num3 struct {
			UserID        string `json:"user_id"`
			Location      string `json:"location"`
			GunWithUserID string `json:"gun_with_user_id"`
			Position      string `json:"position"`
		} `json:"3"`
		Num4 struct {
			UserID        string `json:"user_id"`
			Location      string `json:"location"`
			GunWithUserID string `json:"gun_with_user_id"`
			Position      string `json:"position"`
		} `json:"4"`
		Num5 struct {
			UserID        string `json:"user_id"`
			Location      string `json:"location"`
			GunWithUserID string `json:"gun_with_user_id"`
			Position      string `json:"position"`
		} `json:"5"`
	} `json:"team_in_coin_mission_info"`
	SkinWithUserInfo interface{} `json:"skin_with_user_info"`
	EventInfo        struct {
		Num315 struct {
			ID        string `json:"id"`
			Code      string `json:"code"`
			Condition string `json:"condition"`
			StartTime int    `json:"start_time"`
			EndTime   int    `json:"end_time"`
		} `json:"315"`
	} `json:"event_info"`
	DormScoreInfo struct {
		Num1 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"1"`
		Num2 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"2"`
		Num3 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"3"`
		Num4 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"4"`
		Num5 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"5"`
		Num6 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"6"`
		Num7 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"7"`
		Num8 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"8"`
		Num9 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"9"`
		Num10 struct {
			UserID string `json:"user_id"`
			DormID string `json:"dorm_id"`
			Score  string `json:"score"`
		} `json:"10"`
	} `json:"dorm_score_info"`
	FurnitureWithUserInfo []struct {
		ID          string `json:"id"`
		Dorm        string `json:"dorm"`
		FurnitureID string `json:"furniture_id"`
		X           string `json:"x"`
		Y           string `json:"y"`
	} `json:"furniture_with_user_info"`
	FurnitureCollectInfo []struct {
		FurnitureID string `json:"furniture_id"`
		UserID      string `json:"user_id"`
	} `json:"furniture_collect_info"`
	FairyCollectInfo []struct {
		FairyID string `json:"fairy_id"`
		UserID  string `json:"user_id"`
	} `json:"fairy_collect_info"`
	TrialWithUserInfo struct {
		Counter  int    `json:"counter"`
		UserID   string `json:"user_id"`
		MaxTrial string `json:"max_trial"`
		MaxTime  string `json:"max_time"`
	} `json:"trial_with_user_info"`
	TrialExerciseInfo struct {
		GunsInfo             string `json:"guns_info"`
		BattleTeam           int    `json:"battle_team"`
		LastBattleFinishTime int    `json:"last_battle_finish_time"`
		Cheat                int    `json:"cheat"`
		TrialID              int    `json:"trial_id"`
		TeamIds              string `json:"team_ids"`
		UserID               string `json:"user_id"`
	} `json:"trial_exercise_info"`
	GunPresetInfo []struct {
		UserID   string `json:"user_id"`
		PresetNo string `json:"preset_no"`
		Gun1     string `json:"gun1"`
		Gun2     string `json:"gun2"`
		Gun3     string `json:"gun3"`
		Gun4     string `json:"gun4"`
		Gun5     string `json:"gun5"`
	} `json:"gun_preset_info"`
	OuthouseEstablishInfo []struct {
		ID                 string      `json:"id"`
		UserID             string      `json:"user_id"`
		RoomID             string      `json:"room_id"`
		EstablishID        string      `json:"establish_id"`
		EstablishType      string      `json:"establish_type"`
		FurnitureID        string      `json:"furniture_id"`
		UpgradeEstablishID string      `json:"upgrade_establish_id"`
		UpgradeStarttime   string      `json:"upgrade_starttime"`
		BuildStarttime     string      `json:"build_starttime"`
		BuildNum           string      `json:"build_num"`
		BuildTmpData       interface{} `json:"build_tmp_data"`
		BuildGetTime       string      `json:"build_get_time"`
		UpdateFurnitureID  string      `json:"update_furniture_id"`
		FurniturePostion   string      `json:"furniture_postion"`
		EstablishLv        string      `json:"establish_lv"`
		UpgradeCoin        string      `json:"upgrade_coin"`
		UpgradeTime        string      `json:"upgrade_time"`
		UpgradeCondition   string      `json:"upgrade_condition"`
		Parameter1         string      `json:"parameter_1"`
		Parameter2         string      `json:"parameter_2"`
		Parameter3         string      `json:"parameter_3"`
	} `json:"outhouse_establish_info"`
	DormRestFriendBuildCoinCount int         `json:"dorm_rest_friend_build_coin_count"`
	FairyWithUserInfo            interface{} `json:"fairy_with_user_info"`
	TipsInfo                     interface{} `json:"tips_info"`
	FairySkinWithUserInfo        interface{} `json:"fairy_skin_with_user_info"`
	UserGameSetting              struct {
		UID      string `json:"uid"`
		Settings string `json:"settings"`
	} `json:"user_game_setting"`
	MissionKeyWithUserInfo interface{} `json:"mission_key_with_user_info"`
	SquadDataDaily         []struct {
		UserID         int    `json:"user_id"`
		SquadID        string `json:"squad_id"`
		Type           string `json:"type"`
		LastFinishTime string `json:"last_finish_time"`
		Count          int    `json:"count"`
		Receive        int    `json:"receive"`
	} `json:"squad_data_daily"`
	SquadWithUserInfo   []interface{} `json:"squad_with_user_info"`
	ChipWithUserInfo    []interface{} `json:"chip_with_user_info"`
	DataAnalysisActInfo []interface{} `json:"data_analysis_act_info"`
	SquadTrainActInfo   []interface{} `json:"squad_train_act_info"`
	SquadSkillActInfo   []interface{} `json:"squad_skill_act_info"`
	SquadFixActInfo     []interface{} `json:"squad_fix_act_info"`
	SurplusDatacellInfo struct {
		UserID int `json:"user_id"`
		Base   int `json:"base"`
		Senior int `json:"senior"`
	} `json:"surplus_datacell_info"`
	GunMemoirList       []interface{} `json:"gun_memoir_list"`
	GunInTheaterInfo    []interface{} `json:"gun_in_theater_info"`
	SquadInTheaterInfo  []interface{} `json:"squad_in_theater_info"`
	FairyInTheaterInfo  []interface{} `json:"fairy_in_theater_info"`
	TheaterExerciseInfo interface{}   `json:"theater_exercise_info"`
	TheaterWithUserInfo []interface{} `json:"theater_with_user_info"`
	FriendWithUserInfo  []struct {
		Type          string      `json:"type"`
		EndTime       int         `json:"end_time"`
		FUserid       string      `json:"f_userid"`
		Name          string      `json:"name"`
		Lv            string      `json:"lv"`
		HeadpicID     interface{} `json:"headpic_id"`
		HomepageTime  interface{} `json:"homepage_time"`
		GiveTeamToday interface{} `json:"give_team_today"`
	} `json:"friend_with_user_info"`
	UserFriendInfo struct {
		FUserid int    `json:"f_userid"`
		Name    string `json:"name"`
		Lv      string `json:"lv"`
		Medal   []struct {
			ID    string `json:"id"`
			Level string `json:"level"`
			Num   string `json:"num"`
		} `json:"medal"`
		CardID    string      `json:"card_id"`
		HeadpicID interface{} `json:"headpic_id"`
		HomebgID  string      `json:"homebg_id"`
		Intro     string      `json:"intro"`
		TeamInfo  []struct {  // Contains all the user's gun information
			ID             string `json:"id"`
			UserID         string `json:"user_id"`
			GroupID        string `json:"group_id"`
			GunID          string `json:"gun_id"`
			GunExp         string `json:"gun_exp"`
			GunLevel       string `json:"gun_level"`
			Location       string `json:"location"`
			Position       string `json:"position"`
			Pow            string `json:"pow"`
			Hit            string `json:"hit"`
			Dodge          string `json:"dodge"`
			Rate           string `json:"rate"`
			Skill1         string `json:"skill1"`
			Skill2         string `json:"skill2"`
			Number         string `json:"number"`
			Equip1         string `json:"equip1"`
			Equip2         string `json:"equip2"`
			Equip3         string `json:"equip3"`
			Favor          string `json:"favor"`
			Skin           string `json:"skin"`
			SoulBond       string `json:"soul_bond"`
			IfModification string `json:"if_modification"`
			GunWithUserID  string `json:"gun_with_user_id"`
		} `json:"team_info"`
		BorrowTeamToday int `json:"borrow_team_today"`
		FairyInfo       []struct {
			ID              string `json:"id"`
			UserID          string `json:"user_id"`
			GroupID         string `json:"group_id"`
			FairyID         string `json:"fairy_id"`
			FairyLv         string `json:"fairy_lv"`
			FairyExp        string `json:"fairy_exp"`
			QualityLv       string `json:"quality_lv"`
			QualityExp      string `json:"quality_exp"`
			SkillLv         string `json:"skill_lv"`
			PassiveSkill    string `json:"passive_skill"`
			EquipID         string `json:"equip_id"`
			Skin            string `json:"skin"`
			FairyWithUserID string `json:"fairy_with_user_id"`
		} `json:"fairy_info"`
		IsReturnUser int `json:"is_return_user"`
	} `json:"user_friend_info"`
	GiftWithUserInfo []struct {
		ItemID string `json:"item_id"`
		Number string `json:"number"`
	} `json:"gift_with_user_info"`
	MedalWithUserInfo []struct {
		MedalID string `json:"medal_id"`
		Num     string `json:"num"`
	} `json:"medal_with_user_info"`
	KalinaWithUserInfo struct {
		UserID       string `json:"user_id"`
		ClickNum     int    `json:"click_num"`
		ClickTime    int    `json:"click_time"`
		Level        string `json:"level"`
		Favor        string `json:"favor"`
		LastFavor    string `json:"last_favor"`
		Skin         string `json:"skin"`
		SendMailTime string `json:"send_mail_time"`
	} `json:"kalina_with_user_info"`
	ShareWithUserInfo struct {
		LastTime int `json:"last_time"`
	} `json:"share_with_user_info"`
	AutoMissionActInfo interface{} `json:"auto_mission_act_info"`
	GunWithUserInfo    []struct {
		ID             string `json:"id"`
		UserID         string `json:"user_id"`
		GunID          string `json:"gun_id"`
		GunExp         string `json:"gun_exp"`
		GunLevel       string `json:"gun_level"`
		TeamID         string `json:"team_id"`
		IfModification string `json:"if_modification"`
		Location       string `json:"location"`
		Position       string `json:"position"`
		Life           string `json:"life"`
		Ammo           string `json:"ammo"`
		Mre            string `json:"mre"`
		Pow            string `json:"pow"`
		Hit            string `json:"hit"`
		Dodge          string `json:"dodge"`
		Rate           string `json:"rate"`
		Skill1         string `json:"skill1"`
		Skill2         string `json:"skill2"`
		FixEndTime     string `json:"fix_end_time"`
		IsLocked       string `json:"is_locked"`
		Number         string `json:"number"`
		Equip1         string `json:"equip1"`
		Equip2         string `json:"equip2"`
		Equip3         string `json:"equip3"`
		Equip4         string `json:"equip4"`
		Favor          string `json:"favor"`
		MaxFavor       string `json:"max_favor"`
		FavorToplimit  string `json:"favor_toplimit"`
		SoulBond       string `json:"soul_bond"`
		Skin           string `json:"skin"`
		CanClick       string `json:"can_click"`
		SoulBondTime   string `json:"soul_bond_time"`
	} `json:"gun_with_user_info"`
	ShowCdkey        int `json:"show_cdkey"`
	CanUsePayceo     int `json:"can_use_payceo"`
	Hexie            int `json:"hexie"`
	Nf               int `json:"nf"`
	IosAds           int `json:"ios_ads"`
	MissionEventInfo struct {
		ID                   string `json:"id"`
		MissionCampaign      string `json:"mission_campaign"`
		DrawEventID          string `json:"draw_event_id"`
		StartTime            int    `json:"start_time"`
		EndTime              int    `json:"end_time"`
		NormalCombatCampaign string `json:"normal_combat_campaign"`
		NormalCombatInit     string `json:"normal_combat_init"`
		DrawInfo             []struct {
			DrawEventID string      `json:"draw_event_id"`
			DrawNum     interface{} `json:"draw_num"` // str or int
			DrawEndtime string      `json:"draw_endtime"`
		} `json:"draw_info"`
		ItemLimitWithUser struct {
			Num8002 struct {
				ItemID         string `json:"item_id"`
				UserID         string `json:"user_id"`
				DailyGet       string `json:"daily_get"`
				DailyClearTime string `json:"daily_clear_time"`
				EventGet       string `json:"event_get"`
				EndTime        string `json:"end_time"`
			} `json:"8002"`
			Num100023 struct {
				ItemID         string `json:"item_id"`
				UserID         string `json:"user_id"`
				DailyGet       string `json:"daily_get"`
				DailyClearTime string `json:"daily_clear_time"`
				EventGet       string `json:"event_get"`
				EndTime        string `json:"end_time"`
			} `json:"100023"`
		} `json:"item_limit_with_user"`
		TomorrowUnix int `json:"tomorrow_unix"`
		AfterDays    int `json:"after_days"`
	} `json:"mission_event_info"`
	AuthenticationURL    string `json:"authentication_url"`
	NaiveBuildGunFormula string `json:"naive_build_gun_formula"`
	GameConfigInfo       struct {
		ShareGlobalSwitch    string `json:"share_global_switch"`
		SpecialDevelopSwitch string `json:"special_develop_switch"`
		EquipEnhanceSwitch   string `json:"equip_enhance_switch"`
		EquipRectifySwitch   string `json:"equip_rectify_switch"`
		TrailSwitch          string `json:"trail_switch"`
		SoulbindGlobalSwitch string `json:"soulbind_global_switch"`
	} `json:"game_config_info"`
}
