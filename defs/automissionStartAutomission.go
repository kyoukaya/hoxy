package defs


// CAutomission/startAutomission
type CAutomissionStartAutomission struct {
	TeamIds       []int `json:"team_ids"`
	AutoMissionID int   `json:"auto_mission_id"`
	Number        int   `json:"number"`
}

type SAutomissionStartAutomission int