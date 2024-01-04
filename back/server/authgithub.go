package forum

import (
	"encoding/json"
	"fmt"
	f "forum/back/func"
	q "forum/back/func/queries"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Définir les constantes GitHub
const (
	githubClientID     = "1a3fdbb57dd1031f6a68"
	githubClientSecret = "36c8a887c0b900fcf6f94a841a61438c5e292f64"
	githubRedirectURL  = "http://localhost:8080/GitCallBack"
)

// Struct pour la réponse du token GitHub
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// Fonction pour gérer la redirection vers l'authentification GitHub
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	authURL := "https://github.com/login/oauth/authorize" +
		"?client_id=" + githubClientID +
		"&redirect_uri=" + githubRedirectURL +
		"&scope=user"
	http.Redirect(w, r, authURL, http.StatusFound)
}

// Fonction pour gérer la réponse de GitHub après l'authentification
func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	// Effectue l'échange du code d'autorisation, obtenu après que l'utilisateur a autorisé l'application, contre un jeton d'accès GitHub

	//Préparation des données pour la requête POST
	data := url.Values{
		"client_id":     {githubClientID},
		"client_secret": {githubClientSecret},
		"code":          {code},
		"redirect_uri":  {githubRedirectURL},
	}

	//Envoie de la requête POST à l'URL avec les données, ce qui va permettre d'échanger le code d'audorisation contre un jeton d'accès
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", data)
	if err != nil {
		Error(w, r, err, "Erreur lors de la requête d'échange de code", http.StatusInternalServerError)
		return
	}

	//Lecture de la réponse qui contient le jeton d'accès GitHub
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Error(w, r, err, "Erreur lors de la lecture du corps de la réponse", http.StatusInternalServerError)
		return
	}

	// Traitement de la réponse obtenu pour obtenir le jeton d'accès

	//Analyse des données de la réponse en les transformant en paramètre utilisable
	formData, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		Error(w, r, err, "Erreur lors de l'analyse des données", http.StatusInternalServerError)
		return
	}

	// Extraction du jeton d'accès
	accessToken := formData.Get("access_token")
	if accessToken == "" {
		fmt.Println("Aucun jeton d'accès trouvé dans la réponse")
		return
	}

	//Requête pour obtenir les informations utilisateur avec le jeton d'accès
	userInfoURL := "https://api.github.com/user"
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		Error(w, r, err, "Erreur lors de la création de la requête utilisateur", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	//Envoi de la requête et traitement de la réponse
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		Error(w, r, err, "Erreur lors de la demande d'informations utilisateur", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Vérification due code de statut de la réponse
	if resp.StatusCode != http.StatusOK {
		Error(w, r, nil, "La demande d'informations utilisateur a renvoyé un code de statut non OK : ", resp.StatusCode)
		return
	}

	// Lecture des informations de l'utilisateur depuis la réponse, dans l'API Github le login = username, l'avatarURL est l'emplacement de l'image de profil
	var userData struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		Error(w, r, err, "Erreur lors de la lecture des informations utilisateur JSON", http.StatusInternalServerError)
		return
	}

	//Ajoute l'utilisateur à la base de données user et session_user, et créer le cookie

	// Insère l'ID GitHub dans la base de données uniquement s'il n'existe pas déjà
	var user_id int
	var err3 error
	if userData.Login != "" {
		user_id, err3 = q.AddGitHubUser(userData.Login)
		if err3 != nil {
			Error(w, r, err, "Erreur lors de l''ajout à la DB", http.StatusInternalServerError)
			return
		}
	} else {
		Error(w, r, nil, "Aucune Infos récupéré de Github", http.StatusInternalServerError)
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
