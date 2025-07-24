package models

import "github.com/oklog/ulid"

type DeviceModel struct {
	Id              ulid.ULID
	ClassId         ulid.ULID
	Name            string
	Description     string
	DocumentationId ulid.ULID
}
