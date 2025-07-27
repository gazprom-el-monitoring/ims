package models

import "github.com/google/uuid"

type DeviceModel struct {
	Id              uuid.UUID
	ClassId         uuid.UUID
	Name            string
	Description     string
	DocumentationId uuid.UUID
}
