package forum

import (
	"fmt"
	"net/http"
)

func Handlers() {
	// Page d'accueil
	http.HandleFunc("/", IndexHandler)
	//Page de connexion
	http.HandleFunc("/login", LoginHandler)
	//Requête de connexion avec Github
	http.HandleFunc("/loginwithgithub", HandleGitHubLogin)
	http.HandleFunc("/GitCallBack", HandleGitHubCallback)
	//Requête de connexion avec Google
	http.HandleFunc("/loginwithgoogle", HandleLoginGoogle)
	http.HandleFunc("/auth", HandleGoogleCallback)
	//Page d'inscription
	http.HandleFunc("/register", RegisterHandler)
	//Requête de déconnexion
	http.HandleFunc("/logout", LogoutHandler)
	//Page pour écrire une publication
	http.HandleFunc("/write", WriteHandler)
	//Page pour voir et commenter une publication
	http.HandleFunc("/view_post", ViewPost)

	fmt.Println("Listening to port 8080... https://localhost:8080")

}
