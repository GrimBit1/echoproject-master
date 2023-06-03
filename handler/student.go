package handler

import (
	"fmt"
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

	return c.JSON(http.StatusOK, studentlogic.GiveStudents())
}

// Route 2 Get Student as per id

func (s studentHandler) createStudent(c echo.Context) error {

	// Take all the inputs
	name := c.FormValue("name")
	age, err1 := strconv.ParseInt(c.FormValue("age"), 10, 64)
	rollno, err2 := strconv.ParseInt(c.FormValue("rollNo"), 10, 64)

	// Check error and give string
	err := checkErr(err1, err2)

	// if got error then return error
	if err == "Got Error" {
		return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Bad Format"})
	}

	// if no error then add student
	if len(name) != 0 {

		var newStudent = studentlogic.Student{Name: name, Age: age, Rollno: rollno}
		studentlogic.AddStudent(newStudent)
		return c.JSON(http.StatusOK, newStudent)
	}
	return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Invalid Format"})
}

// Route 3 Get Student as per id
func (s studentHandler) getStudent(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(id)

	// If student name is empty the give 404 Not Found
	if len(Student.Name) != 0 {
		// If not then send student as json
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusNotFound, studentlogic.Error{Message: "Not Found"})

}

// Route 4 Get Student as per id

func (s studentHandler) updateStudent(c echo.Context) error {
	// Already intializing the variables
	var name = c.FormValue("name")
	var stringage = c.FormValue("age")
	var stringrollno = c.FormValue("rollno")
	var age int64
	var rollno int64
	var err1 error
	var err2 error

	// checking if the user has given the age or rollno / if empty then use default values
	if len(stringage) != 0 {
		age, err1 = strconv.ParseInt(stringage, 10, 64)
	}

	if len(stringrollno) != 0 {
		rollno, err2 = strconv.ParseInt(stringrollno, 10, 64)
	}

	var id, err3 = strconv.ParseInt(c.Param("id"), 10, 64)

	// Check error and give string
	err := checkErr(err1, err2, err3)

	// if got error then return error
	if err == "Got Error" {
		return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Bad Format"})
	}

	fmt.Println(name, age, rollno, id)

	// Check for the Student if it is available then
	var oldstudent = studentlogic.Filter(id)

	if len(oldstudent.Name) == 0 {
		return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Bad Request"})
	}

	// if no errors then update student
	var updatedStudent = studentlogic.Student{Name: name, Age: age, Rollno: rollno} // create a template of the Student struct with values

	var Student = studentlogic.UpdateStudent(oldstudent, updatedStudent, id) // send to updateStudent function

	return c.JSON(http.StatusOK, Student)

}

// Route 5 Get Student as per id

func (s studentHandler) deleteStudent(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var Student = studentlogic.Filter(id)
	if len(Student.Name) != 0 {

		var index = id
		studentlogic.Remove(index)
		return c.JSON(http.StatusOK, Student)
	}
	return c.JSON(http.StatusBadRequest, studentlogic.Error{Message: "Bad Request"})

}

func checkErr(values ...error) string {
	for _, v := range values {
		if v != nil {
			return "Got Error"
		}
	}
	return "Got None"
}
