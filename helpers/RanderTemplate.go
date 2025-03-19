package helpers

import (
	"bytes"
	"forum/utils"
	"net/http"
)

func RanderTemplate(w http.ResponseWriter, template string, statusCode int, data interface{}) {

	var buffer bytes.Buffer

	err := utils.Tp.ExecuteTemplate(&buffer, template, data)
	if err != nil {
		buffer.Reset()
		statusCode = http.StatusInternalServerError

		err := utils.Tp.ExecuteTemplate(&buffer, "statusPage.html", data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return

		}

	}
	w.WriteHeader(statusCode)
	w.Write(buffer.Bytes())

}
