package defs

// CFriend/messageget
type CFriendMessageget struct {
	FUserid int `json:"f_userid"`
}

type SFriendMessageget struct {
	List []interface{} `json:"list"`
}
