package models

import (
	"net/netip"

	"github.com/oklog/ulid"
)

type GatewayProtocol string

type Gateway struct {
	Id       ulid.ULID
	Protocol GatewayProtocol
	IP       netip.Addr
	Port     uint16
}
