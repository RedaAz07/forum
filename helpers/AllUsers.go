package helpers

import (
	"fmt"
	"net/http"

	"forum/utils"
)

func AllUsers(w http.ResponseWriter) []string {
	// get all users
	var allusers []string
	var username string
	queryUsers := `select username from users order by username ASC`
	users, errUsers := utils.Db.Query(queryUsers)
	for users.Next() {
		errUsers = users.Scan(&username)
		if errUsers != nil {
			fmt.Println("DB Query error:", errUsers)
			RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return nil
		}
		allusers = append(allusers, username)
	}
	return allusers
	// end of get all usrers
}
