package db

import (
	"database/sql"
	"os"

	//This is driver need to connect to mysqldb  server.
	_ "github.com/go-sql-driver/mysql"
)

var (
	//DB that holds database connection
	DB *sql.DB
)

func init() {
	url, _ := os.LookupEnv("MYSQL_DB_URL")
	//url := "root:my-secret-pw@tcp(127.0.0.1:3306)/main"
	DB = connect(url)
}

func connect(cstring string) *sql.DB {
	conn, err := sql.Open("mysql", cstring)

	checkErr(err)

	return conn
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
