package api_type

import (
	"database/sql/driver"
	"fmt"
)

type ApiType int

const (
	Unknown  ApiType = 0
	Internal ApiType = 1
	Public   ApiType = 2
	Swagger  ApiType = 3
)

type NullApiType struct {
	Enum  ApiType
	Valid bool
}

var enumApiTypeName = map[int]string{
	0: "unknown",
	1: "internal",
	2: "public",
	3: "swagger",
}

var enumApiTypeValue = map[string]int{
	"unknown":  0,
	"internal": 1,
	"public":   2,
	"swagger":  3,
}

// GENERATE
func ParseApiType(s string) (ApiType, bool) {
	val, ok := enumApiTypeValue[s]
	return ApiType(val), ok
}

func (e ApiType) Enum() int {
	return int(e)
}

func (e ApiType) Name() string {
	return enumApiTypeName[e.Enum()]
}

func (e ApiType) Value() (driver.Value, error) {
	if e == 0 {
		return nil, nil
	}
	return e.String(), nil
}

func (e ApiType) String() string {
	s, ok := enumApiTypeName[e.Enum()]
	if ok {
		return s
	}
	return fmt.Sprintf("ApiType(%v)", e.Enum())
}

func WrapApiType(enum ApiType) NullApiType {
	return NullApiType{Enum: enum, Valid: true}
}

func (n NullApiType) Apply(s ApiType) ApiType {
	if n.Valid {
		return n.Enum
	}
	return s
}

func (n NullApiType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Enum.Value()
}
