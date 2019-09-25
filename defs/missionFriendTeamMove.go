package defs

// SMission/friendTeamMove
type SMissionFriendTeamMove struct {
	FromSpotID             int           `json:"from_spot_id"`
	ToSpotID               int           `json:"to_spot_id"`
	FriendTeamID           int           `json:"friend_team_id"`
	AllFriendTeamMoved     int           `json:"all_friend_team_moved"`
	BuildingDefenderChange []interface{} `json:"building_defender_change"`
	Type5Score             string        `json:"type5_score"`
}
