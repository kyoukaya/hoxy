package defs

// Operation/finishOperation
type COperationFinishOperation struct {
	OperationID int `json:"operation_id"`
}

type SOperationFinishOperation struct {
	ItemID     string `json:"item_id"`
	BigSuccess int    `json:"big_success"`
}
