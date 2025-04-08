package handlers

import (
	"forum/helpers"
	"forum/utils"
	"net/http"
)

func LoginShowHandler(w http.ResponseWriter, r *http.Request) {

if exists , _ :=helpers.SessionChecked(w,r) ; exists{
	http.Redirect(w,r,"/", 302)
	return
}


if r.URL.Path != "/login" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusNotFound, utils.ErrorNotFound)
		return
}


	if r.Method != "GET" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}


	helpers.RanderTemplate(w, "login.html", http.StatusOK, nil)

}
