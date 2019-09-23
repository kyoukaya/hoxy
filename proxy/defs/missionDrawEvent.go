package defs

// Mission/drawEvent
type SMissionDrawEvent []struct {
	DrawEventID string `json:"draw_event_id"`
	DrawNum     int    `json:"draw_num"`
	DrawEndtime int    `json:"draw_endtime"`
}
