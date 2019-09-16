package defs

import "github.com/iancoleman/orderedmap"

// Gun/fixGuns
type CGunFixGuns struct {
	IfQuick int `json:"if_quick"`
	// Contains a key/value mapping of a gun_with_user_id to repair bay represented as an integer.
	FixGuns *orderedmap.OrderedMap `json:"fix_guns"`
}

type SGunFixGuns int
