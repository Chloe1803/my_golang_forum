package forum

import (
	"fmt"
	q "forum/back/func/queries"
)

func Convert_args(args []interface{})(string, string, string, int, int){
	table1, _ := args[0].(string)
	table2, _ := args[1].(string)
	col_type, _ := args[2].(string)
	col_id, _ := args[3].(int)
	user_id, _ := args[4].(int)
	
	return table1, table2, col_type, col_id, user_id
}

func Like(args []interface{}) error {
	table1, table2, col_type, col_id, user_id := Convert_args(args)

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

func Dislike(args []interface{}) error {

	table1, table2, col_type, col_id, user_id := Convert_args(args)
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

func Unlike(args []interface{}) error {
	table1, table2, col_type, col_id, user_id := Convert_args(args)
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

func Undislike(args []interface{}) error {
	table1, table2, col_type, col_id, user_id := Convert_args(args)

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
