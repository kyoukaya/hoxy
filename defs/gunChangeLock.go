package defs

// Gun/changeLock
type CGunChangeLock struct {
	Lock   []int         `json:"lock"`
	Unlock []interface{} `json:"unlock"`
}

type SGunChangeLock int