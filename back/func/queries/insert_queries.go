package forum

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func AddNewUser(email string, username string, password string) error {

	query := "INSERT INTO users (email, username, password) VALUES (?, ?, ?)"

	var args []interface{}
	args = append(args, email, username, password)

	err := QueryAddUser(query, args)

	return err
}

func DeleteUserSession(sessionIDstr string) error {

	query := "DELETE FROM user_sessions WHERE session_id = ?"

	sessionID, err := strconv.Atoi(sessionIDstr)
	if err != nil {
		return err
	}

	var args []interface{}
	args = append(args, sessionID)

	err = QueryAddUser(query, args)

	return err
}

func AddnewPost(title string, user_id int, time_stamp time.Time, content string) (sql.Result, error) {

	var args []interface{}
	args = append(args, title, user_id, time_stamp, content)

	query := "INSERT INTO posts (title, user_id, created, content) VALUES (?, ?, ?, ?)"

	return QueryAddPost(query, args)
}

func AddCat_Post(post_id int, cat_ID int) error {

	var args []interface{}
	args = append(args, post_id, cat_ID)

	query := "INSERT INTO post_category (post_id, category_id) VALUES (?, ?)"

	return QueryAddUser(query, args)
}

func AddNewComment(user_id int, post_id int, com_content string, com_TimeStamp time.Time) error {
	var args []interface{}
	args = append(args, user_id, post_id, com_content, com_TimeStamp)

	query := "INSERT INTO comments (user_id, post_id, content, created) VALUES (?, ?, ?, ?)"

	return QueryAddUser(query, args)
}

func UpdateNbLike(table string, type_id string, post_id int, column string, add int) error {
	var args []interface{}
	args = append(args, add, post_id)

	query := fmt.Sprintf("UPDATE %s SET %s = %s + ? WHERE %s = ?", table, column, column, type_id)

	return QueryAddUser(query, args)
}

func UpdateStatusLike(table string, type_id string, id int, user_id int, status int) error {

	var args1 []interface{}
	args1 = append(args1, id, user_id)

	query1 := "SELECT COUNT(*) FROM " + table + " WHERE " + type_id + " = ? AND user_id = ?"

	count, err1 := QueryReturnOneNb(query1, args1)
	if err1 != nil {
		return err1
	}

	var args2 []interface{}
	var query2 string

	if count == 0 {
		args2 = append(args2, id, user_id, status)
		query2 = "INSERT INTO " + table + " (" + type_id + ", user_id, type) VALUES (?, ?, ?)"
		return QueryAddUser(query2, args2)
	} else {
		args2 = append(args2, status, id, user_id)
		query2 = "UPDATE " + table + " SET type = ? WHERE " + type_id + " = ? AND user_id = ?"
		return QueryAddUser(query2, args2)
	}
}

func AddGitHubUser(username string) (int, error) {

	var args []interface{}
	args = append(args, username)

	query1 := "SELECT user_id FROM users WHERE username = ?"

	user_id, err := QueryReturnOneNb(query1, args)

	if err != nil {
		query2 := "INSERT INTO users (username) VALUES ? "
		err1 := QueryAddUser(query2, args)
		if err1 != nil {
			return 0, err1
		}
		user_id, err = QueryReturnOneNb(query1, args)
	}
	return user_id, err
}

func AddGoogleUser(username string, email string) (int, error) {
	var args []interface{}
	args = append(args, username, email)

	query1 := "SELECT id FROM users WHERE username = ?"

	user_id, err := QueryReturnOneNb(query1, args)

	if err != nil {
		query2 := "INSERT INTO users (username, email) VALUES (?, ?)"
		err1 := QueryAddUser(query2, args)
		if err1 != nil {
			return 0, err1
		}
		user_id, err = QueryReturnOneNb(query1, args)
	}
	return user_id, err
}

func UpdatePostImagePathandType(imageType string, imageFilePath string, post_id int) error {
	var args []interface{}
	args = append(args, imageType, imageFilePath, post_id)

	query := "UPDATE posts SET img_type = ?, img_url = ? WHERE post_id = ?"

	return QueryAddUser(query, args)
}
