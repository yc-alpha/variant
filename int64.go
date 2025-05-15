package variant

import (
	"encoding/binary"
	"math"
	"time"
)

var _ IConvertStrategy[int64] = (*int64Converter)(nil)

type int64Converter struct {
	Converter[int64]
}

func (c int64Converter) FromString(v Variant) int64 {
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

	var n int64 = 0

	if (intSize == 32 && sLen < 10) || (intSize == 64 && sLen < 19) {
		for _, ch := range v.Data[start:] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int64(ch)
		}
	}

	var sub int64 = 0
	if (intSize == 32 && sLen == 10) || (intSize == 64 && sLen == 19) {
		for _, ch := range v.Data[start : len(v.Data)-1] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + int64(ch)
		}

		var cutoff int64 = maxInt / 10
		var remainder int64 = maxInt % 10
		ch := v.Data[len(v.Data)-1] - '0'
		if ch > 9 {
			return 0
		}
		digit := int64(ch)
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

func (c int64Converter) FromBool(v Variant) int64 {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return 0
	}
	return 1
}

func (c int64Converter) FromInt(v Variant) int64 {
	var i int64
	if len(v.Data) == 4 && intSize == 32 {
		i = int64(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int64(binary.BigEndian.Uint64(v.Data))
	}
	return i
}

func (c int64Converter) FromInt8(v Variant) int64 {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	return int64(i)
}

func (c int64Converter) FromInt16(v Variant) int64 {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return int64(i)
}

func (c int64Converter) FromInt32(v Variant) int64 {
	return int64(binary.BigEndian.Uint32(v.Data))
}

func (c int64Converter) FromInt64(v Variant) int64 {
	return int64(binary.BigEndian.Uint64(v.Data))
}

func (c int64Converter) FromUint(v Variant) int64 {
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
	return int64(i)
}

func (c int64Converter) FromUint8(v Variant) int64 {
	return int64(v.Data[0])
}

func (c int64Converter) FromUint16(v Variant) int64 {
	return int64(binary.BigEndian.Uint16(v.Data))
}

func (c int64Converter) FromUint32(v Variant) int64 {
	return int64(binary.BigEndian.Uint32(v.Data))
}

func (c int64Converter) FromUint64(v Variant) int64 {
	return int64(binary.BigEndian.Uint64(v.Data))
}

func (c int64Converter) FromFloat32(v Variant) int64 {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return int64(f)
}

func (c int64Converter) FromFloat64(v Variant) int64 {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return int64(f)
}

func (c int64Converter) FromTime(v Variant) int64 {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return 0
	}
	return int64(t.UnixNano())
}

func newInt64Converter() IConvertStrategy[int64] {
	c := &int64Converter{}
	c.m = map[Kind]func(v Variant) int64{
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
