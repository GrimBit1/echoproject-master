package connectdb

import (
	"database/sql"
	checkerror "server/checkError"
)

var db *sql.DB
var err error

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
