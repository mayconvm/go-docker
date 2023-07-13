package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	Entity "github.com/<TEMPLATE>/internal/entity/person"
	InUserCase "github.com/<TEMPLATE>/internal/usercase/person"
)

type studentsRoutes struct {
	StudentsUserCase *InUserCase.StudentUserCase
}

func StudentsRegister(httpserver *echo.Echo, duc *InUserCase.StudentUserCase) {
	dr := &studentsRoutes{duc}

	g := httpserver.Group("/students")
	g.GET("/", dr.get)
	g.POST("/", dr.post)
}

func (dr *studentsRoutes) get(c echo.Context) error {
	var entity Entity.StudentEntity
	if err := c.Bind(&entity); err != nil {
		return err
	}

	//TODO: check if there is paramets to search

	result := dr.StudentsUserCase.SearchDocuments(entity)
	return c.JSON(http.StatusOK, result)
}

func (dr *studentsRoutes) post(c echo.Context) error {
	entity := new(Entity.StudentEntity)
	if err := c.Bind(entity); err != nil {
		return err
	}

	//TODO: check if entity is empty

	result := dr.StudentsUserCase.SaveDocument(entity)

	return c.JSON(http.StatusCreated, result)
}
