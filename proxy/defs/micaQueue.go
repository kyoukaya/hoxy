package defs

// Index/getUidEnMicaQueue
type SIndexGetUidEnMicaQueue struct {
	UID               string `json:"uid"`
	Sign              string `json:"sign"`
	IsUsernameExist   bool   `json:"is_username_exist"`
	AppGuardID        string `json:"app_guard_id"`
	RealName          int    `json:"real_name"`
	AuthenticationURL string `json:"authentication_url"`
	// ErrMsg and ErrNo are only sent when the servers are down for maintenance.
	ErrMsg string `json:"Err_Msg"`
	ErrNo  int    `json:"Err_No"`
}
