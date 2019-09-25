package defs

// Operation/startOperation
type COperationStartOperation struct {
	TeamID      int `json:"team_id"`
	OperationID int `json:"operation_id"`
	MaxLevel    int `json:"max_level"`
}

type SOperationStartOperation int
