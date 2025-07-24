package models

import "github.com/oklog/ulid"

type DeviceProtocol string

type Device struct {
	Id        ulid.ULID
	ModelId   ulid.ULID
	Location  string
	Protocol  *DeviceProtocol
	GatewayId *ulid.ULID
	Address   *string
}
