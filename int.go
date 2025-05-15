package variant

import (
	"encoding/binary"
	"math"
	"time"
)

var _ IConvertStrategy[int] = (*intConverter)(nil)

type intConverter struct {
	Converter[int]
}

func (c intConverter) FromString(v Variant) int {

	tail := []byte{}
	for index, ch := range v.Data {
		if ch == '.' {
			tail = v.Data[index+1:]
			if len(tail) == 0 {
				return 0
			}
			v.Data = v.Data[:index]
			break
		}
	}
	// check if tail is not a number
	for _, ch := range tail {
		if ch < '0' || ch > '9' {
			return 0
		}
	}

	start := 0
	sLen := len(v.Data)
	if sLen == 0 {
		return 0
	}
	if v.Data[0] == '-' || v.Data[0] == '+' {
		start = 1
		sLen -= 1
		if len(v.Data) < 2 {
			return 0
		}
	}

	n := 0

	if (intSize == 32 && sLen < 10) || (intSize == 64 && sLen < 19) {
		for _, ch := range v.Data[start:] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int(ch)
		}
	}

	sub := 0
	if (intSize == 32 && sLen == 10) || (intSize == 64 && sLen == 19) {
		for _, ch := range v.Data[start : len(v.Data)-1] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int(ch)
		}

		cutoff := maxInt / 10
		remainder := maxInt % 10
		ch := v.Data[len(v.Data)-1] - '0'
		if ch > 9 {
			return 0
		}
		digit := int(ch)
		if v.Data[0] == '-' && cutoff == n && remainder == digit-1 {
			sub = 1
			n = n*10 + remainder
		} else if (cutoff == n && remainder >= digit) || cutoff > n {
			n = n*10 + digit
		} else {
			return 0
		}
	}

	if v.Data[0] == '-' {
		n = -n - sub
	}
	return n
}

func (c intConverter) FromBool(v Variant) int {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return 0
	}
	return 1
}

func (c intConverter) FromInt(v Variant) int {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	return i
}

func (c intConverter) FromInt8(v Variant) int {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	return int(i)
}

func (c intConverter) FromInt16(v Variant) int {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return int(i)
}

func (c intConverter) FromInt32(v Variant) int {
	return int(int32(binary.BigEndian.Uint32(v.Data)))
}

func (c intConverter) FromInt64(v Variant) int {
	return int(int64(binary.BigEndian.Uint64(v.Data)))
}

func (c intConverter) FromUint(v Variant) int {
	var i uint
	if intSize == 32 {
		i = uint(binary.BigEndian.Uint32(v.Data))
		if i > 1<<31-1 {
			return 0
		}
	} else if intSize == 64 {
		i = uint(binary.BigEndian.Uint64(v.Data))
		if i > 1<<63-1 {
			return 0
		}
	}
	return int(i)
}

func (c intConverter) FromUint8(v Variant) int {
	return int(v.Data[0])
}

func (c intConverter) FromUint16(v Variant) int {
	return int(binary.BigEndian.Uint16(v.Data))
}

func (c intConverter) FromUint32(v Variant) int {
	return int(binary.BigEndian.Uint32(v.Data))
}

func (c intConverter) FromUint64(v Variant) int {
	return int(binary.BigEndian.Uint64(v.Data))
}

func (c intConverter) FromFloat32(v Variant) int {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return int(f)
}

func (c intConverter) FromFloat64(v Variant) int {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return int(f)
}

func (c intConverter) FromTime(v Variant) int {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return 0
	}
	return int(t.UnixNano())
}

func newIntConverter() IConvertStrategy[int] {
	c := &intConverter{}
	c.m = map[Kind]func(v Variant) int{
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
