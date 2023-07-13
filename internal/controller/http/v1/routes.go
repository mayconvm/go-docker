package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/<TEMPLATE>/internal/usercase"
)

func RegisterRoutes(httpserver *echo.Echo, usercase usercase.UserCase) {

	// documents
	// DocumentsRegister(httpserver, usercase.Documents())
	StudentsRegister(httpserver, usercase.Students())
}
