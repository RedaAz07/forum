package helpers

import (
	"database/sql"
	"fmt"
	"forum/utils"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func DataBase() {
	var err error
	utils.Db, err = sql.Open("sqlite", "./db/db.db")
	if err != nil {
		log.Fatal("open error:", err)
	}

	sqlfile, err := os.ReadFile("./db/query.sql")
	if err != nil {
		log.Fatal("read error:", err)
	}

	_, err = utils.Db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal("exec error: ", err)
	}




/*	stmt := `INSERT INTO users (username, email, password) VALUES (?, ?, ?);`

	_, err = utils.Db.Exec(stmt, "redaanniz", "reda@gmail.com", "123456789")
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println("Queries executed successfully!")
}
