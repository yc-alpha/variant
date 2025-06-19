package variant

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestVariant_ToString(t *testing.T) {
	// time
	tt := time.Now()
	targets := []Pair[string]{
		{true, "true"},
		{false, "false"},
		{"Hello World!你好！", "Hello World!你好！"},
		{math.MaxInt8, "127"},
		{math.MinInt8, "-128"},
		{math.MaxInt16, "32767"},
		{math.MinInt16, "-32768"},
		{math.MaxInt32, "2147483647"},
		{math.MinInt32, "-2147483648"},
		{math.MaxInt64, "9223372036854775807"},
		{math.MinInt64, "-9223372036854775808"},
		{math.MaxUint8, "255"},
		{math.MaxUint16, "65535"},
		{math.MaxUint32, "4294967295"},
		{float32(12345.123), "12345.123"},
		{123456789.12345678, "123456789.12345678"},
		{map[string]int{}, ""},
		{tt, tt.Format(time.DateTime)},
	}
	for _, pair := range targets {
		t.Run("ToString", func(t *testing.T) {
			v := New(pair.Key)
			assert(v.ToString() == pair.Val)
		})
	}
}

func Benchmark_ToString(b *testing.B) {
	var Nil *string
	str := "Hello World!你好！"
	boolean := true
	var i8 int8 = math.MaxInt8
	var i16 int16 = math.MaxInt16
	var i32 int32 = math.MaxInt32
	var i64 int64 = math.MaxInt64
	var ui8 uint8 = math.MaxUint8
	var ui16 uint16 = math.MaxUint16
	var ui32 uint32 = math.MaxUint32
	var ui64 uint64 = math.MaxUint64
	var f32 float32 = 12345.123
	var f64 float64 = 123456789.12345678
	tt := time.Now()
	for i := 0; i < b.N; i++ {
		fmt.Println(New(Nil).ToString())
		fmt.Println(New(&str).ToString())
		fmt.Println(New(&boolean).ToString())
		fmt.Println(New(&i8).ToString())
		fmt.Println(New(&i16).ToString())
		fmt.Println(New(&i32).ToString())
		fmt.Println(New(&i64).ToString())
		fmt.Println(New(&ui8).ToString())
		fmt.Println(New(&ui16).ToString())
		fmt.Println(New(&ui32).ToString())
		fmt.Println(New(&ui64).ToString())
		fmt.Println(New(&f32).ToString())
		fmt.Println(New(&f64).ToString())
		fmt.Println(New(&tt).ToString())
	}
	b.ReportAllocs()
}
