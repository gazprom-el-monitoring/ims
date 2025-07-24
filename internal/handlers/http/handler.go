package http

import "github.com/beego/beego/v2/server/web"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() {
	web.Get("/ping", h.Ping)
}
