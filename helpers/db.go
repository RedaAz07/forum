package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"forum/utils"

	_ "github.com/mattn/go-sqlite3"
)

func DataBase() {
	var err error
	utils.Db, err = sql.Open("sqlite3", "./db/db.db")
	if err != nil {
		log.Fatal("open error:", err)
	}
	//	defer utils.Db.Close()

	sqlfile, err := os.ReadFile("./db/query.sql")
	if err != nil {
		log.Fatal("read error:", err)
	}

	_, err = utils.Db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal("exec error: ", err)
	}
	// utils.Db.Exec("delete from posts where description = 'assasas'")
	fmt.Println("Queries executed successfully!")
}
