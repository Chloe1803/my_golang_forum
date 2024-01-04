package forum

import (
	"fmt"
	f "forum/back/func"
	q "forum/back/func/queries"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func ViewPost(w http.ResponseWriter, r *http.Request) {

	var data q.VP_data
	tmpl := template.Must(template.ParseGlob("./front/tmpl/viewpost.html"))
	var err error

	//Rate limite

	rateLimiter := f.NewRateLimiter(1, 1, 2)
	if !rateLimiter.IsAllowed() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	//Récupère le cookie
	data.Connected = false
	user_id, err := f.FetchCookie(r)
	if err != nil {
		Error(w, r, err, "Problème récupération de cookie", http.StatusInternalServerError)
		return
	}
	if user_id != 0 {
		data.Connected = true
	}

	//Vérification et récupération du post_id
	if r.URL.Query().Has("Post_id") {
		data.Post.ID, err = strconv.Atoi(r.FormValue("Post_id"))
		if err != nil {
			Error(w, r, nil, "Error : invalid post_id in page viewpost URL", http.StatusBadRequest)
			return
		}

	} else {
		Error(w, r, nil, "Error : post_id not specified in page viewpost", 404)
		return
	}

	if r.Method == http.MethodPost {
		//Gestion de l'ajout de commentaire
		if r.FormValue("comment_content") != "" {
			com_content := r.FormValue("comment_content")
			com_TimeStamp := time.Now()

			err = q.AddNewComment(user_id, data.Post.ID, com_content, com_TimeStamp)
			if err != nil {
				Error(w, r, err, "Erreur lors de l'ajout d'un commentaire", http.StatusInternalServerError)
				return
			}
			err = q.UpdateNbLike("comments", "post_id", data.Post.ID, "comments", 1)
			if err != nil {
				Error(w, r, err, "Erreur lors de l'ajout d'un commentaire", http.StatusInternalServerError)
				return
			}
		}
	}

	if r.Method == http.MethodGet {

		switch {
		// Gestion des likes/dislikes des postes
		case r.FormValue("p_like") == "true":
			// Pour un like en état 0
			err = f.Like("posts", "like_post", "post_id", data.Post.ID, user_id)
		case r.FormValue("p_dislike") == "true":
			// Pour un dislike en état 0
			err = f.Dislike("posts", "like_post", "post_id", data.Post.ID, user_id)

		case r.FormValue("p_unlike") == "true":
			//Si l'utilisateur reclique sur Like après l'avoir liker en état 1
			err = f.Unlike("posts", "like_post", "post_id", data.Post.ID, user_id)

		case r.FormValue("p_change_to_dislike") == "true":
			//Si l'utilisateur clique sur dislike après avoir cliquer sur like
			err = f.Unlike("posts", "like_post", "post_id", data.Post.ID, user_id)
			if err != nil {
				Error(w, r, err, "Problème Unlike", http.StatusInternalServerError)
			}
			err = f.Dislike("posts", "like_post", "post_id", data.Post.ID, user_id)

		case r.FormValue("p_change_to_like") == "true":
			//Si l'utilisateur clique sur like après avoir cliquer sur dislike
			err = f.Undislike("posts", "like_post", "post_id", data.Post.ID, user_id)
			if err != nil {
				Error(w, r, err, "Problème Unlike", http.StatusInternalServerError)
			}
			err = f.Like("posts", "like_post", "post_id", data.Post.ID, user_id)

		case r.FormValue("p_undislike") == "true":
			//Si l'utilisateur reclique sur disLike après l'avoir disliker
			err = f.Undislike("posts", "like_post", "post_id", data.Post.ID, user_id)

			//Gestion des likes/dislikes des commentaires -----------------------------------------------------------------

		case r.FormValue("com_like") == "true":
			// Pour un like en état 0
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}
			err = f.Like("comments", "like_comment", "comment_id", com_id, user_id)

		case r.FormValue("com_dislike") == "true":
			// Pour un dislike en état 0
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}

			err = f.Dislike("comments", "like_comment", "comment_id", com_id, user_id)

		case r.FormValue("com_unlike") == "true":
			//Si l'utilisateur reclique sur Like après l'avoir liker en état 1
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}

			err = f.Unlike("comments", "like_comment", "comment_id", com_id, user_id)

		case r.FormValue("com_change_to_dislike") == "true":
			//Si l'utilisateur clique sur dislike après avoir cliquer sur like
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}

			err = f.Unlike("comments", "like_comment", "comment_id", com_id, user_id)
			if err != nil {
				Error(w, r, err, "Problème Unlike", http.StatusInternalServerError)
			}
			err = f.Dislike("comments", "like_comment", "comment_id", com_id, user_id)

		case r.FormValue("com_change_to_like") == "true":
			//Si l'utilisateur clique sur like après avoir cliquer sur dislike
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}

			err = f.Undislike("comments", "like_comment", "comment_id", com_id, user_id)
			if err != nil {
				Error(w, r, err, "Problème Unlike", http.StatusInternalServerError)
			}
			err = f.Like("comments", "like_comment", "comment_id", com_id, user_id)

		case r.FormValue("com_undislike") == "true":
			//Si l'utilisateur reclique sur disLike après l'avoir disliker
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				fmt.Println("Problème Atoi com_id")
			}

			err = f.Undislike("comments", "like_comment", "comment_id", com_id, user_id)

		}
		if err != nil {
			Error(w, r, err, "Problème like/dislike", http.StatusInternalServerError)
		}
	}

	//Récupération des données du post
	data.Post, err = q.GetPostByID(data.Post.ID)
	if err != nil {
		Error(w, r, nil, "Error : récupération du post", http.StatusInternalServerError)
		return
	}

	//Vérifie le status des likes
	data.Post.Like_status, err = q.CheckLike("post_id", data.Post.ID, user_id, "like_post")
	if err != nil {
		data.Post.Like_status = 0
	}

	//Récupère les commentaires du post
	data.Comments, err = q.GetComments(data.Post.ID)
	if err != nil {
		Error(w, r, nil, "Error : récupération des commentaires du post", http.StatusInternalServerError)
		return
	}

	//Vérifie le status des likes de commentaire

	for i := 0; i < len(data.Comments); i++ {
		data.Comments[i].Like_status, err = q.CheckLike("comment_id", data.Comments[i].ID, user_id, "like_comment")
		if err != nil {
			data.Comments[i].Like_status = 0
		}
		data.Comments[i].Connected = data.Connected
	}
	err = tmpl.ExecuteTemplate(w, "viewpost.html", data)
	if err != nil {
		Error(w, r, err, "Erreur lors de l'exécution du template Register", http.StatusInternalServerError)
		return
	}

}
