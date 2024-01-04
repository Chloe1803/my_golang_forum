package forum

import (
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
	var action string
	var args []interface{}

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

	//Gestion des likes et dislikes sur les postes et les commentaires

	if r.Method == http.MethodGet {
		if r.FormValue("post_action")!=""{
			action = r.FormValue("post_action")
			args = append(args,"posts", "like_post", "post_id", data.Post.ID, user_id)
		}else if r.FormValue("comment_action")!=""{
			action = r.FormValue("comment_action")
			com_id, err3 := strconv.Atoi(r.FormValue("Comment_id"))
			if err3 != nil {
				Error(w, r, err, "Problème Atoi comment_id", http.StatusInternalServerError)
			}
			args = append(args, "comments", "like_comment", "comment_id", com_id, user_id)
		}

		switch action {
		case "like" :
			f.Like(args)
		case "dislike" :
			f.Dislike(args)
		case "unlike" :
			f.Unlike(args)
		case "undislike":
			f.Undislike(args)
		case "change_to_like" :
			f.Undislike(args)
			f.Like(args)
		case "change_to_dislike" :
			f.Unlike(args)
			f.Dislike(args)
		}
	}


	//Récupération des données du post
	data.Post, err = q.GetPostByID(data.Post.ID)
	if err != nil {
		Error(w, r, nil, "Error : récupération du post", http.StatusInternalServerError)
		return
	}

	//Vérifie le status des likes
	data.Like_status, err = q.CheckLike("post_id", data.Post.ID, user_id, "like_post")
	if err != nil {
		data.Like_status = 0
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
