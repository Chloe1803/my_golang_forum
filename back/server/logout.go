package forum

import (
	f "forum/back/func/queries"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	//Récupère le cookie
	cookie, err := r.Cookie("userdata")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//Supprime la session de la DB
	err = f.DeleteUserSession(cookie.Value)
	if err != nil {
		Error(w, r, err, "Erreur dans la suppression du la session", http.StatusInternalServerError)
	}

	// Supprimer le cookie
	deleteCookie := &http.Cookie{
		Name:    "userdata",
		Value:   "",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, deleteCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
