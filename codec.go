package variant

import (
	"bytes"
	"encoding/json"
)

type JSONMarshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONUnmarshaler interface {
	UnmarshalJSON([]byte) error
}

// type YAMLMarshaler interface {
// 	MarshalYAML() (any, error)
// }

// type YAMLUnmarshaler interface {
// 	UnmarshalYAML(node any) error
// }

// MarshalJSON implements the JSONMarshaler interface.
func (v Variant) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case String, Time:
		return json.Marshal(v.ToString())
	case Bool:
		return json.Marshal(v.ToBool())
	case Int, Int8, Int16, Int32, Int64:
		return json.Marshal(v.ToInt64())
	case Uint, Uint8, Uint16, Uint32, Uint64:
		return json.Marshal(v.ToUint64())
	case Float32, Float64:
		return json.Marshal(v.ToFloat64())
	default:
		return json.Marshal(v.Data)
	}
}

// UnmarshalJSON implements the JSONUnmarshaler interface.
func (v *Variant) UnmarshalJSON(data []byte) error {
	switch {
	case len(data) == 0: // defensive programming
		v.Type = Invalid
		v.Data = nil
	case data[0] == '"' && data[len(data)-1] == '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*v = New(s)
	case bytes.Equal(data, []byte("true")):
		*v = New(true)
	case bytes.Equal(data, []byte("false")):
		*v = New(false)
	case bytes.Equal(data, []byte("null")):
		v.Type = Invalid
		v.Data = nil
	default:
		v.Type = String
		v.Data = data
	}
	return nil
}

// func (v Variant) MarshalYAML() (any, error) {
// 	switch v.Type {
// 	case String, Time:
// 		return v.ToString(), nil
// 	case Bool:
// 		return v.ToBool(), nil
// 	case Int, Int8, Int16, Int32, Int64:
// 		return v.ToInt64(), nil
// 	case Uint, Uint8, Uint16, Uint32, Uint64:
// 		return v.ToUint64(), nil
// 	case Float32, Float64:
// 		return v.ToFloat64(), nil
// 	case Invalid:
// 		return nil, nil
// 	default:
// 		return v.Data, nil
// 	}
// }
