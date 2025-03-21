package handlers

import (
	"forum/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	
	helpers.RanderTemplate(w, "home.html", 200, nil)
}
