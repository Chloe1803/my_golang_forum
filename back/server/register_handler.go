package forum

import (
	f "forum/back/func"
	q "forum/back/func/queries"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var errorMessage bool
	tmpl := template.Must(template.ParseGlob("./front/tmpl/register.html"))

	//Rate limite

	rateLimiter := f.NewRateLimiter(1, 1, 2)
	if !rateLimiter.IsAllowed() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	switch r.Method {
	case http.MethodPost:

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		var errhashed error

		password, errhashed = HashPassword(password)
		if errhashed != nil {
			Error(w, r, errhashed, "Erreur lors du hashage du mot de passe", http.StatusInternalServerError)
		}

		check, err := f.CheckEmailAndUsername(email, username)
		if err != nil {
			Error(w, r, err, "Erreur lors de la vérification de l'email et l'utilisateur", http.StatusInternalServerError)
		}

		if check {

			err = q.AddNewUser(email, username, password)
			if err != nil {
				Error(w, r, err, "Erreur lors de la vérification de l'email et l'utilisateur", http.StatusInternalServerError)
				return
			}

			// Connect le nouvelle utilisateur
			_, err = f.Login(email, password, w)
			if err != nil {
				Error(w, r, err, "Erreur lors de la connexion du nouvel utilisateur", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		} else {

			var data q.ErrorInput
			data.ErrorMessage = true

			err := tmpl.ExecuteTemplate(w, "register.html", data)
			if err != nil {
				Error(w, r, err, "Erreur lors de l'exécution du template Register", http.StatusInternalServerError)
				return
			}
		}

	case http.MethodGet:

		data := q.ErrorInput{
			ErrorMessage: errorMessage,
		}

		// Exécuter le template pour afficher la page d'inscription
		err := tmpl.ExecuteTemplate(w, "register.html", data)
		if err != nil {
			Error(w, r, err, "Erreur lors de l'exécution du template Register", http.StatusInternalServerError)
			return
		}
	}
}

func HashPassword(password string) (string, error) {
	// Utilisez bcrypt pour hacher le mot de passe avec le coùut du hashage, qui définit le niveau de complexité
	//bcrypt.DefaultCost est une constante prédéfinie pour un coût de hachage raisonnable.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
