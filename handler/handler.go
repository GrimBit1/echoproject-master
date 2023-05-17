package handler

import "github.com/labstack/echo/v4"

var studenthandler studentHandler

func Api(e *echo.Echo) {
	studentRouteGroup := e.Group("/students")

	studentRouteGroup.GET("/", studenthandler.getStudents)
	studentRouteGroup.GET("/:id", studenthandler.getStudent)
	studentRouteGroup.POST("/", studenthandler.createStudent)
	studentRouteGroup.PUT("/:id", studenthandler.updateStudent)
	studentRouteGroup.DELETE("/:id", studenthandler.deleteStudent)
}
