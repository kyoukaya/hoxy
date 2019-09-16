package defs

import (
	"github.com/iancoleman/orderedmap"
)

// CIndexQuest has no request body

// SIndexQuest Server's Index/Quest response
type SIndexQuest struct {
	Daily             *orderedmap.OrderedMap `json:"daily"`
	Weekly            *orderedmap.OrderedMap `json:"weekly"`
	Career            *orderedmap.OrderedMap `json:"career"`
	StaticCareerQuest []struct {
		ID           string `json:"id"`
		UnlockLv     string `json:"unlock_lv"`
		UnlockIds    string `json:"unlock_ids"`
		UnlockLabel  string `json:"unlock_label"`
		Type         string `json:"type"`
		Count        string `json:"count"`
		PrizeID      string `json:"prize_id"`
		Title        string `json:"title"`
		Content      string `json:"content"`
		UnlockCourse string `json:"unlock_course"`
	} `json:"static_career_quest"`
}
