package app

import (
	"github.com/<TEMPLATE>/config"
	"github.com/<TEMPLATE>/internal/controller/http"
	"github.com/<TEMPLATE>/internal/usercase"
	"github.com/<TEMPLATE>/pkg/httpserver"
	"github.com/<TEMPLATE>/pkg/logger"
	"github.com/<TEMPLATE>/pkg/mongo"
)

func Start(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	mc, _ := mongo.New(cfg)

	all_usercase := usercase.New(usercase.UserCase{
		Cfg:   cfg,
		Log:   l,
		Mongo: mc.Client,
	})

	http_server := httpserver.New()
	http.RegisterRoutes(http_server, all_usercase)

	//Start httpserver
	httpserver.Start(http_server, cfg, l)

	// Waiting signal
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// select {
	// case s := <-interrupt:
	// 	l.Info("app - Run - signal: " + s.String())
	// case err = <-httpServer.Notify():
	// 	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	// case err = <-rmqServer.Notify():
	// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	// }

}
