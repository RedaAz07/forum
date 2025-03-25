package handlers

import (
	"fmt"
	"forum/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {


	session, _ := r.Cookie("session")



	fmt.Println("sses0", session)

	helpers.RanderTemplate(w, "home.html", 200, session)
}
