package variant

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"
	"unsafe"
)

var _ IConvertStrategy[float64] = (*float64Converter)(nil)

type float64Converter struct {
	Converter[float64]
}

func (c float64Converter) FromString(v Variant) float64 {
	s := *(*string)(unsafe.Pointer(&v.Data))
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f64
}

func (c float64Converter) FromBool(v Variant) float64 {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return 0
	}
	return 1
}

func (c float64Converter) FromInt(v Variant) float64 {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	return float64(i)
}

func (c float64Converter) FromInt8(v Variant) float64 {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	return float64(i)
}

func (c float64Converter) FromInt16(v Variant) float64 {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return float64(i)
}

func (c float64Converter) FromInt32(v Variant) float64 {
	return float64(int32(binary.BigEndian.Uint32(v.Data)))
}

func (c float64Converter) FromInt64(v Variant) float64 {
	return float64(int64(binary.BigEndian.Uint64(v.Data)))
}

func (c float64Converter) FromUint(v Variant) float64 {
	var i uint
	if intSize == 32 {
		i = uint(binary.BigEndian.Uint32(v.Data))
	} else if intSize == 64 {
		i = uint(binary.BigEndian.Uint64(v.Data))
	}
	return float64(i)
}

func (c float64Converter) FromUint8(v Variant) float64 {
	return float64(v.Data[0])
}

func (c float64Converter) FromUint16(v Variant) float64 {
	return float64(binary.BigEndian.Uint16(v.Data))
}

func (c float64Converter) FromUint32(v Variant) float64 {
	return float64(binary.BigEndian.Uint32(v.Data))
}

func (c float64Converter) FromUint64(v Variant) float64 {
	return float64(binary.BigEndian.Uint64(v.Data))
}

func (c float64Converter) FromFloat32(v Variant) float64 {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	str := strconv.FormatFloat(float64(f), 'g', -1, 32)
	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f64
}

func (c float64Converter) FromFloat64(v Variant) float64 {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return f
}

func (c float64Converter) FromTime(v Variant) float64 {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return 0
	}
	return float64(t.UnixNano())
}

func newFloat64Converter() IConvertStrategy[float64] {
	c := &float64Converter{}
	c.m = map[Kind]func(v Variant) float64{
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
