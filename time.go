package variant

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"
	"unsafe"
)

var _ IConvertStrategy[time.Time] = (*timeConverter)(nil)

type timeConverter struct {
	Converter[time.Time]
}

// FromBool implements IConvertStrategy.
func (t *timeConverter) FromBool(v Variant) time.Time {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return time.Time{}
	}
	return time.Now()
}

// FromFloat32 implements IConvertStrategy.
func (t *timeConverter) FromFloat32(v Variant) time.Time {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return time.Unix(0, int64(f))
}

// FromFloat64 implements IConvertStrategy.
func (t *timeConverter) FromFloat64(v Variant) time.Time {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return time.Unix(0, int64(f))
}

// FromInt implements IConvertStrategy.
func (t *timeConverter) FromInt(v Variant) time.Time {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	return time.Unix(0, int64(i))
}

// FromInt16 implements IConvertStrategy.
func (t *timeConverter) FromInt16(v Variant) time.Time {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return time.Unix(0, int64(i))
}

// FromInt32 implements IConvertStrategy.
func (t *timeConverter) FromInt32(v Variant) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint32(v.Data)))
}

// FromInt64 implements IConvertStrategy.
func (t *timeConverter) FromInt64(v Variant) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint64(v.Data)))
}

// FromInt8 implements IConvertStrategy.
func (t *timeConverter) FromInt8(v Variant) time.Time {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	return time.Unix(0, int64(i))
}

// FromString implements IConvertStrategy.
func (t *timeConverter) FromString(v Variant) time.Time {
	s := *(*string)(unsafe.Pointer(&v.Data))
	tt, err := time.Parse(v.layout, s)
	if err != nil {
		if i, e := strconv.Atoi(s); e == nil {
			return time.Unix(0, int64(i))
		}
	}
	return tt
}

// FromTime implements IConvertStrategy.
func (t *timeConverter) FromTime(v Variant) time.Time {
	var tt time.Time
	err := tt.UnmarshalBinary(v.Data)
	if err != nil {
		return time.Time{}
	}
	return tt
}

// FromUint implements IConvertStrategy.
func (t *timeConverter) FromUint(v Variant) time.Time {
	var i uint
	if intSize == 32 {
		i = uint(binary.BigEndian.Uint32(v.Data))
	} else if intSize == 64 {
		i = uint(binary.BigEndian.Uint64(v.Data))
	}
	return time.Unix(0, int64(i))
}

// FromUint16 implements IConvertStrategy.
func (t *timeConverter) FromUint16(v Variant) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint16(v.Data)))
}

// FromUint32 implements IConvertStrategy.
func (t *timeConverter) FromUint32(v Variant) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint32(v.Data)))
}

// FromUint64 implements IConvertStrategy.
func (t *timeConverter) FromUint64(v Variant) time.Time {
	return time.Unix(0, int64(binary.BigEndian.Uint64(v.Data)))
}

// FromUint8 implements IConvertStrategy.
func (t *timeConverter) FromUint8(v Variant) time.Time {
	return time.Unix(0, int64(v.Data[0]))
}

func newTimeConverter() IConvertStrategy[time.Time] {
	c := &timeConverter{}
	c.m = map[Kind]func(v Variant) time.Time{
		String:  c.FromString,
		Bool:    c.FromBool,
		Int:     c.FromInt,
		Int8:    c.FromInt8,
		Int16:   c.FromInt16,
		Int32:   c.FromInt32,
		Int64:   c.FromInt64,
		Uint:    c.FromUint,
		Uint8:   c.FromUint8,
		Uint16:  c.FromUint16,
		Uint32:  c.FromUint32,
		Uint64:  c.FromUint64,
		Float32: c.FromFloat32,
		Float64: c.FromFloat64,
		Time:    c.FromTime,
	}
	return c
}
