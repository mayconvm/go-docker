package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/<TEMPLATE>/config"
	"github.com/<TEMPLATE>/pkg/logger"
)

func New() *echo.Echo {
	e := echo.New()
	// e.Validator = &CustomValidator{validator: validator.New()}
	// transaction.Echo = e

	// TODO: refactor
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// RegisterRoutes(transaction)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}

func Start(e *echo.Echo, cfg *config.Config, log *logger.Logger) {
	defer e.Close()

	log.Info("Start http server :" + cfg.HTTP.Port)
	e.Logger.Fatal(e.Start(":" + cfg.HTTP.Port))
}
