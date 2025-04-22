package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"forum/utils"

	_ "modernc.org/sqlite"
)

func DataBase() {
	var err error
	utils.Db, err = sql.Open("sqlite", "./db/db.db")
	if err != nil {
		log.Fatal("open error:", err)
	}
	//	defer utils.Db.Close()

	sqlfile, err := os.ReadFile("./db/query.sql")
	if err != nil {
		log.Fatal("read error:", err)
	}

	// i should to use     status ENUM('depend', 'refus', 'success') NOT NULL look chatgpt
	_, err = utils.Db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal("exec error: ", err)
	}

	fmt.Println("Queries executed successfully!")
}
