package bootstrap

import (
	"context"
	"errors"
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/api/mails"
	"github.com/agung96tm/miblog/api/middlewares"
	"github.com/agung96tm/miblog/api/policies"
	"github.com/agung96tm/miblog/api/repositories"
	"github.com/agung96tm/miblog/api/routes"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	lib.Module,
	routes.Module,
	controllers.Module,
	repositories.Module,
	services.Module,
	middlewares.Module,
	mails.Module,
	policies.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler lib.HttpHandler, routes routes.Routes, config lib.Config, middlewares middlewares.Middlewares, database lib.Database) {
	db, err := database.ORM.DB()
	if err != nil {
		log.Fatalf("[Database] Error to get database connection: %v", err)
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting Application")

			if err := db.Ping(); err != nil {
				log.Fatalf("[Database] Error to get database connection: %v", err)
			}

			go func() {
				middlewares.Setup()
				routes.Setup()

				if err := handler.Engine.Start(config.Http.ListenAddr()); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						log.Debug("Shutting down the Application")
					} else {
						log.Fatalf("Error to Start Application: %v", err)
					}
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping Application")

			_ = handler.Engine.Close()
			_ = db.Close()

			return nil
		},
	})
}
