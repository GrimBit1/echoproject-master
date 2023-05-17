package handler

import (
	"net/http"

	"strconv"

	studentlogic "server/buisness/student"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	Name string
}

var httpImaTeapot int = 418

// Route 1 Get Student as per id

func (s studentHandler) getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, studentlogic.Students)
}

// Route 2 Get Student as per id

func (s studentHandler) createStudent(c echo.Context) error {
	name := c.FormValue("name")
	age, _ := strconv.ParseInt(c.FormValue("age"), 10, 64)
	rollno, _ := strconv.ParseInt(c.FormValue("rollNo"), 10, 64)
	if len(name) != 0 {

		var newStudent = studentlogic.Student{Name: name, Age: age, Rollno: rollno, Index: studentlogic.Index}
		*studentlogic.IndexPointer++
		studentlogic.AddStudent(newStudent)
		return c.JSON(http.StatusOK, newStudent)
	}
	return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Invalid Format"})
}

// Route 3 Get Student as per id
func (s studentHandler) getStudent(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	if len(Student.Name) != 0 {

		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusNotFound, studentlogic.Error{Message: "Not Found"})

}

// Route 4 Get Student as per id

func (s studentHandler) updateStudent(c echo.Context) error {
	var name = c.FormValue("name")
	if len(name) == 0 {
		return c.JSON(httpImaTeapot, studentlogic.Error{Message: "Name must be valid"})
	}
	// var age = c.FormValue("name")
	// var rollNo = c.FormValue("name")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	if len(Student.Name) != 0 {

		var index int64 = studentlogic.GiveIndex(studentlogic.Students, id)

		Student.Name = name
		studentlogic.Students[index] = Student
		// students = remove(students, id)
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Not Allowed"})

}

// Route 5 Get Student as per id

func (s studentHandler) deleteStudent(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	if len(Student.Name) != 0 {

		var index int64 = studentlogic.GiveIndex(studentlogic.Students, id)
		studentlogic.Students = studentlogic.Remove(studentlogic.Students, index)
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Bad Request"})

}
