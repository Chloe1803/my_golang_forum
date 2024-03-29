package forum

import (
	"database/sql"
)

//Fonction qui initialise la connexion avec la DB
func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "back/database/forum.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

//Fonction pour ajouter des éléments à la DB
func QueryAddToDB(query string, args []interface{}) error {

	db, err := GetDatabase()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)

	return err
}

//Fonction pour ajouter des éléments à la DB et qui renvoie un résultat
func QueryAddToDBwithRes(query string, args []interface{}) (sql.Result, error) {

	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	res, err1 := db.Exec(query, args...)

	return res, err1
}


//Fonction qui consulte la DB et renvoie un entier
func QueryReturnOneNb(query string, args []interface{}) (int, error) {

	var res int

	db, err := GetDatabase()
	if err != nil {
		return 0, err
	}

	err = db.QueryRow(query, args...).Scan(&res)
	if err != nil {
		return 0, err
	}

	return res, err
}


func QueryReturnPasswordAndId(query string, args []interface{}) (string, int, error) {

	var password string
	var id int

	db, err := GetDatabase()
	if err != nil {
		return "", 0, err
	}

	err = db.QueryRow(query, args...).Scan(&password, &id)
	if err != nil {
		return "", 0, err
	}

	return password, id, err
}

func QueryTabCat(query string, args []interface{}) ([]Category, error) {

	var res []Category
	var cat Category

	db, err := GetDatabase()
	if err != nil {
		return res, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&cat.ID, &cat.Name)
		if err != nil {
			return res, err
		}
		res = append(res, cat)
	}

	return res, nil
}

func QueryTabPosts(query string, args []interface{}) ([]Post, error) {

	var res []Post
	var post Post

	db, err := GetDatabase()
	if err != nil {
		return res, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Author, &post.Timestamp, &post.Content, &post.Nb_like, &post.Nb_dislike, &post.Nb_comments, &post.ImageURL, &post.ImageType)
		if err != nil {
			return res, err
		}
		post.Categories, err = GetCategoriesForPost(post.ID)
		if err != nil {
			return res, err
		}
		if post.ImageType != "" {
			post.ImageAvailable = true
		} else {
			post.ImageAvailable = false
		}
		res = append(res, post)
	}

	return res, nil
}

func QueryVPpost(query string, args []interface{}) (Post, error) {

	var post Post

	db, err := GetDatabase()
	if err != nil {
		return post, err
	}

	err = db.QueryRow(query, args...).Scan(&post.ID, &post.Title, &post.Author, &post.Timestamp, &post.Content, &post.Nb_like, &post.Nb_dislike, &post.Nb_comments, &post.ImageURL, &post.ImageType)
	if err != nil {
		return post, err
	}
	post.Categories, err = GetCategoriesForPost(post.ID)
	if err != nil {
		return post, err
	}
	if post.ImageType != "" {
		post.ImageAvailable = true
	} else {
		post.ImageAvailable = false
	}

	return post, nil
}

func QueryComments(query string, args []interface{}) ([]Comment, error) {

	var com Comment
	var res []Comment

	db, err := GetDatabase()
	if err != nil {
		return res, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&com.ID, &com.Author, &com.Time_stamp, &com.Content, &com.Nb_like, &com.Nb_dislike)
		if err != nil {
			return res, err
		}
		res = append(res, com)
	}

	return res, nil
}
