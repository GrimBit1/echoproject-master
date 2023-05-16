package main

import (
	studentmethods "server/handler/student"

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

	studentRouteGroup := e.Group("/students")
	// var c  echo.Context
	studentRouteGroup.GET("/", studentmethods.GetStudents)
	studentRouteGroup.GET("/:id", studentmethods.GetStudent)
	studentRouteGroup.POST("/", studentmethods.CreateStudent)
	studentRouteGroup.PUT("/:id", studentmethods.UpdateStudent)
	studentRouteGroup.DELETE("/:id", studentmethods.DeleteStudent)
	e.Logger.Fatal(e.Start(":8000"))
}
