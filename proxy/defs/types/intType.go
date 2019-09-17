package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// Int is a custom type for handling ambiguous int or string encoded int data in json.
// Somtimes ints values are encoded as an empty string too.
type Int struct {
	Value       int
	FromStr     bool
	EmptyString bool
}

var emptyString = []byte(`""`)

func (ios *Int) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		if bytes.Equal(data, emptyString) {
			ios.EmptyString = true
			return nil
		}
		ios.FromStr = true
		i, err := strconv.Atoi(string(data[1 : len(data)-1]))
		if err != nil {
			return errors.New("json: cannot unmarshal to types.Int: " + err.Error())
		}
		ios.Value = i
		return nil
	}

	err := json.Unmarshal(data, &ios.Value)
	if err != nil {
		return err
	}
	return nil
}

func (ios *Int) MarshalJSON() ([]byte, error) {
	if ios.EmptyString {
		return []byte(`""`), nil
	}
	ret := strconv.Itoa(ios.Value)
	if ios.FromStr {
		ret = `"` + ret + `"`
	}
	return []byte(ret), nil
}
