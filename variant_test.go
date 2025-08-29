package variant

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"
	"time"
)

type Pair[T any] struct {
	Key any
	Val T
}

func assert(condition bool, msg ...any) {
	if !condition {
		log.Fatal(msg...)
	}
}

func TestVariant_ToBool(t *testing.T) {
	tt := time.Now()
	targets := []Pair[bool]{
		{true, true},
		{false, false},
		{"Hello World!你好！", true},
		{math.MaxInt8, true},
		{int8(0), false},
		{int16(0), false},
		{int32(0), false},
		{int64(0), false},
		{uint8(0), false},
		{uint16(0), false},
		{uint32(0), false},
		{uint64(0), false},
		{float32(12345.123), true},
		{float32(0), false},
		{123456789.12345678, true},
		{tt, true},
	}
	for _, pair := range targets {
		t.Run("ToInt", func(t *testing.T) {
			v := New(pair.Key)
			assert(v.ToBool() == pair.Val)
		})
	}
}

func TestVariant_Empty(t *testing.T) {
	v := Variant{Type: Bool}
	assert(v.ToString() == "false", "")
	assert(!v.ToBool(), "")
	assert(v.ToInt() == 0, "")
	assert(v.ToUint() == 0, "")
	assert(v.ToFloat32() == 0, "")
	assert(v.ToFloat64() == 0, "")

	v = Variant{Type: String}
	assert(v.ToString() == "", "1")
	assert(!v.ToBool(), "2")
	assert(v.ToInt() == 0, "3")
	assert(v.ToUint() == 0, "4")
	assert(v.ToFloat32() == 0, "5")
	assert(v.ToFloat64() == 0, "6")

	v = Variant{Type: Int}
	assert(v.ToString() == "0", "")
	assert(!v.ToBool(), "")
	assert(v.ToInt() == 0, "")
	assert(v.ToUint() == 0, "")
	assert(v.ToFloat32() == 0, "")
	assert(v.ToFloat64() == 0, "")

	v = Variant{Type: Float32}
	assert(v.ToString() == "0", "")
	assert(!v.ToBool(), "")
	assert(v.ToInt() == 0, "")
	assert(v.ToUint() == 0, "")
	assert(v.ToFloat32() == 0, "")
	assert(v.ToFloat64() == 0, "")
	v = Variant{Type: Float64}
	assert(v.ToString() == "0", "")
	assert(!v.ToBool(), "")
	assert(v.ToInt() == 0, "")
	assert(v.ToUint() == 0, "")
	assert(v.ToFloat32() == 0, "")
	assert(v.ToFloat64() == 0, "")

	v = Variant{Type: Time}
	assert(v.ToString() == "", "")
	assert(!v.ToBool(), "")
	assert(v.ToInt() == 0, "")
	assert(v.ToUint() == 0, "")
	assert(v.ToFloat32() == 0, "")
	assert(v.ToFloat64() == 0, "")
}

func Test_ToTime(t *testing.T) {
	v := New(true)
	assert(v.ToTime().Equal(time.Now()), "1")
	v = New(false)
	assert(v.ToTime().Equal(time.Time{}), "2")
	v = New(uint(0))
	assert(v.ToTime().Equal(time.Unix(0, 0)), "3")
	v = New(uint8(86))
	assert(v.ToTime().Equal(time.Unix(0, 86)), "4")
	v = New(10086)
	assert(v.ToTime().Equal(time.Unix(0, 10086)), "5")
	v = New(int8(86))
	assert(v.ToTime().Equal(time.Unix(0, 86)), "6")
	v = New(int16(0))
	assert(v.ToTime().Equal(time.Unix(0, 0)), "7")
	v = New(int32(86))
	assert(v.ToTime().Equal(time.Unix(0, 86)), "8")
	v = New(int64(86))
	assert(v.ToTime().Equal(time.Unix(0, 86)), "9")
	v = New("100.86")
	assert(v.ToTime().Equal(time.Time{}), "10")
	v = New("-100.86")
	assert(v.ToTime().Equal(time.Time{}), "11")
	v = New("1abc")
	assert(v.ToTime().Equal(time.Time{}), "12")
	v = New(float32(-86.2))
	assert(v.ToTime().Equal(time.Unix(0, -86)), "13")
	v = New(-86.1)
	assert(v.ToTime().Equal(time.Unix(0, -86)), "14")

	tt := time.Now()
	v = New(tt)
	assert(v.ToTime().UnixNano() == tt.UnixNano(), "14")
}

type Data struct {
	Name   string
	Age    int
	Weight float32
	IsMale any
}

type Data2 struct {
	Name   Variant
	Age    Variant
	Weight Variant
	IsMale Variant
}

func Test_Marshal(t *testing.T) {
	data := map[string]interface{}{
		"Name":    New("John"),
		"Age":     New(42),
		"Weight":  New(1.2),
		"IsMale":  New(true),
		"Now":     New(time.Now()),
		"Address": New(nil),
	}
	bytes, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

// BenchmarkVariant-10    	  110092	     10829 ns/op	     456 B/op	       8 allocs/op
func Benchmark_UnmarshalJson(b *testing.B) {
	data := []byte(`{"name": "lili", "age": 36, "weight": 60.123, "is_male": null}`)
	for i := 0; i < b.N; i++ {
		d := Data2{}
		err := json.Unmarshal(data, &d)
		if err != nil {
			panic(err)
		}
		fmt.Println(d.Name.ToString())
		fmt.Println(d.Age.ToString())
		fmt.Println(d.Weight.ToString())
		fmt.Println(d.IsMale.ToBool())
	}
	b.ReportAllocs()
}

// Benchmark_UnmarshalJson-10    	  119211	     10826 ns/op	     368 B/op	      11 allocs/op
func Benchmark_UnmarshalJson2(b *testing.B) {
	data := []byte(`{"name": "lili", "age": 36, "weight": 60.123, "is_male": null}`)
	for i := 0; i < b.N; i++ {
		d := Data{}
		err := json.Unmarshal(data, &d)
		if err != nil {
			panic(err)
		}
		fmt.Println(d.Name)
		fmt.Println(strconv.Itoa(d.Age))
		fmt.Println(strconv.FormatFloat(float64(d.Weight), 'f', -1, 64))
		if d.IsMale != nil {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
	b.ReportAllocs()
}

func Test_Equal(t *testing.T) {
	v := Nil
	assert(v.Equal(nil))

	v1 := New(123)
	assert(v1.Equal(123))
	assert(v1.Equal(v1))
	assert(v1.Equal(New(123)))

	assert(!v1.Equal(1234))
	assert(!v1.Equal(Nil))
	assert(!v1.Equal(New(1234)))
}
