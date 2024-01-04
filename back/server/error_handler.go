package forum

import (
	"net/http"
	"strconv"
	"text/template"
)

func Error(w http.ResponseWriter, r *http.Request, err error, message string, errorStatus int) {
	tmpl, _ := template.ParseFiles("./front/tmpl/error.html")

	s := ""

	if err != nil {
		s += err.Error() + "\n"
	}

	if message != "" {
		s += message + "\n"
	}

	s += "http status : " + strconv.Itoa(errorStatus)

	w.WriteHeader(errorStatus)
	tmpl.Execute(w, s)
}
