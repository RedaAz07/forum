package handlers

import (
	"forum/helpers"
	"forum/utils"
	"net/http"
)

func RegisterShowHandler(w http.ResponseWriter, r *http.Request) {
	if exists , _ :=helpers.SessionChecked(w,r) ; exists {
http.Redirect(w,r,"/", 302)
		return
	}
	if r.URL.Path != "/register" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusNotFound, utils.ErrorNotFound)
	}

	if r.Method != "GET" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
	}

	helpers.RanderTemplate(w, "register.html", http.StatusOK, nil)

}
