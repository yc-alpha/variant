package variant

import (
	"testing"
	"time"
)

func Test_ToTime2(t *testing.T) {
	targets := []Pair[time.Time]{
		{false, time.Time{}},
		{"100.86", time.Time{}},
		{int8(0), time.Unix(0, 0)},
		{int16(0), time.Unix(0, 0)},
		{int32(10086), time.Unix(0, 10086)},
		{int64(86), time.Unix(0, 86)},
		{uint8(0), time.Unix(0, 0)},
		{uint16(0), time.Unix(0, 0)},
		{uint32(10086), time.Unix(0, 10086)},
		{uint64(86), time.Unix(0, 86)},
		{float32(-12345.123), time.Unix(0, -12345)},
		{float32(0), time.Unix(0, 0)},
		{123456789.12345678, time.Unix(0, 123456789)},
	}
	for _, pair := range targets {
		t.Run("ToTime", func(t *testing.T) {
			v := New(pair.Key)
			assert(v.ToTime() == pair.Val)
		})
	}
}
