package models

import "github.com/oklog/ulid"

type DeviceModel struct {
	Id          ulid.ULID
	Class       ulid.ULID
	Name        string
	Description string
}
