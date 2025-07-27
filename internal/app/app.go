package app

import (
	"context"
	"fmt"
	"github.com/gazprom-el-monitoring/ims/internal/config"
	"github.com/gazprom-el-monitoring/ims/internal/handlers/http"
	"github.com/gazprom-el-monitoring/ims/internal/repositories"
	"github.com/gazprom-el-monitoring/ims/internal/services"
	"github.com/gazprom-el-monitoring/ims/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"log/slog"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	l := logger.NewConsole()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	conf, err := config.Init()
	if err != nil {
		l.Fatal().Err(err).Msg("cannot load config")
	}

	pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Name))
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pgPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		l.Fatal().Err(err).Msg("Unable to connect to db")

	}
	defer pgPool.Close()

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &l,
	}))

	r := repositories.NewRepositories(pgPool)
	s := services.NewServices(r)
	h := http.NewHandler(s)
	h.Init(app)

	go func() {
		err := app.Listen(":8000")
		if err != nil {
			l.Error().Err(err).Msg("Can not start http server")
			cancel()
		}
	}()

	<-ctx.Done()
	slog.Error("Shutdown")

	// Создаем контекст с таймаутом для завершения
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		l.Fatal().Err(err).Msg("Unable to shutdown")
	}
}
