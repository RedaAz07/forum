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
	err = utils.Db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	sqlfile, err := os.ReadFile("./db/query.sql")
	if err != nil {
		log.Fatal("read error:", err)
	}

	_, err = utils.Db.Exec(string(sqlfile))
	if err != nil {
		log.Fatal("exec error: ", err)
	}
	utils.Db.Exec("DELETE FROM posts where title = 'Cumque voluptatem it'")
	utils.Db.Exec(`INSERT INTO categories (name, icon) VALUES('Sport', '<i class="fa-solid fa-medal"></i>'),('Music', '<i class="fa-solid fa-music"></i>'),('Movies', '<i class="fa-solid fa-film"></i>'),('Science', '<i class="fa-solid fa-flask"></i>'),('Gym', '<i class="fa-solid fa-dumbbell"></i>'),('Tecknology', '<i class="fa-solid fa-microchip"></i>'),('Culture', '<i class="fa-solid fa-person-walking"></i>'),('Politics', '<i class="fa-solid fa-landmark"></i>');`)

	fmt.Println("Queries executed successfully!")
}
