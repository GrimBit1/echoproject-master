package studentlogic

import (
	_ "github.com/lib/pq"
)

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Rollno int64  `json:"rollno"`
}
type Error struct {
	Message string `json:"message"`
}

const UserSchema = `
CREATE TABLE users(
	id Bigserial PRIMARY KEY,
	name varchar(150),
	age int,
	rollno int
);`
