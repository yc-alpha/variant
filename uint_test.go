package variant

import (
	"math"
	"testing"
	"time"
)

func TestVariant_ToUint(t *testing.T) {
	tt := time.Now()
	targets := []Pair[uint]{
		{true, 1},
		{false, 0},
		{"Hello World!你好！", 0},
		{math.MaxInt8, 127},
		{math.MinInt8, 0},
		{math.MaxInt16, 32767},
		{math.MinInt16, 0},
		{math.MaxInt32, 2147483647},
		{math.MinInt32, 0},
		{math.MaxInt64, 9223372036854775807},
		{math.MinInt64, 0},
		{math.MaxUint8, 255},
		{math.MaxUint16, 65535},
		{math.MaxUint32, 4294967295},
		{float32(12345.123), 12345},
		{float32(-12345.123), 0},
		{123456789.12345678, 123456789},
		{tt, uint(tt.UnixNano())},
	}
	for _, pair := range targets {
		t.Run("ToInt", func(t *testing.T) {
			v := New(pair.Key)
			assert(v.ToUint() == pair.Val)
		})
	}
}
