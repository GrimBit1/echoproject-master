package studentlogic

type Student struct {
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Rollno int64  `json:"rollno"`
	Index  int64  `json:"index"`
}
type Error struct {
	Message string `json:"message"`
}