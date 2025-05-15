package variant

import (
	"math"
	"testing"
	"time"
)

func TestVariant_ToUint64(t *testing.T) {
	tt := time.Now()
	targets := []Pair[uint64]{
		{true, 1},
		{false, 0},
		{"192hello.你好！", 0},
		{".234", 0},
		{"234.", 0},
		{"127.0.0.1", 0},
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
		{tt, uint64(tt.UnixNano())},
	}
	for _, pair := range targets {
		t.Run("ToUint64", func(t *testing.T) {
			v := New(pair.Key)
			assert(v.ToUint64() == pair.Val)
		})
	}
}
