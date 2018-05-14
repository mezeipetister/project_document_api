package main

import (
	"Projects/project_document_api/src/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.POST("/login", user.Login)

	// Create
	e.GET("/create", user.Create)

	// Mongo
	e.GET("/mongo", user.Mongo)

	var counter int

	e.GET("/counter", func(c echo.Context) error {
		counter++
		return c.JSON(http.StatusOK, map[string]string{
			"counter": fmt.Sprintf("%v", counter),
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
