package studentmethods

import (
	"fmt"
	"net/http"

	"strconv"

	studentlogic "server/buisness/student"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Message string `json:"message"`
}

// Route 1 Get Student as per id

func GetStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, studentlogic.Students)
}

// Route 2 Get Student as per id

func CreateStudent(c echo.Context) error {
	name := c.FormValue("name")
	age, _ := strconv.ParseInt(c.FormValue("age"), 10, 64)
	rollno, _ := strconv.ParseInt(c.FormValue("rollNo"), 10, 64)
	if len(name) != 0 {

		var newStudent = studentlogic.Student{Name: name, Age: age, Rollno: rollno, Index: studentlogic.Index}
		*studentlogic.IndexPointer++
		studentlogic.AddStudent(newStudent)
		return c.JSON(http.StatusOK, newStudent)
	}
	return c.JSON(http.StatusBadRequest, Error{"Invalid Format"})
}

// Route 3 Get Student as per id
func GetStudent(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	if len(Student.Name) != 0 {

		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusNotFound, Error{"Not Found"})

}

// Route 4 Get Student as per id

func UpdateStudent(c echo.Context) error {
	var name = c.FormValue("name")
	// var age = c.FormValue("name")
	// var rollNo = c.FormValue("name")
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	if len(Student.Name) != 0 {

		var index int64 = studentlogic.GiveIndex(studentlogic.Students, id)

		fmt.Println(Student)
		Student.Name = name
		studentlogic.Students[index] = Student
		// students = remove(students, id)
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusBadRequest, Error{"Not Allowed"})

}

// Route 5 Get Student as per id

func DeleteStudent(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(studentlogic.Students, id)
	fmt.Println(Student)
	if len(Student.Name) != 0 {

		var index int64 = studentlogic.GiveIndex(studentlogic.Students, id)
		fmt.Println(index)
		studentlogic.Students = studentlogic.Remove(studentlogic.Students, index)
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusBadRequest, Error{"Invalid Format"})

}
