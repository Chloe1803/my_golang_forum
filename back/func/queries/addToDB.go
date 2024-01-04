package forum

import "database/sql"

func QueryAddUser(query string, args []interface{}) error {

	db, err := GetDatabase()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)

	return err
}

func QueryAddPost(query string, args []interface{}) (sql.Result, error) {

	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	res, err1 := db.Exec(query, args...)

	return res, err1
}
