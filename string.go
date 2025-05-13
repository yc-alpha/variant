package variant

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"
	"unsafe"
)

var _ IConvertStrategy[string] = (*stringConverter)(nil)

type stringConverter struct {
	Converter[string]
}

func (c stringConverter) FromString(v Variant) string {
	return *(*string)(unsafe.Pointer(&v.Data))
}

func (c stringConverter) FromBool(v Variant) string {
	if len(v.Data) == 0 || v.Data[0] == 0x00 {
		return "false"
	}
	return "true"
}

func (c stringConverter) FromInt(v Variant) string {
	var i int
	if len(v.Data) == 4 && intSize == 32 {
		i = int(binary.BigEndian.Uint32(v.Data))
	} else if len(v.Data) == 8 && intSize == 64 {
		i = int(binary.BigEndian.Uint64(v.Data))
	}
	return strconv.Itoa(i)
}

func (c stringConverter) FromInt8(v Variant) string {
	var i int8
	if len(v.Data) == 1 {
		i = int8(v.Data[0])
	}
	b := strconv.AppendInt(make([]byte, 0), int64(i), 10)
	return *(*string)(unsafe.Pointer(&b))
}

func (c stringConverter) FromInt16(v Variant) string {
	var i int16
	if len(v.Data) == 2 {
		i = int16(binary.BigEndian.Uint16(v.Data))
	}
	return strconv.Itoa(int(i))
}

func (c stringConverter) FromInt32(v Variant) string {
	i := int32(binary.BigEndian.Uint32(v.Data))
	return strconv.Itoa(int(i))
}

func (c stringConverter) FromInt64(v Variant) string {
	i := int64(binary.BigEndian.Uint64(v.Data))
	return strconv.Itoa(int(i))
}

func (c stringConverter) FromUint(v Variant) string {
	var i uint64
	if intSize == 32 {
		i = uint64(binary.BigEndian.Uint32(v.Data))
	} else if intSize == 64 {
		i = binary.BigEndian.Uint64(v.Data)
	}
	return strconv.FormatUint(i, 10)
}

func (c stringConverter) FromUint8(v Variant) string {
	i := v.Data[0]
	b := strconv.AppendUint(make([]byte, 0), uint64(i), 10)
	return *(*string)(unsafe.Pointer(&b))
}

func (c stringConverter) FromUint16(v Variant) string {
	i := uint64(binary.BigEndian.Uint16(v.Data))
	return strconv.FormatUint(i, 10)
}

func (c stringConverter) FromUint32(v Variant) string {
	i := uint64(binary.BigEndian.Uint32(v.Data))
	return strconv.FormatUint(i, 10)
}

func (c stringConverter) FromUint64(v Variant) string {
	i := binary.BigEndian.Uint64(v.Data)
	return strconv.FormatUint(i, 10)
}

func (c stringConverter) FromFloat32(v Variant) string {
	var f float32
	if len(v.Data) == 4 {
		f = math.Float32frombits(binary.BigEndian.Uint32(v.Data))
	}
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func (c stringConverter) FromFloat64(v Variant) string {
	var f float64
	if len(v.Data) == 8 {
		f = math.Float64frombits(binary.BigEndian.Uint64(v.Data))
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (c stringConverter) FromTime(v Variant) string {
	var t time.Time
	err := t.UnmarshalBinary(v.Data)
	if err != nil {
		return ""
	}
	return t.Format(v.layout)
}

func newStringConverter() IConvertStrategy[string] {
	c := &stringConverter{}
	c.m = map[Kind]func(v Variant) string{
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
