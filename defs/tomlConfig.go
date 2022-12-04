package defs

type (
	Config struct {
		Lists   map[string]Blocklist
		General GenInit
	}

	GenInit struct {
		BaseURL string
	}

	Blocklist struct {
		FromFile bool
		FileLoc  string
		List     string
	}
)
