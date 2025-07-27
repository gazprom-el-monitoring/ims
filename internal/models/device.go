package models

import (
	"github.com/google/uuid"
)

type DeviceProtocol string

type Device struct {
	Id        uuid.UUID
	ModelId   uuid.UUID
	Location  string
	Protocol  *DeviceProtocol
	GatewayId *uuid.UUID
	Address   *string
}
