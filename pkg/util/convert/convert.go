// Package convert convert string to number
package convert

import (
	"encoding/json"
	"strconv"
)

type StrTo string

// String string转string
func (s StrTo) String() string {
	return string(s)
}

// Int string转int
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

// MustInt string强制转为int
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

// Uint string转为uint
func (s StrTo) Uint() (uint, error) {
	v, err := strconv.Atoi(s.String())
	return uint(v), err
}

// MustUint string强制转为uint
func (s StrTo) MustUint() uint {
	v, _ := s.Uint()
	return v
}

// UInt32 string转为uint32
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// MustUInt32 string强制转为uint32
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

// UInt64 string转为uint64
func (s StrTo) UInt64() (uint64, error) {
	v, err := strconv.Atoi(s.String())
	return uint64(v), err
}

// MustUInt64 string强制转为uint64
func (s StrTo) MustUInt64() uint64 {
	v, _ := s.UInt64()
	return v
}

// Int64 string转为int64
func (s StrTo) Int64() (int64, error) {
	v, err := strconv.Atoi(s.String())
	return int64(v), err
}

// MustInt64 string强制转为int64
func (s StrTo) MustInt64() int64 {
	v, _ := s.Int64()
	return v
}

// ToStr 将interface转为string
func ToStr(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
