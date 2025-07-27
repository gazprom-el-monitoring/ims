package models

import "github.com/google/uuid"

type DeviceClass struct {
	Id      uuid.UUID
	BusName string
}
