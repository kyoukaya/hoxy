package types

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Int is a custom type for handling ambiguous int or string encoded int data in json.
type Int struct {
	Value   int
	fromStr bool
}

func (ioe *Int) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		ioe.fromStr = true
		i, err := strconv.Atoi(string(data[1 : len(data)-1]))
		if err != nil {
			return errors.New("types.Int.UnmarshallingJSON: " + err.Error())
		}
		ioe.Value = i
		return nil
	}

	err := json.Unmarshal(data, &ioe.Value)
	if err != nil {
		return err
	}
	return nil
}

func (ioe *Int) MarshalJSON() ([]byte, error) {
	ret := strconv.Itoa(ioe.Value)
	if ioe.fromStr {
		ret = `"` + ret + `"`
	}
	return []byte(ret), nil
}
