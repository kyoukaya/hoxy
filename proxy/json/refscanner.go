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
	EmptyKind RefKind = 0
	String    RefKind = 1 << iota
	Int
	Float
	Array
	Object
	ObjectKey
)

func (n RefKind) String() string {
	if n&String != 0 {
		switch n {
		case String | Int:
			return "string_int"
		case String | Float:
			return "string_float"
		}
	}
	switch n {
	case EmptyKind:
		return "empty"
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
	}
	panic(fmt.Sprintf("Undefined RefKind %#v", n))
}

type RefNode struct {
	key           string
	parent        *RefNode
	children      []*RefNode
	kind          RefKind
	precission    uint8
	isEmptyString bool
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

func (cur *RefNode) assignKind(buf *bytes.Buffer) error {
	if buf.Len() == 0 {
		return fmt.Errorf("Something went wrong")
	}
	if sl := buf.Bytes(); sl[0] == '"' && sl[len(sl)-1] == '"' {
		cur.kind = String
		if len(sl) == 2 {
			cur.isEmptyString = true
			return nil
		}
		// Trim quotation marks for number processing
		buf = bytes.NewBuffer(sl[1 : len(sl)-1])
	}

	isNumber := true
	isFp := false
	var precision uint8 = 0
	r, _, buffErr := buf.ReadRune()
	for ; buffErr == nil; r, _, buffErr = buf.ReadRune() {
		if isFp {
			precision++
		}
		if r < '0' || r > '9' {
			if r != '.' {
				isNumber = false
				break
			}
			if isFp {
				isNumber = false
				break
			}
			isFp = true
		}
	}

	if !isNumber {
		return nil
	}
	if isFp {
		cur.kind |= Float
		cur.precission = precision
		return nil
	}
	cur.kind |= Int
	return nil
}

// ParseRef parses input JSON data and returns a reference syntax tree.
func ParseRef(data []byte) (*RefNode, error) {
	scan := &refScanner{}
	root := &RefNode{}
	cur := root
	scan.reset()
	buf := bytes.Buffer{}
	for _, c := range data {
		scan.bytes++
		step := scan.step(&scan.scanner, c)
		if step == scanError {
			return root, scan.err
		}
		stepStr := scanStateString(step)
		fmt.Println(stepStr)
		switch step {
		case scanBeginLiteral:
			buf.Reset()
			buf.WriteByte(c)
		case scanContinue:
			buf.WriteByte(c)
		case scanObjectKey:
			key := bytes.Trim(buf.Bytes(), `"`)
			cur.key = string(key)
			cur.kind = ObjectKey
			cur = &RefNode{
				parent: cur,
			}
			fmt.Println(buf.String())
			buf.Reset()
		case scanArrayValue:
			if buf.Len() != 0 {
				cur.assignKind(&buf)
			}
			cur.parent.children = append(cur.parent.children, cur)
			cur = &RefNode{
				parent: cur.parent,
			}
			fmt.Println(buf.String())
			buf.Reset()
		case scanEndArray:
			// Could contain an object or array here
			if cur.kind == EmptyKind {
				cur.assignKind(&buf)
			}
			if buf.Len() > 0 || len(cur.children) > 0 {
				cur.parent.children = append(cur.parent.children, cur)
			}
			cur = cur.parent
			fmt.Println(buf.String())
			buf.Reset()
		case scanEndObject:
			if cur.kind == EmptyKind {
				cur.assignKind(&buf)
				fmt.Println(buf.String())
				buf.Reset()
			}
			// We can be in the scanEndObject state if the root value is an empty object
			// In that case, cur.parent would be nil as there would be no ObjectKey.
			if cur.parent != nil {
				break
			}
			rootObj := cur.parent.parent
			rootObj.parent.children = append(rootObj.parent.children, rootObj)
			cur = &RefNode{
				parent: rootObj.parent,
			}
		case scanObjectValue:
			if cur.kind == EmptyKind {
				cur.assignKind(&buf)
			}
			cur.parent.children = append(cur.parent.children, cur)
			cur = &RefNode{
				// This node's immediate parent is the ObjectKey node, so we take its
				// parent again to get to the actual object node.
				parent: cur.parent.parent,
			}
			fmt.Println(buf.String())
			buf.Reset()
		case scanBeginArray:
			cur.kind = Array
			new := &RefNode{
				parent: cur,
			}
			cur = new
		case scanBeginObject:
			cur.kind = Object
			new := &RefNode{
				parent: cur,
			}
			cur = new
		}
	}
	if scan.eof() == scanError {
		return root, scan.err
	}
	if buf.Len() != 0 {
		// assignKind when completed
		cur.kind = Int
	}
	return root, nil
}
