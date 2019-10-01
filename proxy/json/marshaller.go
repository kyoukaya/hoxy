package json

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/orderedmap"
)

// RefValue represents the type of data that was in the original JSON bytes.
type RefValue struct {
	// k can be an Int, Float64, String, Bool
	k             reflect.Kind
	isEmptyString bool
}

func isEmptyValue(data interface{}) (bool, error) {
	if data == nil {
		return true, nil
	}
	k := reflect.TypeOf(data).Kind()
	switch k {
	case reflect.Int:
		i := data.(int)
		return i == 0, nil
	case reflect.Float64:
		f := data.(float64)
		return f == 0., nil
	case reflect.Slice:
		return reflect.ValueOf(data).Len() == 0, nil
	case reflect.String:
		s := data.(string)
		return s == "", nil
	case reflect.Bool:
		b := data.(bool)
		return b, nil
	case reflect.Struct:
		return false, nil
	}
	return false, fmt.Errorf("marshal: unable to determine if %#v is empty, data: %#v", k, data)
}

func coerceDataIntoRef(ref RefValue, data interface{}, dataT reflect.Type) ([]byte, error) {
	dataK := dataT.Kind()
	switch dataK {
	// Numbers may be encoded in a string, e.g., "123", empty string, e.g., "", if the value is effectively nil, or as a JSON number, e.g., 123.
	case reflect.Int:
		if ref.isEmptyString {
			return []byte(`""`), nil
		}
		i := data.(int)
		is := strconv.Itoa(i)
		switch ref.k {
		case reflect.Int:
			return []byte(is), nil
		case reflect.String:
			return []byte(`"` + is + `"`), nil
		case reflect.Float64:
			return []byte(is), nil
		}
	case reflect.Float64:
		if ref.isEmptyString {
			return []byte(`""`), nil
		}
		f := data.(float64)
		fs := strconv.FormatFloat(f, 'f', 1, 64)
		fss := strconv.FormatFloat(f, 'f', -1, 64)
		if len(fs) < len(fss) {
			fs = fss
		}
		switch ref.k {
		case reflect.Float64:
			return []byte(fs), nil
		case reflect.String:
			return []byte(`"` + fs + `"`), nil
		}
	case reflect.String:
		s := data.(string)
		if ref.k == reflect.String {
			s = strings.Replace(s, "\n", `\n`, -1)
			s = strings.Replace(s, "\t", `\t`, -1)
			s = strings.Replace(s, `/`, `\/`, -1)
			s = strings.Replace(s, `"`, `\"`, -1)
			return []byte(`"` + s + `"`), nil
		}
		// Allow strings to be coerced into numbers as well
		if ref.k == reflect.Int || ref.k == reflect.Float64 {
			// validate that the data is a number
			for _, v := range s {
				if (v < '0' || v > '9') && v != '.' {
					return []byte(`0`), fmt.Errorf("unable to coerce string '%s' into JSON number", s)
				}
			}
			return []byte(s), nil
		}
	case reflect.Bool:
		b := data.(bool)
		switch ref.k {
		case reflect.Bool:
			if b {
				return []byte(`true`), nil
			}
			return []byte(`false`), nil
		case reflect.String:
			if b {
				return []byte(`"1"`), nil
			}
			return []byte(`"0"`), nil
		}
	}
	return nil, fmt.Errorf("marshal: unable to coerce data value %#v of kind %s into kind %s", data, dataK.String(), ref.k.String())
}

