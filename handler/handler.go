package handler

import (
	"github.com/labstack/echo/v4"
)

var studenthandler = studentHandler{"Student"}

func Api(e *echo.Echo) {
	// fmt.Println(studenthandler,"Hi")
	studentRouteGroup := e.Group("/students")

	studentRouteGroup.GET("/", studenthandler.getStudents)
	studentRouteGroup.GET("/:id", studenthandler.getStudent)
	studentRouteGroup.POST("/", studenthandler.createStudent)
	studentRouteGroup.PUT("/:id", studenthandler.updateStudent)
	studentRouteGroup.DELETE("/:id", studenthandler.deleteStudent)
}
