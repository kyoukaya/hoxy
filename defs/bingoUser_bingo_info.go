package defs

type SBingoUser_bingo_info struct {
	UserID                 int      `json:"user_id"`
	Grid                   []int    `json:"grid"`
	UserGrid               []string `json:"user_grid"`
	PayTaskRefreshCount    int      `json:"pay_task_refresh_count"`
	LinesIndex             []string `json:"lines_index"`
	BingoNumber            string   `json:"bingo_number"`
	EventStage             string   `json:"event_stage"`
	Createtime             int      `json:"createtime"`
	StartTime              int      `json:"start_time"`
	ID                     int      `json:"id"`
	DifficultyRefreshCount string   `json:"difficulty_refresh_count"`
	EndTime                int      `json:"end_time"`
}
