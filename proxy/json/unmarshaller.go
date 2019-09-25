package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/iancoleman/orderedmap"
	"github.com/mitchellh/mapstructure"
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
	Op   string
	data []byte
}

// MarshalMismatchErr is returned when the marshalled packet is unable to be marshalled
// back into its original []byte form.
type MarshalMismatchErr struct {
	Op        string
	original  []byte
	marshaled []byte
}

// MarshalFunc is a function that will perform json marshalling for ann interface
type MarshalFunc func(v interface{}) ([]byte, error)

// UnMarshalFunc is a function that will perform json unmarshalling into an interface
type UnMarshalFunc func(op string, data []byte) (interface{}, error)

func (e MarshalNoDefErr) Error() string {
	return fmt.Sprintf("MarshalNoDefErr: no defs found for %s\n%s", e.Op, e.data)
}

func (e MarshalMismatchErr) Error() string {
	return fmt.Sprintf("MarshalMismatchErr: mismatch on data for %s", e.Op)
}

var omType = reflect.TypeOf(orderedmap.OrderedMap{})

func processOrderedMap(m interface{}) (interface{}, error) {
	if m == nil {
		return nil, nil
	}

	if _, ok := m.(orderedmap.OrderedMap); ok {
		om := m.(orderedmap.OrderedMap)
		keys := om.Keys()
		for _, k := range keys {
			i, _ := om.Get(k)
			sub, err := processOrderedMap(i)
			if err != nil {
				return m, err
			}
			om.Set(k, sub)
		}
		return m, nil
	}

	t := reflect.TypeOf(m)
	k := t.Kind()
	switch k {
	case reflect.Int:
		return RefValue{k: reflect.Int}, nil
	case reflect.Float64:
		return RefValue{k: reflect.Float64}, nil
	case reflect.Slice:
		sl := reflect.ValueOf(m)
		var newSl []interface{}
		for i := 0; i < sl.Len(); i++ {
			sub, err := processOrderedMap(sl.Index(i).Interface())
			if err != nil {
				return newSl, err
			}
			newSl = append(newSl, sub)
		}
		return newSl, nil
	case reflect.String:
		return RefValue{k: reflect.String, isEmptyString: m.(string) == ""}, nil
	case reflect.Bool:
		return RefValue{k: reflect.Bool}, nil
	}
	return m, fmt.Errorf("unexpected kind: %#v of type %#v", k, t)
}

func decodeHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	// Coerce empty slice to empty struct if the expected kind is a struct.
	if from.Kind() == reflect.Slice && to.Kind() == reflect.Struct {
		return struct{}{}, nil
	}
	// Coerce empty string into an int
	if from.Kind() == reflect.String && to.Kind() == reflect.Int {
		fs := data.(string)
		if fs == "" {
			return 0, nil
		}
		return data, nil
	}
	return data, nil
}

func unMarshalfunc(op string, data []byte) (interface{}, MarshalFunc, error) {
	var ret interface{}
	var err error
	DefMapLock.Lock()
	t, ok := DefMap[strings.ToLower(op)]
	DefMapLock.Unlock()
	if ok {
		v := reflect.New(t)
		ret = v.Interface()

		if data[0] != '{' && data[0] != '[' {
			err = json.Unmarshal(data, &ret)
			return ret, json.Marshal, err
		}
		// hack: incredibly inefficient way to stop encoding/json from converting
		// unicode escape sequences into unicode runes.
		data = bytes.ReplaceAll(data, []byte(`\u`), []byte(`\\u`))

		// This whole block needs to be refactored as it violates DRY quite severely.

		// Hack to handle json arrays of values such as in CGun/retireGun
		if data[0] == '[' {
			if data[1] != '{' {
				err = json.Unmarshal(data, &ret)
				f := json.Marshal
				// Hack to make SEquip/adjust like packets where the empty array is
				// possible as a root value marshal back into the empty array.
				if bytes.Equal(data, []byte(`[]`)) {
					f = func(v interface{}) ([]byte, error) {
						return []byte(`[]`), nil
					}
				}
				return ret, f, err
			}
			tm := []orderedmap.OrderedMap{}
			err = json.Unmarshal(data, &tm)
			if err != nil {
				return ret, nil, err
			}
			_, err := processOrderedMap(tm)
			if err != nil {
				return ret, nil, err
			}
			m := []map[string]interface{}{}
			err = json.Unmarshal(data, &m)
			if err != nil {
				return ret, nil, err
			}
			conf := &mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				Result:           ret,
				TagName:          "json",
				ErrorUnused:      true,
				DecodeHook:       decodeHook,
			}
			dec, _ := mapstructure.NewDecoder(conf)
			err = dec.Decode(&m)
			marFunc := func(v interface{}) ([]byte, error) {
				return marshal(tm, v)
			}
			return ret, marFunc, err
		}

		tm := orderedmap.OrderedMap{}
		err = json.Unmarshal(data, &tm)
		if err != nil {
			return ret, nil, err
		}
		_, err := processOrderedMap(tm)
		if err != nil {
			return ret, nil, err
		}
		m := map[string]interface{}{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			return ret, nil, err
		}
		conf := &mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			Result:           ret,
			TagName:          "json",
			ErrorUnused:      true,
			DecodeHook:       decodeHook,
		}
		dec, _ := mapstructure.NewDecoder(conf)
		err = dec.Decode(&m)
		marFunc := func(v interface{}) ([]byte, error) {
			return marshal(tm, v)
		}
		return ret, marFunc, err
	}
	err = MarshalNoDefErr{op, data}
	return ret, nil, err
}

// UnMarshal takes json data in byte form and unmarshals it to a struct specified by
// the op string.
func UnMarshal(op string, data []byte) (interface{}, MarshalFunc, error) {
	var ret interface{}
	var err error

	if data == nil {
		return &SignCode{}, nil, nil
	}

	ret, marFunc, err := unMarshalfunc(op, data)

	// TODO: allow user to toggle marshalling validation.
	// compare UnMarshaled->json data with original data
	if err == nil {
		var remarshal []byte
		remarshal, err = marFunc(ret)
		if err != nil {
			return ret, marFunc, err
		}
		if !bytes.Equal(data, remarshal) {
			err = MarshalMismatchErr{op, data, remarshal}
		}
	}
	return ret, marFunc, err
}
