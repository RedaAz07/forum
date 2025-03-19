package handlers

import (
	"forum/helpers"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RanderTemplate(w, "login.html", 200, nil)

}
