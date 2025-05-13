package variant

import (
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
