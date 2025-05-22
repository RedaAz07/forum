package handlers

import (
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func LoginShowHandler(w http.ResponseWriter, r *http.Request) {
	if exists, _ := helpers.SessionChecked(w, r); exists {
		http.Redirect(w, r, "/", 303)
		return
	}

	if r.Method != "GET" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}

	helpers.RanderTemplate(w, "login.html", http.StatusOK, nil)
}
