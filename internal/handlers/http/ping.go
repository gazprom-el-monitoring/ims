package http

import "github.com/beego/beego/v2/server/web/context"

func (h *Handler) Ping(ctx *context.Context) {
	k := map[string]interface{}{
		"ok": true,
	}
	ctx.JSONResp(k)
}
