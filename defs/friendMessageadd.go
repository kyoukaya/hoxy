package defs

// CFriend/messageadd
type CFriendMessageadd struct {
	FUserid int    `json:"f_userid"`
	Message string `json:"message"`
}

type SFriendMessageadd struct {
	List []interface{} `json:"list"`
}
