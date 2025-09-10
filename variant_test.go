package variant

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"
	"time"
	// "gopkg.in/yaml.v3"
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
	Age    uint64
	Weight float32
	IsMale any
}

type Data2 struct {
	Name   Variant
	Age    Variant
	Weight Variant
	IsMale Variant `json:"is_male"`
	Data   map[string]Variant
}

func Test_Marshal(t *testing.T) {
	data := map[string]interface{}{
		"Name":   New("John"),
		"Age":    New(1<<53 + 1),
		"Weight": New(1.2),
		"IsMale": New(true),
		"Now":    New(time.Now()),
		"data": map[string]any{
			"key1": New("value1"),
			"key2": New(10086),
		},
	}
	bytes, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func Test_Unmarshal(t *testing.T) {
	data := []byte(`{"name": "li\nli", "age": 36, "weight": 60.123, "is_male": false` +
		`, "now": "2024-06-20T14:00:00Z", "data": {"key1": "value1", "key2": []}}`)

	d := Data2{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Name.ToString())
	fmt.Println(d.Age.ToInt())
	fmt.Println(d.Weight.ToInt())
	fmt.Println(d.IsMale.ToBool())
	fmt.Println(d.Data)
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
		fmt.Println(d.Age.ToInt())
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
		fmt.Println(strconv.Itoa(int(d.Age)))
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

type Data3 struct {
	Name   string
	Age    Variant
	Weight float32
	IsMale any
}

/*
func Test_YmlMarshal(t *testing.T) {
	data := map[string]interface{}{
		"Name":    New("John\nDoe"),
		"Age":     New(uint64(42)),
		"Weight":  New(1.2),
		"IsMale":  New(true),
		"Now":     New(time.Now()),
		"Address": New(nil),
		"data1": map[string]any{
			"key1": New("value1"),
			"key2": New(10086),
		},
		"data2": []Variant{
			New("item1"),
			New(10086),
		},
	}
	bytes, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
*/

type UserProfile struct {
	Name     Variant `yaml:"name"`
	Age      Variant `yaml:"age"`
	Active   Variant `yaml:"active"`
	Nickname Variant `yaml:"nickname"`
	Scores   struct {
		Math    Variant `yaml:"math"`
		English Variant `yaml:"english"`
	} `yaml:"scores"`
	Settings struct {
		Theme         Variant `yaml:"theme"`
		Notifications Variant `yaml:"notifications"`
		Languages     []Variant
	} `yaml:"settings"`
	Items []struct {
		ID    Variant `yaml:"id"`
		Name  Variant `yaml:"name"`
		Price Variant `yaml:"price"`
	} `yaml:"items"`
	Metadata Variant `yaml:"metadata"`
}

/*
func TestVariant_UnmarshalYAML(t *testing.T) {
	yamlData := `
name: "Alice"
age: 30
active: true
nickname: null
scores:
  math: 95.5
  english: 88
settings:
  theme: "dark"
  notifications: false
  languages:
    - "en"
    - "fr"
items:
  - id: 1
    name: "Item1"
    price: 12.5
  - id: 2
    name: "Item2"
    price: 20.0
metadata: null
`
	var profile UserProfile
	if err := yaml.Unmarshal([]byte(yamlData), &profile); err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}
	if got := profile.Name.ToString(); got != "Alice" {
		t.Errorf("expected Name=Alice, got %q", got)
	}
	if got := profile.Age.ToInt64(); got != 30 {
		t.Errorf("expected Age=30, got %d", got)
	}
	if got := profile.Active.ToBool(); !got {
		t.Errorf("expected Active=true, got %v", got)
	}
	if profile.Nickname.Type != Invalid {
		t.Errorf("expected Nickname=Invalid, got %+v", profile.Nickname)
	}
	if got := profile.Scores.Math.ToFloat64(); math.Abs(got-95.5) != 0 {
		t.Errorf("expected Scores.Math=95.5, got %f", got)
	}
	if got := profile.Scores.English.ToInt(); got != 88 {
		t.Errorf("expected Scores.English=88, got %d", got)
	}
	if got := profile.Settings.Theme.ToString(); got != "dark" {
		t.Errorf("expected Settings.Theme=dark, got %q", got)
	}
	if got := profile.Settings.Notifications.ToBool(); got {
		t.Errorf("expected Settings.Notifications=false, got %v", got)
	}
	if !reflect.DeepEqual(profile.Settings.Languages, []Variant{New("en"), New("fr")}) {
		t.Errorf("expected Settings.Languages=[en fr], got %+v", profile.Settings.Languages)
	}
	if profile.Metadata.Type != Invalid {
		t.Errorf("expected Metadata=Invalid, got %+v", profile.Metadata)
	}
}
*/
