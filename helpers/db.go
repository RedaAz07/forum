package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB 

func Db() {
	var err error
	db, err = sql.Open("sqlite3", "./db/db.db") 
	if err != nil {
		log.Fatal("open error:", err)
	}

	sqlfile, err := os.ReadFile("./db/query.sql")
	if err != nil {
		log.Fatal("read error:", err)
	}

	_, err = db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal("exec error:", err)
	}

	fmt.Println("Queries executed successfully!")
}
