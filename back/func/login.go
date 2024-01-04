package forum

import (
	"fmt"
	q "forum/back/func/queries"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, IPpassword string, w http.ResponseWriter) (bool, error) {

	//Récupère le mot de passe et l'ID dans la base de données
	storedPassword, ID, err := q.GetIDandPasswordFromEmail(email)
	if err != nil {
		return false, err
	}

	//Comparaison des mots de passe avec bcrypte
	err = VerifyPassword(storedPassword, IPpassword)

	if err == nil {

		// si le mot de passe est correct, on va créer une nouvelle session
		sessionID, err1 := q.UpdateSessionsDB(ID)
		if err1 != nil {
			fmt.Println("Problème update session, ", err)
		}

		// avec cette nouvelle session on va généré un cookie
		cookie := CreateCookie(sessionID)
		http.SetCookie(w, cookie)

		return true, nil
	}

	return false, nil
}

func VerifyPassword(hashedPassword, password string) error {
	// Utilisez bcrypt pour comparer le mot de passe en clair avec le hash stocké.
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
