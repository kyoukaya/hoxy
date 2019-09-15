package defs

// CIndexAttendance Client's request for Index/attendance has no body.

// SIndexAttendance Server's response for Index/attendance
type SIndexAttendance struct {
	AttendanceType1 []struct {
		Day            string `json:"day"`
		StartTime      int    `json:"start_time"`
		EndTime        int    `json:"end_time"`
		AttendanceType string `json:"attendance_type"`
		Mp             string `json:"mp"`
		Ammo           string `json:"ammo"`
		Mre            string `json:"mre"`
		Part           string `json:"part"`
		Gem            string `json:"gem"`
		Core           string `json:"core"`
		GunID          string `json:"gun_id"`
		ItemIds        string `json:"item_ids"`
		Gift           string `json:"gift"`
		PrizeID        string `json:"prize_id"`
	} `json:"attendance_type1"`
	AttendanceType2 []interface{} `json:"attendance_type2"`
	Attendant1      struct {
		AttendanceType1Day  int `json:"attendance_type1_day"`
		AttendanceType1Time int `json:"attendance_type1_time"`
	} `json:"attendant1"`
	SevenAttendant string `json:"seven_attendant"`
	UserRecordUp   struct {
		SevenType           string `json:"seven_type"`
		SevenStartTime      string `json:"seven_start_time"`
		SevenAttendanceDays string `json:"seven_attendance_days"`
		SevenSpendPoint     string `json:"seven_spend_point"`
	} `json:"user_record_up"`
	Attendant2 []interface{} `json:"attendant2"`
}
