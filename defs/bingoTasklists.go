package defs

// SBingo/tasklists
type SBingoTasklists []struct {
	TaskID            int    `json:"task_id"`
	TaskName          string `json:"task_name"`
	UserSize          int    `json:"user_size"`
	TaskSize          int    `json:"task_size"`
	Type              string `json:"type"`
	TicketNum         int    `json:"ticket_num"`
	Title             string `json:"title"`
	Content           string `json:"content"`
	FunctionControlID int    `json:"function_control_id"`
	TaskTypeID        int    `json:"task_type_id"`
	Size              int    `json:"size"`
}
