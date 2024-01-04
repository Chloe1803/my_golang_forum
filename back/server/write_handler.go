package forum

import (
	f "forum/back/func"
	q "forum/back/func/queries"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

func WriteHandler(w http.ResponseWriter, r *http.Request) {
	var data q.Index_data

	tmpl := template.Must(template.ParseGlob("./front/tmpl/writepost.html"))

	//Rate limite

	rateLimiter := f.NewRateLimiter(1, 1, 2)
	if !rateLimiter.IsAllowed() {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	// Récupération des catégories à afficher
	var err error

	data.Categories, err = q.FetchCategories()
	if err != nil {
		Error(w, r, err, "Erreur lors de l'import des catégories", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodPost:

		//Récupère les données du postes
		title := r.FormValue("title")
		content := r.FormValue("content")
		time_stamp := time.Now()

		user_id, err := f.FetchCookie(r)
		if err != nil {
			Error(w, r, err, "Problème récupération de cookie", http.StatusInternalServerError)
			return
		}

		//Insère les données dans la table "posts"

		res, err2 := q.AddnewPost(title, user_id, time_stamp, content)
		if err2 != nil {
			Error(w, r, err2, "Problème d'ajout du post à la db", http.StatusInternalServerError)
			return
		}
		post_id, err1 := res.LastInsertId()
		if err1 != nil {
			Error(w, r, err1, "Problème d'ajout du post à la db", http.StatusInternalServerError)
			return
		}

		//Récupère les catégories et les insères dans la table "post_catégorie"

		for _, cat := range data.Categories {
			if r.FormValue(cat.Name) != "" {
				cat_ID, err3 := strconv.Atoi(r.FormValue(cat.Name))
				if err3 != nil {
					Error(w, r, err3, "Problème Atoi cat_ID", http.StatusInternalServerError)
					return
				}
				err3 = q.AddCat_Post(int(post_id), cat_ID)
				if err3 != nil {
					Error(w, r, err3, "Problème ajout cat_post", http.StatusInternalServerError)
					return
				}
			}
		}

		//Récupération et traitement de l'image

		imageFile, imageHeader, err := r.FormFile("post_image")
		var imageType string
		if err != nil && err != http.ErrMissingFile {
			// Une erreur inattendue s'est produite lors du téléchargement de l'image.
			Error(w, r, err, "Error uploading image in createpost", http.StatusInternalServerError)
			return
		}
		//Vérifie s'il y a une image
		if imageHeader != nil {
			defer imageFile.Close()

			// type MIME de l'image téléchargée
			imageType = imageHeader.Header.Get("Content-Type")
			// Vérifie que le type d'image est autorisé (JPEG, PNG ou GIF).
			if imageType != "image/jpeg" && imageType != "image/png" && imageType != "image/gif" {
				// Type d'image non autorisé.
				Error(w, r, nil, "Unsupported image format in createpost", http.StatusUnsupportedMediaType)
				return
			}
			// Construit le chemin complet du fichier image en utilisant le nom de fichier unique.
			uniqueImageFilename := "imageOfPostnumber" + strconv.Itoa(int(post_id))
			imageFilePath := "img/" + uniqueImageFilename
			img, err := os.Create("./front/static/" + imageFilePath)
			if err != nil {
				Error(w, r, err, "Error saving image in createpost", http.StatusInternalServerError)
				return
			}
			defer img.Close()
			_, err = io.Copy(img, imageFile)
			if err != nil {
				Error(w, r, err, "Error copying image in createpost", http.StatusInternalServerError)
				return
			}

			//MAJ de la base de données avec l'image
			err = q.UpdatePostImagePathandType(imageType, imageFilePath, int(post_id))
			if err != nil {
				Error(w, r, err, "Erreur lors de l'ajout des images à la db", http.StatusInternalServerError)
				return
			}
		}

		//Redirection vers la page d'accueil

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodGet:

		err = tmpl.ExecuteTemplate(w, "writepost.html", data)
		if err != nil {
			Error(w, r, err, "Erreur lors de l'exécution du template Register", http.StatusInternalServerError)
			return
		}

	}
}
