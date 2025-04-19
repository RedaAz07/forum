package helpers

import (
	"bytes"
	"net/http"
	// زيد هادي
	"forum/utils"
)

func RanderTemplate(w http.ResponseWriter, template string, statusCode int, data interface{}) {
	var buf bytes.Buffer
	// execute the template with buffer to check if there is an error in our templates
	err := utils.Tp.ExecuteTemplate(&buf, template, data)
	if err != nil {
		buf.Reset()
		statusCode = http.StatusInternalServerError
		err := utils.Tp.ExecuteTemplate(&buf, "statusPage.html", utils.ErrorInternalServerErr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(PageDeleted()))
			return
		}
	}
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
