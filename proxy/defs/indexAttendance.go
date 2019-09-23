package defs

// CIndexAttendance Client's request for Index/attendance has no body.

// SIndexAttendance Server's response for Index/attendance
type SIndexAttendance struct {
	AttendanceType1 []struct {
		Day            string `json:"day"`
		StartTime      int    `json:"start_time"`
		EndTime        int    `json:"end_time"`
		AttendanceType int    `json:"attendance_type"`
		Mp             int    `json:"mp"`
		Ammo           int    `json:"ammo"`
		Mre            int    `json:"mre"`
		Part           int    `json:"part"`
		Gem            int    `json:"gem"`
		Core           int    `json:"core"`
		GunID          int    `json:"gun_id"`
		ItemIds        string `json:"item_ids"`
		Gift           string `json:"gift"`
		PrizeID        int    `json:"prize_id"`
	} `json:"attendance_type1"`
	AttendanceType2 []struct {
		Day            int    `json:"day"`
		StartTime      int    `json:"start_time"`
		EndTime        int    `json:"end_time"`
		AttendanceType int    `json:"attendance_type"`
		Mp             int    `json:"mp"`
		Ammo           int    `json:"ammo"`
		Mre            int    `json:"mre"`
		Part           int    `json:"part"`
		Gem            int    `json:"gem"`
		Core           int    `json:"core"`
		GunID          int    `json:"gun_id"`
		ItemIds        string `json:"item_ids"`
		Gift           string `json:"gift"`
		PrizeID        int    `json:"prize_id"`
	} `json:"attendance_type2"`
	Attendant1 struct {
		AttendanceType1Day  int `json:"attendance_type1_day"`
		AttendanceType1Time int `json:"attendance_type1_time"`
	} `json:"attendant1"`
	SevenAttendant string `json:"seven_attendant"`
	UserRecordUp   struct {
		SevenType           int    `json:"seven_type"`
		SevenStartTime      int    `json:"seven_start_time"`
		SevenAttendanceDays string `json:"seven_attendance_days"`
		SevenSpendPoint     int    `json:"seven_spend_point"`
	} `json:"user_record_up"`
	Attendant2 *struct {
		AttendanceType2Day  int `json:"attendance_type2_day"`
		AttendanceType2Time int `json:"attendance_type2_time"`
	} `json:"attendant2"`
}
