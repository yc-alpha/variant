package variant

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"time"
)

const (
	intSize = 32 << (^uint(0) >> 63)
	maxUint = 1<<intSize - 1
	maxInt  = 1<<(intSize-1) - 1
)

type Variant struct {
	Type   Kind
	Data   []byte
	layout string
}

var Nil = Variant{Type: Invalid}

func (v *Variant) SetLayout(layout string) *Variant {
	v.layout = layout
	return v
}

// Implement Stringer interface
func (v Variant) String() string {
	return fmt.Sprintf("Variant(%v, %v)", v.Type, v.Data)
}

// ToBytes converts the Variant to a byte slice.
func (v Variant) ToBytes() []byte {
	return v.Data
}

// ToBool converts the Variant to a boolean value.
func (v Variant) ToBool() bool {
	for _, ch := range v.Data {
		if ch != 0 {
			return true
		}
	}
	return false
}

// ToInt converts the Variant to an int value.
func (v Variant) ToInt() int {
	if fn := Strategies.int.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToInt64 converts the Variant to an int64 value.
func (v Variant) ToInt64() int64 {
	if fn := Strategies.int64.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToUint converts the Variant to an uint value.
func (v Variant) ToUint() uint {
	if fn := Strategies.uint.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToUint64 converts the Variant to an uint64 value.
func (v Variant) ToUint64() uint64 {
	if fn := Strategies.uint64.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToFloat32 converts the Variant to a float32 value.
func (v Variant) ToFloat32() float32 {
	if fn := Strategies.float32.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToFloat64 converts the Variant to a float64 value.
func (v Variant) ToFloat64() float64 {
	if fn := Strategies.float64.Get(v.Type); fn != nil {
		return fn(v)
	}
	return 0
}

// ToString converts the Variant to a string value.
func (v Variant) ToString() string {

	if fn := Strategies.string.Get(v.Type); fn != nil {
		return fn(v)
	}
	return ""
}

// ToTime converts the Variant to a time.Time value.
func (v Variant) ToTime() time.Time {
	if fn := Strategies.time.Get(v.Type); fn != nil {
		return fn(v)
	}
	return time.Time{}
}

// Equal checks if the Variant is equal to another value.
func (v Variant) Equal(other any) bool {
	if _, ok := other.(Variant); ok {
		return reflect.DeepEqual(v, other)
	}
	variant := New(other)
	variant.layout = v.layout
	return reflect.DeepEqual(v, variant)
}

func New(v any) Variant {
	variant := Variant{
		Type:   Invalid,
		layout: time.DateTime,
	}

	switch v := v.(type) {
	case *string:
		if v != nil {
			variant = New(*v)
		}
	case *bool:
		if v != nil {
			variant = New(*v)
		}
	case *int:
		if v != nil {
			variant = New(*v)
		}
	case *int8:
		if v != nil {
			variant = New(*v)
		}
	case *int16:
		if v != nil {
			variant = New(*v)
		}
	case *int32:
		if v != nil {
			variant = New(*v)
		}
	case *int64:
		if v != nil {
			variant = New(*v)
		}
	case *uint:
		if v != nil {
			variant = New(*v)
		}
	case *uint8:
		if v != nil {
			variant = New(*v)
		}
	case *uint16:
		if v != nil {
			variant = New(*v)
		}
	case *uint32:
		if v != nil {
			variant = New(*v)
		}
	case *uint64:
		if v != nil {
			variant = New(*v)
		}
	case *float32:
		if v != nil {
			variant = New(*v)
		}
	case *float64:
		if v != nil {
			variant = New(*v)
		}
	case *time.Time:
		if v != nil {
			variant = New(*v)
		}
	case string:
		variant.Type = String
		variant.Data = []byte(v)
	case bool:
		variant.Type = Bool
		if v {
			variant.Data = []byte{0x01}
		} else {
			variant.Data = []byte{0x00}
		}
	case int:
		variant.Type = Int
		switch intSize {
		case 32:
			variant.Data = make([]byte, 4)
			binary.BigEndian.PutUint32(variant.Data, uint32(v))
		case 64:
			variant.Data = make([]byte, 8)
			binary.BigEndian.PutUint64(variant.Data, uint64(v))
		}
	case int8:
		variant.Type = Int8
		variant.Data = append(variant.Data, byte(v))

	case int16:
		variant.Type = Int16
		variant.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(variant.Data, uint16(v))
	case int32:
		variant.Type = Int32
		variant.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(variant.Data, uint32(v))
	case int64:
		variant.Type = Int64
		variant.Data = make([]byte, 8)
		binary.BigEndian.PutUint64(variant.Data, uint64(v))
	case uint:
		variant.Type = Uint
		switch intSize {
		case 32:
			variant.Data = make([]byte, 4)
			binary.BigEndian.PutUint32(variant.Data, uint32(v))
		case 64:
			variant.Data = make([]byte, 8)
			binary.BigEndian.PutUint64(variant.Data, uint64(v))
		}
	case uint8:
		variant.Type = Uint8
		variant.Data = append(variant.Data, v)
	case uint16:
		variant.Type = Uint16
		variant.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(variant.Data, v)
	case uint32:
		variant.Type = Uint32
		variant.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(variant.Data, v)
	case uint64:
		variant.Type = Uint64
		variant.Data = make([]byte, 8)
		binary.BigEndian.PutUint64(variant.Data, v)
	case float32:
		variant.Type = Float32
		variant.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(variant.Data, math.Float32bits(v))
	case float64:
		variant.Type = Float64
		variant.Data = make([]byte, 8)
		binary.BigEndian.PutUint64(variant.Data, math.Float64bits(v))
	case time.Time:
		data, err := v.MarshalBinary()
		if err == nil {
			variant.Type = Time
			variant.Data = data
		}
	}
	return variant
}
