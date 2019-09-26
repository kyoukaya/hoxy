package json_test

import (
	"testing"

	"github.com/kyoukaya/hoxy/proxy/json"
)

func TestLiteral(t *testing.T) {
	data := []byte(`1`)
	json.ParseRef(data)
}

func TestEmptySlice(t *testing.T) {
	data := []byte(`[]`)
	json.ParseRef(data)
}

func TestEmptyObj(t *testing.T) {
	data := []byte(`{}`)
	json.ParseRef(data)
}

func TestArrayWithLiterals(t *testing.T) {
	data := []byte(`[1,2,3]`)
	json.ParseRef(data)
}

func TestArrayWithObj(t *testing.T) {
	data := []byte(`[1,{"2":[3,4],"5":6}, [7]]`)
	json.ParseRef(data)
}

func TestSliceWithSlice(t *testing.T) {
	data := []byte(`[11,[2,3],4]`)
	json.ParseRef(data)
}

func TestObjWithObj(t *testing.T) {
	data := []byte(`{"1":{}}`)
	json.ParseRef(data)
}

func TestObjEmptyKey(t *testing.T) {
	data := []byte(`{"":1}`)
	json.ParseRef(data)
}

func TestObjWithSlice(t *testing.T) {
	data := []byte(`{"1":[2]}`)
	json.ParseRef(data)
}
