package types

import "bytes"

// Str prevents escaped unicode characters in the JSON from being turned into literal
// unicode characters. It also allows for a null value in which case it'll write the
// null back when marshalling.
type Str struct {
	Str    string
	IsNull bool
}

func (s *Str) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`null`)) {
		s.IsNull = true
		return nil
	}
	s.Str = string(b)
	return nil
}

func (s Str) MarshalJSON() ([]byte, error) {
	if s.IsNull && len(s.Str) == 0 {
		return []byte(`null`), nil
	}
	return []byte(s.Str), nil
}
