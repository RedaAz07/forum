package handlers

import (
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func RegisterShowHandler(w http.ResponseWriter, r *http.Request) {
	if exists, _ := helpers.SessionChecked(w, r); exists {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
	}

	helpers.RanderTemplate(w, "register.html", http.StatusOK, nil)
}
