package defs

import "hoxy/proxy/defs/types"

// CIndexVersion has no body

// SIndexVersion server's response for Index/version
type SIndexVersion struct {
	Now               string    `json:"now"`
	TomorrowZero      string    `json:"tomorrow_zero"`
	MonthZero         int       `json:"month_zero"`
	NextMonthZero     int       `json:"next_month_zero"`
	Timezone          string    `json:"timezone"`
	DataVersion       string    `json:"data_version"`
	ClientVersion     string    `json:"client_version"`
	AbVersion         string    `json:"ab_version"`
	IsKick            string    `json:"is_kick"`
	Weekday           int       `json:"weekday"`
	AuthenticationURL types.Str `json:"authentication_url"`
}
