package models

import "github.com/oklog/ulid"

type DeviceClass struct {
	Id       ulid.ULID
	Bus_name string
}
