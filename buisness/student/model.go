package studentlogic

import (
	"database/sql"
	checkerror "server/checkError"

	_ "github.com/lib/pq"
)

type Student struct {
	Id     int64 `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Rollno int64  `json:"rollno"`
}
type Error struct {
	Message string `json:"message"`
}

var db *sql.DB
var err error

const UserSchema = `
CREATE TABLE users(
	id Bigserial PRIMARY KEY,
	name varchar(150),
	age int,
	rollno int
);`

func ConnectDB() {
	connStr := "user=nlab-7 dbname=nlab-7 password=nlab"
	db, err = sql.Open("postgres", connStr)
	checkerror.Checkerror(err)
	// result, err := db.Exec(UserSchema)
	// checkerror.Checkerror(err)
	// fmt.Println(result)
}
func GiveDB() *sql.DB {
	return db
}
