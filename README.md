# Variant

Variant is a Go package that provides a flexible type conversion system for handling different data types in a unified way.

## Installation

```bash
go get github.com/yc-alpha/variant
```

## Features

- Support for multiple primitive types including bool, integers, floats, string and time
- Type-safe conversions between different data types
- JSON marshaling/unmarshalling support
- Thread-safe operations

## Supported Types

- Bool
- Int (int8, int16, int32, int64)
- Uint (uint8, uint16, uint32, uint64)
- Float (float32, float64)
- String
- Time

## Usage

```go
// Create a new variant
v := variant.New("123")

// Convert to different types
intVal := v.ToInt()      // 123
strVal := v.ToString()   // "123"
floatVal := v.ToFloat64() // 123.0

// Create time variant with custom layout
t := time.Now()
v = variant.New(t)
str := v.SetLayout("2006-01-02").ToString() // Date in specified format

// JSON Marshal
data := map[string]interface{}{
    "Name":    variant.New("John"),
    "Age":     variant.New(42),
    "Weight":  variant.New(1.2),
    "IsMale":  variant.New(true),
    "Now":     variant.New(time.Now()),
    "Address": variant.New(nil),
}
bytes, err := json.Marshal(&data)

// JSON Unmarshal
data := []byte(`{"name": "lili", "age": 36, "weight": 60.123, "is_male": null}`)
m := map[string]variant.Variant{}
err := json.Unmarshal(data, &m)
```

## License

Distributed under the [MIT license](./LICENSE).
