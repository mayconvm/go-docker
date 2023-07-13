package http

import (
	"github.com/labstack/echo/v4"
	v1 "github.com/<TEMPLATE>/internal/controller/http/v1"
	"github.com/<TEMPLATE>/internal/usercase"
)

func RegisterRoutes(httpserver *echo.Echo, usercase usercase.UserCase) {
	// TODO: register routes by version

	v1.RegisterRoutes(httpserver, usercase)
}
