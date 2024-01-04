package forum

import (
	"fmt"
	"strconv"
	"strings"
)

func NbEmailAndUser(email string, username string) (nb int, err error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ? OR username = ?"
	var args []interface{}
	args = append(args, email, username)
	return QueryReturnOneNb(query, args)
}

func GetIDandPasswordFromEmail(email string) (string, int, error) {
	query := "SELECT password, id FROM users WHERE email = ?"
	var args []interface{}
	args = append(args, email)
	return QueryReturnPasswordAndId(query, args)
}

func UpdateSessionsDB(ID int) (int, error) {

	var args []interface{}
	args = append(args, ID)

	queryOne := "INSERT OR REPLACE INTO user_sessions (user_id) VALUES (?);"

	err := QueryAddToDB(queryOne, args)
	if err != nil {
		return 0, err
	}

	queryTwo := "SELECT session_id FROM user_sessions WHERE user_id = ? ;"

	return QueryReturnOneNb(queryTwo, args)

}

func GetUserIDBySessionID(session_IDstr string) (int, error) {

	session_ID, err := strconv.Atoi(session_IDstr)
	if err != nil {
		return 0, err
	}
	var args []interface{}
	args = append(args, session_ID)

	query := "SELECT user_id FROM user_sessions WHERE session_id = ? ;"

	return QueryReturnOneNb(query, args)
}

func FetchCategories() ([]Category, error) {

	var args []interface{}

	query := "SELECT * FROM categories"

	return QueryTabCat(query, args)
}

//Requête de récupération des publications en fonction des filtres appliqués

func GetPostOfUser(user_id int) ([]Post, error) {
	var args []interface{}
	args = append(args, user_id)

	query := `SELECT P.post_id, P.title, C.username, P.created, P.content, P.likes, P.dislikes, P.comments, P.img_URL, P.img_type
	FROM posts AS P
	INNER JOIN users AS C
	ON P.user_id = C.id
	WHERE P.user_id = ? 
	ORDER BY P.created DESC`

	return QueryTabPosts(query, args)

}

func GetPostLikedbyUser(user_id int) ([]Post, error) {

	var args []interface{}
	args = append(args, user_id)

	query := `SELECT P.post_id, P.title, C.username, P.created, P.content, P.likes, P.dislikes, P.comments P.img_URL, P.img_type
	FROM posts AS P
	INNER JOIN users AS C
	ON P.user_id = C.id
	WHERE P.post_id IN
	(SELECT post_id
	FROM like_post
	WHERE user_id = ? ) 
	ORDER BY P.created DESC`

	return QueryTabPosts(query, args)

}

func GetPostbyCat(filtered_categories []int) ([]Post, error) {

	var args []interface{}
	for _, cat := range filtered_categories {
		args = append(args, cat)
	}

	query := `SELECT P.post_id, P.title, C.username, P.created, P.content, P.likes, P.dislikes, P.comments, P.img_URL, P.img_type
	FROM posts AS P
	INNER JOIN users AS C
	ON P.user_id = C.id
	WHERE P.post_id IN
	(SELECT post_id
	FROM post_category
	WHERE category_id IN (` + buildPlaceholders(len(filtered_categories)) + `)) 
	ORDER BY P.created DESC`

	return QueryTabPosts(query, args)
}

func GetPosts() ([]Post, error) {
	var args []interface{}
	query := `SELECT P.post_id, P.title, C.username, P.created, P.content, P.likes, P.dislikes, P.comments, P.img_URL, P.img_type
	FROM posts AS P
	INNER JOIN users AS C
	ON P.user_id = C.id
	ORDER BY P.created DESC`

	return QueryTabPosts(query, args)
}

func GetCategoriesForPost(ID int) ([]Category, error) {
	var args []interface{}
	args = append(args, ID)

	query := "SELECT R.category_id, C.category_name FROM post_category AS R INNER JOIN categories AS C ON R.category_id = C.category_id WHERE R.post_id = ?"

	return QueryTabCat(query, args)
}

func buildPlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := 0; i < count; i++ {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}

func GetPostByID(post_id int) (VP_post, error) {

	var args []interface{}
	args = append(args, post_id)

	query := `SELECT P.post_id, P.title, C.username, P.created, P.content, P.likes, P.dislikes, P.comments, P.img_URL, P.img_type
	FROM posts AS P
	INNER JOIN users AS C
	ON P.user_id = C.id
	WHERE P.post_id = ?
	ORDER BY P.created DESC`

	return QueryVPpost(query, args)
}

func CheckLike(type_id string, post_id int, user_id int, table string) (int, error) {

	var args []interface{}
	args = append(args, post_id, user_id)

	query := "SELECT type FROM " + table + " WHERE " + type_id + " = ? AND user_id = ?"

	return QueryReturnOneNb(query, args)
}

func GetComments(post_id int) ([]Comment, error) {
	var args []interface{}
	args = append(args, post_id)

	query := `SELECT C.comment_id, U.username, C.created, C.content, C.likes, C.dislikes
	FROM comments AS C
	INNER JOIN users AS U
	ON C.user_id = U.id
	WHERE C.post_id = ? `

	return QueryComments(query, args)
}

func GetNblikes(com_id int) {
	var args []interface{}
	args = append(args, com_id)
	query := "SELECT likes FROM comments WHERE comment_id = ?"

	nb, _ := QueryReturnOneNb(query, args)
	fmt.Println("nombre de like :", nb)
}
