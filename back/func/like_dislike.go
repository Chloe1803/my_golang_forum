package forum

import (
	"fmt"
	q "forum/back/func/queries"
)

func Like(table1 string, table2 string, col_type string, col_id int, user_id int) error {
	err := q.UpdateNbLike(table1, col_type, col_id, "likes", 1)

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = q.UpdateStatusLike(table2, col_type, col_id, user_id, 1)
	if err != nil {
		return err
	}
	return nil
}

func Dislike(table1 string, table2 string, col_type string, col_id int, user_id int) error {
	err := q.UpdateNbLike(table1, col_type, col_id, "dislikes", 1)
	if err != nil {
		return err
	}
	err = q.UpdateStatusLike(table2, col_type, col_id, user_id, -1)
	if err != nil {
		return err
	}
	return nil
}

func Unlike(table1 string, table2 string, col_type string, col_id int, user_id int) error {
	err := q.UpdateNbLike(table1, col_type, col_id, "likes", -1)
	if err != nil {
		return err
	}
	err = q.UpdateStatusLike(table2, col_type, col_id, user_id, 0)
	if err != nil {
		return err
	}

	return nil
}

func Undislike(table1 string, table2 string, col_type string, col_id int, user_id int) error {
	err := q.UpdateNbLike(table1, col_type, col_id, "dislikes", -1)
	if err != nil {
		return nil
	}
	err = q.UpdateStatusLike(table2, col_type, col_id, user_id, 0)
	if err != nil {
		return err
	}
	return nil
}
