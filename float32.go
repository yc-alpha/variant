package variant

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"
	"unsafe"
)

var _ IConvertStrategy[float32] = (*float32Converter)(nil)

type float32Converter struct {
	Converter[float32]
}

func (c float32Converter) FromString(v Variant) float32 {
	s := *(*string)(unsafe.Pointer(&v.Data))
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(f64)
}

func (c float32Converter) FromBool(v Variant) float32 {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return 0
	}
	return 1
}

func (c float32Converter) FromInt(v Variant) float32 {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	return float32(i)
}

func (c float32Converter) FromInt8(v Variant) float32 {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	return float32(i)
}

func (c float32Converter) FromInt16(v Variant) float32 {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return float32(i)
}

func (c float32Converter) FromInt32(v Variant) float32 {
	return float32(int32(binary.BigEndian.Uint32(v.Data)))
}

func (c float32Converter) FromInt64(v Variant) float32 {
	return float32(int64(binary.BigEndian.Uint64(v.Data)))
}

func (c float32Converter) FromUint(v Variant) float32 {
	var i uint
	if intSize == 32 {
		i = uint(binary.BigEndian.Uint32(v.Data))
	} else if intSize == 64 {
		i = uint(binary.BigEndian.Uint64(v.Data))
	}
	return float32(i)
}

func (c float32Converter) FromUint8(v Variant) float32 {
	return float32(v.Data[0])
}

func (c float32Converter) FromUint16(v Variant) float32 {
	return float32(binary.BigEndian.Uint16(v.Data))
}

func (c float32Converter) FromUint32(v Variant) float32 {
	return float32(binary.BigEndian.Uint32(v.Data))
}

func (c float32Converter) FromUint64(v Variant) float32 {
	return float32(binary.BigEndian.Uint64(v.Data))
}

func (c float32Converter) FromFloat32(v Variant) float32 {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return f
}

func (c float32Converter) FromFloat64(v Variant) float32 {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return float32(f)
}

func (c float32Converter) FromTime(v Variant) float32 {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return 0
	}
	return float32(t.UnixNano())
}

func newFloat32Converter() IConvertStrategy[float32] {
	c := &float32Converter{}
	c.m = map[Kind]func(v Variant) float32{
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
