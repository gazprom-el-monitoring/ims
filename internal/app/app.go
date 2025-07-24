package app

import (
	"context"
	"github.com/beego/beego/v2/server/web"
	"github.com/gazprom-el-monitoring/ims/internal/handlers/http"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	h := http.NewHandler()

	h.Init()

	go web.Run()

	<-ctx.Done()
	slog.Error("Shutdown")

	// Создаем контекст с таймаутом для завершения
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	err := web.BeeApp.Server.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("Error while graceful shutdown", "err", err)
	}
}