func marshal(ref, data interface{}) ([]byte, error) {
	var ret []byte
	var err error
	// If the reference value is nil, then the original JSON value is null.
	if ref == nil {
		isEmpty, err := isEmptyValue(data)
		if err != nil {
			return []byte{}, err
		}
		if isEmpty {
			return []byte("null"), nil
		}
		return []byte("null"), fmt.Errorf("marshal: expected empty value but got %#v", data)
	}

	// If the reference value is not nil, then it can either be slice, a RefValue struct
	// signalling a "leaf value", or a OrderedMap struct which means the data is either a map
	// or a struct.
	switch reflect.TypeOf(ref).Kind() {
	case reflect.Slice:
		refSl := reflect.ValueOf(ref)
		dataSl := reflect.ValueOf(data)
		// Dereference the slice if it's a ptr
		if dataSl.Kind() == reflect.Ptr {
			dataSl = dataSl.Elem()
		}
		// we could assert the lengths of the slices here
		refLen := refSl.Len()
		var buf bytes.Buffer
		buf.WriteString("[")

		for i := 0; i < refLen; i++ {
			refV := refSl.Index(i).Interface()
			dataV := dataSl.Index(i).Interface()
			var sub []byte
			sub, err = marshal(refV, dataV)
			buf.Write(sub)
			buf.WriteString(",")
			if err != nil {
				// Make sure we close off the slice even if we incur an error while
				// marshalling the elements within the slice.
				break
			}
		}
		// If we've written anything to the buffer other than the opening "["
		// then we have to strip the last byte off as it contains a "," char.
		if bufLen := buf.Len(); bufLen > 1 {
			buf.Truncate(bufLen - 1)
		}

		buf.WriteString("]")
		return buf.Bytes(), err
	case reflect.Struct:
		dataVal := reflect.ValueOf(data)
		dataT := reflect.TypeOf(data)

		switch ref.(type) {
		case RefValue:
			// Coerce data into RefValue
			return coerceDataIntoRef(ref.(RefValue), data, dataT)
		case orderedmap.OrderedMap:
			// If the reference type is an ordered map, the data must be a
			// map[string]T or a struct.
			om, _ := ref.(orderedmap.OrderedMap)

			// Reduce the map[string]T or struct into a map[string]interface
			dataMap := make(map[string]interface{})
			switch dataVal.Kind() {
			case reflect.Map:
				switch dataT.Key().Kind() {
				case reflect.Int:
					for _, k := range dataVal.MapKeys() {
						dataMap[strconv.FormatInt(k.Int(), 10)] = dataVal.MapIndex(k).Interface()
					}
				case reflect.String:
					for _, k := range dataVal.MapKeys() {
						dataMap[k.String()] = dataVal.MapIndex(k).Interface()
					}
				default:
					// Only allow maps to have ints or strings as keys.
					return []byte("{}"), fmt.Errorf("marshal: unsupported map type in "+
						"definitions %#v", dataVal.Kind())
				}
			case reflect.Struct:
				for i := 0; i < dataT.NumField(); i++ {
					field := dataT.Field(i)
					value := dataVal.FieldByIndex([]int{i}).Interface()
					dataMap[field.Tag.Get("json")] = value
				}
			case reflect.Ptr:
				pValue := dataVal.Elem()
				// Unroll ptr
				return marshal(ref, pValue.Interface())
			default:
				return []byte("null"), fmt.Errorf("marshal: unexpected kind in definition %s, when an OrderedMap is in the reference", dataT.Kind().String())
			}

			// Iterate over the reference keys in order
			keys := om.Keys()
			buf := bytes.Buffer{}
			buf.WriteString(`{`)
			for _, k := range keys {
				buf.WriteString(`"` + k + `":`)
				refV, _ := om.Get(k)
				dataV, ok := dataMap[k]
				if !ok {
					return nil, fmt.Errorf("marshal: key '%s' in reference not found in data", k)
				}
				var sub []byte
				sub, err = marshal(refV, dataV)
				buf.Write(sub)
				buf.WriteString(`,`)
				if err != nil {
					err = fmt.Errorf("marshal: error processing field %s: %s", k, err.Error()[9:])
					break
				}
			}
			if bufLen := buf.Len(); bufLen > 1 {
				buf.Truncate(bufLen - 1)
			}
			buf.WriteString(`}`)
			return buf.Bytes(), err
		}
	}
	// this shouldn't happen
	return ret, fmt.Errorf("marshal: unknown error on data:%#v\nref:%#v", data, ref)
}
