package forum

import (
	l "forum/back/func"
	q "forum/back/func/queries"
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("./front/tmpl/login.html"))

	var data q.ErrorInput
	data.ErrorMessage = false

	//Rate limite

	rateLimiter := l.NewRateLimiter(1, 1, 2)
	if !rateLimiter.IsAllowed() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	switch r.Method {
	case http.MethodPost:
		email := r.FormValue("email")
		password := r.FormValue("password")

		connected, err := l.Login(email, password, w)

		if err != nil {
			data.ErrorMessage = true
			err2 := tmpl.ExecuteTemplate(w, "login.html", data)
			if err2 != nil {
				Error(w, r, err2, "Problème d'exécution du template", http.StatusInternalServerError)
			}

			return
		}

		if !connected && err == nil {
			data.ErrorMessage = true
			err2 := tmpl.ExecuteTemplate(w, "login.html", data)
			if err2 != nil {
				Error(w, r, err2, "Problème d'exécution du template", http.StatusInternalServerError)
			}
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodGet:
		err := tmpl.ExecuteTemplate(w, "login.html", data)
		if err != nil {
			Error(w, r, err, "Erreur lors de l'exécution du template Register", http.StatusInternalServerError)
			return
		}
	}
}
