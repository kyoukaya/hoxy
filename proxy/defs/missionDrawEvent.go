package defs

// Mission/drawEvent
type SMissionDrawEvent []struct {
	DrawEventID string      `json:"draw_event_id"`
	DrawNum     interface{} `json:"draw_num"`     // can be string or int
	DrawEndtime interface{} `json:"draw_endtime"` // can be string or int
}
