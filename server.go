package main

import (
	studentlogic "server/buisness/student"
	"server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// var students = []string{}
	// studentmethods.Greet()
	e := echo.New()

	e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// var c  echo.Context
	e.GET("/", func(c echo.Context) error {
		return c.String(int(200), "jejejehehe")
	})
	studentlogic.ConnectDB()
	handler.Api(e)
	e.Logger.Fatal(e.Start(":8000"))
}
