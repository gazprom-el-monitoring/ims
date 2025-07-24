package http

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/gazprom-el-monitoring/ims/internal/models"
	"github.com/gazprom-el-monitoring/ims/pkg/auth"
)

func AuthMiddleware(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	claims, err := auth.ParseToken(token)

	if err != nil || claims == nil {
		ctx.Output.SetStatus(401)
		ctx.Output.Body([]byte("Unauthorized"))
		return
	}

	role, err := models.TryParseRole((*claims)["role"].(string))

	if err != nil {
		ctx.Output.SetStatus(401)
		ctx.Output.Body([]byte("Unauthorized"))
		return
	}

	ctx.Input.SetData("role", role)
}
