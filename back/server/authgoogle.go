package forum

import (
	"encoding/json"
	"fmt"
	f "forum/back/func"
	q "forum/back/func/queries"
	"net/http"
	"strings"
)

var (
	clientID     = "422408899006-86v80fdaehbnbh3a9vfno0livpcmk78c.apps.googleusercontent.com"
	clientSecret = "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"
	redirectURI  = "http://localhost:8080/auth"
)

//Fonction pour gérer la redirection vers l'authentification Google
func HandleLoginGoogle(w http.ResponseWriter, r *http.Request) {
	url := "https://accounts.google.com/o/oauth2/auth" +
		"?client_id=" + clientID +
		"&redirect_uri=" + redirectURI +
		"&response_type=code" +
		"&scope=email%20profile"
	http.Redirect(w, r, url, http.StatusFound)
}

// Fonction pour gérer la réponse de Google après l'authentification
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	//Récupération du code d'accès
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code d'autorisation manquant. Connexion avec Google requise.", http.StatusUnauthorized)
		return
	}

	//Préparation des données pour la requête POST
	tokenURL := "https://accounts.google.com/o/oauth2/token"
	data := "code=" + code +
		"&client_id=" + clientID +
		"&client_secret=" + clientSecret +
		"&redirect_uri=" + redirectURI +
		"&grant_type=authorization_code"

	//Envoie de la requête POST à l'URL avec les données, ce qui va permettre d'échanger le code d'audorisation contre un jeton d'accès
	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Printf("Erreur lors de la demande de token: %v", err)
		Error(w, r, err, "Erreur lors de la connexion avec Google", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//Vérification du statut de la réponse
	if resp.StatusCode != http.StatusOK {
		Error(w, r, err, "Erreur lors de la connexion avec Google", http.StatusInternalServerError)
		return
	}

	// Extraction du jeton d'accès
	var tokenData struct {
		AccessToken string `json:"access_token"`
	}
	err = json.NewDecoder(resp.Body).Decode(&tokenData)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse JSON: %v", err)
		Error(w, r, err, "Erreur lors de la lecture de la réponse de Google", http.StatusInternalServerError)
		return
	}

	//Requête pour obtenir les informations utilisateur avec le jeton d'accès
	userInfoURL := "https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + tokenData.AccessToken
	resp, err = http.Get(userInfoURL)
	if err != nil {
		fmt.Printf("Erreur lors de la demande d'informations utilisateur: %v", err)
		Error(w, r, err, "Erreur lors de la demande d'informations utilisateur à Google", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lecture des informations de l'utilisateur depuis la réponse, dans l'API Google le nom et l'email de l'utilisateur
	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture des informations utilisateur JSON: %v", err)
		Error(w, r, err, "Erreur lors de la lecture des informations utilisateur depuis Google", http.StatusInternalServerError)
		return
	}

	username := strings.ReplaceAll(userInfo.Name, " ", "_")

	//Ajoute l'utilisateur à la base de données user et session_user, et créer le cookie

	// Insère l'utilisateur dans la base de données uniquement s'il n'existe pas déjà
	user_id, err3 := q.AddGoogleUser(username, userInfo.Email)
	if err3 != nil {
		fmt.Println(err)
		Error(w, r, err, "Erreur lors de l'ajout à la DB", http.StatusInternalServerError)
		return
	}

	// Création d'une nouvelle session
	sessionID, err1 := q.UpdateSessionsDB(user_id)
	if err1 != nil {
		Error(w, r, err1, "Erreur lors de la mise à jour d'une session", http.StatusInternalServerError)
		return
	}

	// avec cette nouvelle session on va généré un cookie
	cookie := f.CreateCookie(sessionID)
	http.SetCookie(w, cookie)

	//Redirige l'utilisateur une fois qu'il est connecté
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
