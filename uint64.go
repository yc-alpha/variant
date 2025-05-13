package variant

import (
	"encoding/binary"
	"math"
	"time"
)

var _ IConvertStrategy[uint64] = (*uint64Converter)(nil)

type uint64Converter struct {
	Converter[uint64]
}

func (u uint64Converter) FromString(v Variant) uint64 {
	for index, ch := range v.Data {
		if ch == '.' {
			v.Data = v.Data[:index]
			break
		}
	}
	start := 0
	sLen := len(v.Data)
	if sLen == 0 || v.Data[0] == '-' {
		return 0
	}
	if v.Data[0] == '+' {
		start = 1
		sLen -= 1
		if len(v.Data) < 2 {
			return 0
		}
	}

	var n uint64 = 0

	if (intSize == 32 && sLen < 10) || (intSize == 64 && sLen < 20) {
		for _, ch := range v.Data[start:] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + uint64(ch)
		}
	}

	if (intSize == 32 && sLen == 10) || (intSize == 64 && sLen == 20) {
		for _, ch := range v.Data[start : len(v.Data)-1] {
			ch -= '0'
			if ch > 9 {
				return 0
			}
			n = n*10 + uint64(ch)
		}

		var cutoff uint64 = maxUint / 10
		var remainder uint64 = maxUint % 10
		ch := v.Data[len(v.Data)-1] - '0'
		if ch > 9 {
			return 0
		}
		digit := uint64(ch)
		if (cutoff == n && remainder >= digit) || cutoff > n {
			n = n*10 + digit
		} else {
			return 0
		}
	}
	return n
}

func (u uint64Converter) FromBool(v Variant) uint64 {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return 0
	}
	return 1
}

func (u uint64Converter) FromInt(v Variant) uint64 {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	if i <= 0 {
		return 0
	}
	return uint64(i)
}

func (u uint64Converter) FromInt8(v Variant) uint64 {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	if i <= 0 {
		return 0
	}
	return uint64(i)
}

func (u uint64Converter) FromInt16(v Variant) uint64 {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	if i <= 0 {
		return 0
	}
	return uint64(i)
}

func (u uint64Converter) FromInt32(v Variant) uint64 {
	i := int32(binary.BigEndian.Uint32(v.Data))
	if i <= 0 {
		return 0
	}
	return uint64(i)
}

func (u uint64Converter) FromInt64(v Variant) uint64 {
	i := binary.BigEndian.Uint64(v.Data)
	if int(i) <= 0 {
		return 0
	}
	return i
}

func (u uint64Converter) FromUint(v Variant) uint64 {
	var i uint64
	if intSize == 32 {
		i = uint64(binary.BigEndian.Uint32(v.Data))
	} else if intSize == 64 {
		i = binary.BigEndian.Uint64(v.Data)
	}
	return i
}

func (u uint64Converter) FromUint8(v Variant) uint64 {
	return uint64(v.Data[0])
}

func (u uint64Converter) FromUint16(v Variant) uint64 {
	return uint64(binary.BigEndian.Uint16(v.Data))
}

func (u uint64Converter) FromUint32(v Variant) uint64 {
	return uint64(binary.BigEndian.Uint32(v.Data))
}

func (u uint64Converter) FromUint64(v Variant) uint64 {
	return binary.BigEndian.Uint64(v.Data)
}

func (u uint64Converter) FromFloat32(v Variant) uint64 {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return uint64(f)
}

func (u uint64Converter) FromFloat64(v Variant) uint64 {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return uint64(f)
}

func (u uint64Converter) FromTime(v Variant) uint64 {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return 0
	}
	return uint64(t.UnixNano())
}

func newUint64Converter() IConvertStrategy[uint64] {
	c := &uint64Converter{}
	c.m = map[Kind]func(v Variant) uint64{
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
