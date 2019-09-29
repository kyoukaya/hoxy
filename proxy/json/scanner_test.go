package json

import (
	"bytes"
	"testing"
)

func TestAssignKind(t *testing.T) {
	node := &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`""`)))
	if node.kind != String || !node.isEmptyString {
		t.Errorf("Expected empty string, got %#v", node)
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`"1"`)))
	if node.kind != Int|String {
		t.Errorf("Expected Int|String, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`1`)))
	if node.kind != Int {
		t.Errorf("Expected Int, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`"1.1"`)))
	if node.kind != Float|String && node.precission == 1 {
		t.Errorf("Expected Float|String, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`"1.11"`)))
	if node.kind != Float|String && node.precission == 2 {
		t.Errorf("Expected Float|String, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`"1."`)))
	if node.kind != Float|String && node.precission == 0 {
		t.Errorf("Expected Float|String, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`"1..1"`)))
	if node.kind != String {
		t.Errorf("Expected String, got %#v", node.kind.String())
	}
	node = &RefNode{}
	node.assignKind(bytes.NewBuffer([]byte(`1.11`)))
	if node.kind != Float && node.precission == 2 {
		t.Errorf("Expected Float|String, got %#v", node.kind.String())
	}
}

func TestLiteral(t *testing.T) {
	data := []byte(`1`)
	ParseRef(data)
}

func TestEmptySlice(t *testing.T) {
	data := []byte(`[]`)
	ParseRef(data)
}

func TestEmptyObj(t *testing.T) {
	data := []byte(`{}`)
	ParseRef(data)
}

func TestArrayWithLiterals(t *testing.T) {
	data := []byte(`[1,2,3]`)
	ParseRef(data)
}

func TestArrayWithObj(t *testing.T) {
	data := []byte(`[1,{"2":[3,4],"5":6,"7":8}, [9]]`)
	ParseRef(data)
}

func TestSliceWithSlice(t *testing.T) {
	data := []byte(`[11,[2,3],4]`)
	ParseRef(data)
}

func TestObjWithObj(t *testing.T) {
	data := []byte(`{"1":{}}`)
	ParseRef(data)
}

func TestObjEmptyKey(t *testing.T) {
	data := []byte(`{"":1}`)
	ParseRef(data)
}

func TestObjWithSlice(t *testing.T) {
	data := []byte(`{"1":[2]}`)
	ParseRef(data)
}
