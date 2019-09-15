package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// Returns the operation string of a properly named def.
// I.e., defs.SMissionEndTurn -> SMission/endTurn
func defTypeToOpString(t reflect.Type) string {
	upCount := 0
	typeName := t.Name()
	nameBuf := bytes.Buffer{}
	nameBuf.Reset()
	for _, v := range typeName {
		if unicode.IsUpper(v) {
			upCount++
			if upCount == 3 {
				nameBuf.WriteString("/")
				v = unicode.ToLower(v)
			}
		}
		nameBuf.WriteRune(v)
	}
	return nameBuf.String()
}

// SignCode is an empty struct representing a empty packet which only sends its signcode.
type SignCode struct{}

// MarshalNoDefErr is returned when the packet cannot be unmarshalled into a struct
// due to the lack of a definition.
type MarshalNoDefErr struct {
	op   string
	data []byte
}

// MarshalMismatchErr is returned when the marshalled packet is unable to be marshalled
// back into its original []byte form.
type MarshalMismatchErr struct {
	op        string
	original  []byte
	marshaled []byte
}

// MarshalFunc is a function that will perform json marshalling for ann interface
type MarshalFunc func(op string, v interface{}) ([]byte, error)

// UnMarshalFunc is a function that will perform json unmarshalling into an interface
type UnMarshalFunc func(op string, data []byte) (interface{}, error)

func (e MarshalNoDefErr) Error() string {
	return fmt.Sprintf("MarshalNoDefErr: no defs found for %s\n%s", e.op, e.data)
}

func (e MarshalMismatchErr) Error() string {
	return fmt.Sprintf("MarshalMismatchErr: mismatch on data for %s", e.op)
}

func unMarshalfunc(op string, data []byte) (interface{}, error) {
	var ret interface{}
	var err error
	DefMapLock.Lock()
	t, ok := DefMap[strings.ToLower(op)]
	DefMapLock.Unlock()
	if ok {
		v := reflect.New(t)
		ret = v.Interface()
		err = json.Unmarshal(data, ret)
	} else {
		err = MarshalNoDefErr{op, data}
	}
	return ret, err
}

// UnMarshal takes json data in byte form and unmarshals it to a struct specified by
// the op string.
func UnMarshal(op string, data []byte) (interface{}, error) {
	var ret interface{}
	var err error

	if data == nil {
		return &SignCode{}, nil
	}

	// TODO: define custom unmarshalling routine and defs for packets that break the mould
	// potentially return a function to marshal back into bytes.
	DefMapLock.Lock()
	_, ok := DefMap[strings.ToLower(op)]
	DefMapLock.Unlock()
	var unMarFunc UnMarshalFunc
	var marFunc MarshalFunc
	if ok {
		marFunc = Marshal
		unMarFunc = unMarshalfunc
	} else {
		err = MarshalNoDefErr{op, data}
	}
	if err != nil {
		return nil, err
	}

	ret, err = unMarFunc(op, data)

	// compare UnMarshaled->json data with original data
	if err == nil {
		remarshal, err := marFunc(op, ret)
		if err != nil {
			return ret, err
		}
		// TODO: maps in go aren't ordered so map[string]interfaces typically cause
		// mistmatch warnings. Probably not much we can do besides writing a new json
		// library.
		if !bytes.Equal(data, remarshal) {
			err = MarshalMismatchErr{op, data, remarshal}
		}
	}
	return ret, err
}

// Marshal takes in a struct representing the json data and unmarshals it
// back into a slice of bytes.
func Marshal(op string, v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	ret := buffer.Bytes()

	// Hack to get byte perfect return
	ret = bytes.ReplaceAll(ret, []byte(`/`), []byte(`\/`))
	// TODO: Find a better way of escaping unicode characters
	// Json encoder inserts an additional newline at the end
	ret = bytes.TrimRight(ret, "\n")
	return ret, err
}
