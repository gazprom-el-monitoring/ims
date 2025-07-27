package models

import (
	"net/netip"

	"github.com/google/uuid"
)

type GatewayProtocol string

type Gateway struct {
	Id       uuid.UUID
	Protocol GatewayProtocol
	IP       netip.Addr
	Port     uint16
}
