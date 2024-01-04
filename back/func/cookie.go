package forum

import (
	q "forum/back/func/queries"
	"net/http"
	"strconv"
	"time"
)

func CreateCookie(sessionID int) *http.Cookie {
	// Convertir l'ID de session (int) en string
	sessionIDString := strconv.Itoa(sessionID)

	// Créer le cookie avec l'ID de session
	cookie := &http.Cookie{
		Name:    "userdata",
		Value:   sessionIDString,
		Expires: time.Now().Add(24 * time.Hour), // Exemple : expiration dans 24 heures
		// Ajoutez d'autres options si nécessaire
	}
	return cookie
}

func FetchCookie(r *http.Request) (int, error) {

	cookie, err := r.Cookie("userdata")
	if err != nil {
		return 0, nil
	}

	if cookie.Value != "" {
		user_ID, err1 := q.GetUserIDBySessionID(cookie.Value)
		if err1 != nil {
			return 0, nil
		}
		return user_ID, nil
	}

	return 0, nil
}
