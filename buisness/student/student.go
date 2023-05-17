package studentlogic

var Index int64 = 1
var IndexPointer = &Index
var Students = []Student{}

func Filter(slice []Student, index int64) Student {
	for i := range Students {
		if Students[i].Index == index {
			return Students[i]
		}
	}
	return Student{}
}
func GiveIndex(slice []Student, index int64) int64 {
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
	return append(s[:i], s[i+1:]...)
}
