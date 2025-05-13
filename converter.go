package variant

import "time"

type IStrategy[T any] interface {
	Get(k Kind) func(v Variant) T
}

type IConvertStrategy[T any] interface {
	IStrategy[T]
	FromString(v Variant) T
	FromBool(v Variant) T
	FromInt(v Variant) T
	FromInt8(v Variant) T
	FromInt16(v Variant) T
	FromInt32(v Variant) T
	FromInt64(v Variant) T
	FromUint(v Variant) T
	FromUint8(v Variant) T
	FromUint16(v Variant) T
	FromUint32(v Variant) T
	FromUint64(v Variant) T
	FromFloat32(v Variant) T
	FromFloat64(v Variant) T
	FromTime(v Variant) T
	// add more methods below for other types as needed
}

var Strategies = strategies{
	string:  newStringConverter(),
	int:     newIntConverter(),
	int64:   newInt64Converter(),
	uint:    newUintConverter(),
	uint64:  newUint64Converter(),
	float32: newFloat32Converter(),
	float64: newFloat64Converter(),
	time:    newTimeConverter(),
	// add more strategies for other types as needed
}

type strategies struct {
	string  IConvertStrategy[string]
	int     IConvertStrategy[int]
	int64   IConvertStrategy[int64]
	uint    IConvertStrategy[uint]
	uint64  IConvertStrategy[uint64]
	float32 IConvertStrategy[float32]
	float64 IConvertStrategy[float64]
	time    IConvertStrategy[time.Time]
	// add more strategies for other types as needed
}

type Converter[T any] struct {
	m map[Kind]func(v Variant) T
}

func (c Converter[T]) Get(k Kind) func(v Variant) T {
	return c.m[k]
}
