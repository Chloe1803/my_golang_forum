package forum

import (
	f "forum/back/func"
	q "forum/back/func/queries"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("./front/tmpl/index.html"))

	var data q.Index_data
	var filtered_categories []int

	//Rate limite

	rateLimiter := f.NewRateLimiter(1, 1, 2)
	if !rateLimiter.IsAllowed() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	//Récuperation du cookie et du statue de connexion
	data.Connected = false

	user_id, err := f.FetchCookie(r)
	if err != nil {
		Error(w, r, err, "Problème récupération de cookie", http.StatusInternalServerError)
		return
	}

	if user_id != 0 {
		data.Connected = true
	}

	// Récupération des catégories à afficher

	data.Categories, err = q.FetchCategories()
	if err != nil {
		Error(w, r, err, "Erreur lors de l'import des catégories", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {

		//Récupération des filtres à appliquer
		filters := r.URL.Query()

		if len(filters) != 0 {
			for key, value := range filters {
				switch {
				case key == "Filtered_by_user_post":
					data.Index_posts, err = q.GetPostOfUser(user_id)
					if err != nil {
						Error(w, r, err, "Problème Requête GetPostOfUser", http.StatusInternalServerError)
						return
					}
				case key == "Filtered_by_user_like":
					data.Index_posts, err = q.GetPostLikedbyUser(user_id)
					if err != nil {
						Error(w, r, err, "Problème Requête GetPostLikedbyUser", http.StatusInternalServerError)
						return
					}
				default:
					cat, err6 := strconv.Atoi(value[0])
					if err6 != nil {
						Error(w, r, err, "Problème Atoi", http.StatusInternalServerError)
					}
					filtered_categories = append(filtered_categories, cat)
				}
			}

			if len(filtered_categories) != 0 {
				data.Index_posts, err = q.GetPostbyCat(filtered_categories)
				if err != nil {
					Error(w, r, err, "Problème Requête PostbyCat", http.StatusInternalServerError)
					return
				}
			}
		} else {
			data.Index_posts, err = q.GetPosts()
			if err != nil {
				Error(w, r, err, "Problème Requête Get", http.StatusInternalServerError)
				return
			}
		}

		//Exécution du template

		err = tmpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			Error(w, r, err, "Erreur lors de l'exécution du template Index", http.StatusInternalServerError)
			return
		}
	} else {
		//Erreur pour tout autre méthode
		Error(w, r, nil, "Bad Method", http.StatusBadRequest)
		return
	}
}
