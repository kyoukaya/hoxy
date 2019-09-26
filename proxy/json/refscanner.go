package json

import (
	"bytes"
	"fmt"
)

type refScanner struct {
	scanner
}

type RefKind uint32

const (
	Int RefKind = iota
	Float
	String
	Array
	Object
	ObjectKey
)

func (n RefKind) String() string {
	switch n {
	case Int:
		return "int"
	case Float:
		return "float"
	case String:
		return "string"
	case Array:
		return "array"
	case Object:
		return "object"
	default:
		panic(fmt.Sprintf("Undefined RefKind %#v", n))
		return ""
	}
}

type RefNode struct {
	key      string
	parent   *RefNode
	children []*RefNode
	kind     RefKind
}

func scanStateString(i int) string {
	switch i {
	case scanContinue:
		return "scanContinue"
	case scanBeginLiteral:
		return "scanBeginLiteral"
	case scanBeginObject:
		return "scanBeginObject"
	case scanObjectKey:
		return "scanObjectKey"
	case scanObjectValue:
		return "scanObjectValue"
	case scanEndObject:
		return "scanEndObject"
	case scanBeginArray:
		return "scanBeginArray"
	case scanArrayValue:
		return "scanArrayValue"
	case scanEndArray:
		return "scanEndArray"
	case scanSkipSpace:
		return "scanSkipSpace"
	default:
		return ""
	}
}

func (cur *RefNode) assignKind(b []byte) {
	if len(b) == 0 {
		return
	}

	var isNumber, isFp bool
	s := string(b)
	isNumber = true
	isFp = false
	for _, r := range s {
		if r < '0' || r > '9' {
			if r != '.' {
				isNumber = false
				break
			}
			isFp = true
		}
	}
}

func ParseRef(data []byte) (*RefNode, error) {
	scan := &refScanner{}
	var root *RefNode
	var cur *RefNode
	scan.reset()
	accumulator := bytes.Buffer{}
	for _, c := range data {
		scan.bytes++
		step := scan.step(&scan.scanner, c)
		stepStr := scanStateString(step)
		fmt.Println(stepStr)
		switch step {
		case scanBeginLiteral:
			accumulator.Reset()
			accumulator.WriteByte(c)
		case scanContinue:
			accumulator.WriteByte(c)
		case scanObjectKey:
			new := &RefNode{
				parent: cur,
				kind:   ObjectKey,
			}
			key := bytes.Trim(accumulator.Bytes(), `"`)
			cur.key = string(key)
			cur.children = append(cur.children, new)
			cur = new
			fmt.Println(cur.key)
			accumulator.Reset()
		case scanArrayValue:
			if accumulator.Len() > 0 {
				new := &RefNode{
					parent: cur,
				}
				new.assignKind(accumulator.Bytes())
				cur.children = append(cur.children, new)
				fmt.Println(accumulator.String())
				accumulator.Reset()
			}
		case scanEndArray:
			if accumulator.Len() > 0 {
				new := &RefNode{
					parent: cur,
				}
				new.assignKind(accumulator.Bytes())
				cur.children = append(cur.children, new)
				cur = cur.parent
				fmt.Println(accumulator.String())
				accumulator.Reset()
			}
		case scanEndObject:
			if accumulator.Len() > 0 {
				cur.assignKind(accumulator.Bytes())
				fmt.Println(accumulator.String())
				accumulator.Reset()
				cur = cur.parent
			}
			cur = cur.parent
		case scanObjectValue:
			cur.assignKind(accumulator.Bytes())
			fmt.Println(accumulator.String())
			accumulator.Reset()
			cur = cur.parent
		case scanBeginArray:
			new := &RefNode{
				parent: cur,
				kind:   Array,
			}
			if cur == nil {
				root = new
			} else {
				cur.children = append(cur.children, new)
			}
			cur = new
		case scanBeginObject:
			new := &RefNode{
				parent: cur,
				kind:   Object,
			}
			if cur == nil {
				root = new
			} else {
				cur.children = append(cur.children, new)
			}
			cur = new
		}
		if step == scanError {
			return root, scan.err
		}
	}
	if scan.eof() == scanError {
		return root, scan.err
	}
	return root, nil
}
