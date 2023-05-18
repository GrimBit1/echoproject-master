package studentlogic

var Index int64 = 1
var IndexPointer = &Index
var Students = []Student{}

func Filter(index int64) Student {
	for i := range Students {
		if Students[i].Index == index {
			return Students[i]
		}
	}
	return Student{}
}
func GiveIndex(index int64) int64 {
	var j int64 = 0
	for i := range Students {
		if Students[i].Index == index {
			return j
		}
		j++
	}
	return j
}
func AddStudent(_students Student) {
	Students = append(Students, _students)
}
func Remove(s []Student, i int64) []Student {
	return append(Students[:i], Students[i+1:]...)
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

	var index int64 = GiveIndex(id)

	Students[index] = updatedStudent
	return updatedStudent
}
