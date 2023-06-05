package studentlogic

import (
	"context"
	"fmt"
	checkerror "server/checkError"
	"server/connectdb"
)

var ctx context.Context
var Students = []Student{}

func GiveStudents() []Student {
	var tempArray = []Student{}
	tempArr, err := connectdb.GiveDB().Query("Select * from users")
	// tempArr.Scan()
	checkerror.Checkerror(err)
	fmt.Println(tempArr)
	for tempArr.Next() {
		var (
			id     int64
			name   string
			age    int64
			rollno int64
		)
		checkerror.Checkerror(tempArr.Scan(&id, &name, &age, &rollno))
		var student = Student{id, name, age, rollno}
		tempArray = append(tempArray, student)
	}
	Students = tempArray
	checkerror.Checkerror(tempArr.Err())
	return Students
}

func Filter(index int64) Student {
	result := connectdb.GiveDB().QueryRow("Select * from users where id = $1", index)
	var (
		id     int64
		name   string
		age    int64
		rollno int64
	)
	result.Scan(&id, &name, &age, &rollno)
	fmt.Println(result)

	student := Student{id, name, age, rollno}
	return student
}

func AddStudent(_students Student) int64 {
	result, err := connectdb.GiveDB().Exec(`INSERT INTO users(name,age,rollno) VALUES($1,$2,$3)`, _students.Name, _students.Age, _students.Rollno)

	checkerror.Checkerror(err)

	var rowChanged, err1 = result.RowsAffected()

	checkerror.Checkerror(err1)


	if rowChanged > 0 {
		res := connectdb.GiveDB().QueryRow(`Select last_value from users_id_seq`)
		var id int64
		res.Scan(&id)
		// fmt.Println(`Student = `, Student{id, name, age, rollno})
		return id
	}
	return 0
}
func Remove(i int64) {
	_, err := connectdb.GiveDB().Exec("DELETE FROM users WHERE id = $1", i)
	checkerror.Checkerror(err)

}
func UpdateStudent(oldStudent Student, updatedStudent Student, id int64) Student {
	// If the default values are provided then don't change it
	if updatedStudent.Age == 0 {
		updatedStudent.Age = oldStudent.Age
	}
	if updatedStudent.Rollno == 0 {
		updatedStudent.Rollno = oldStudent.Rollno
	}
	if len(updatedStudent.Name) == 0 {
		updatedStudent.Name = oldStudent.Name
	}

	connectdb.GiveDB().Exec("update users set name = $1 ,age =$2 ,rollno = $3 where id = $4", updatedStudent.Name, updatedStudent.Age, updatedStudent.Rollno, id)
	GiveStudents()
	return updatedStudent
}
