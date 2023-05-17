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
func UpdateStudent(student Student, name string, id int64) Student {
	var index int64 = GiveIndex(id)

	student.Name = name
	Students[index] = student
	return student
}
