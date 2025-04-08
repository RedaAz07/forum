package handlers

import (
	"fmt"
	"net/http"

	"forum/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := r.Cookie("session")

	fmt.Println("sses0", session)

	stmt := "SELECT * from posts "
	

	helpers.RanderTemplate(w, "home.html", 200, session)
}
