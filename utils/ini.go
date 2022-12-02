package utils

import (
	"gopkg.in/ini.v1"
)

func INIParsefromByte(data []byte, sect string, key string, typ string) any {
	cfg, err := ini.Load(data)

	if err != nil {
		panic(err)
	}

	switch typ {
	case "string":
		return cfg.Section(sect).Key(key).String()

	case "bool":
		ret, err := cfg.Section(sect).Key(key).Bool()
		if err != nil {
			panic(err)
		}
		return ret

	case "int":
		ret, err := cfg.Section(sect).Key(key).Int()
		if err != nil {
			panic(err)
		}
		return ret
	}

	return nil
}

func INIParsefromFile(floc string, sect string, key string, typ string) any {
	cfg, err := ini.Load(floc)

	if err != nil {
		panic(err)
	}

	switch typ {
	case "string":
		return cfg.Section(sect).Key(key).String()

	case "[]string":
		return cfg.Section(sect).Key(key).Strings(" ")

	case "bool":
		ret, err := cfg.Section(sect).Key(key).Bool()
		if err != nil {
			panic(err)
		}
		return ret

	case "int":
		ret, err := cfg.Section(sect).Key(key).Int()
		if err != nil {
			panic(err)
		}
		return ret
	}

	return nil
}
