package defs

type CIndexCrashLog struct {
	Context       string `json:"context"`
	Condition     string `json:"condition"`
	StackTrace    string `json:"stackTrace"`
	ClientVersion string `json:"clientVersion"`
}

type SIndexCrashLog int
