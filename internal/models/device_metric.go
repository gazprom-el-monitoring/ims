package models

import "github.com/oklog/ulid"

type DeviceMetricType string

const (
	i8  = "i8"
	i16 = "i16"
	i32 = "i32"
	i64 = "i64"
	u8  = "u8"
	u16 = "u16"
	u32 = "u32"
	u64 = "u64"
)

type DeviceMetricClass string

type DeviceMetric struct {
	Id       ulid.ULID
	ModelId  ulid.ULID
	Scale    float64
	Type     DeviceMetricType
	Class    DeviceMetricClass
	Register string
}
