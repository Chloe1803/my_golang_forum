package forum

import (
	q "forum/back/func/queries"
)

func CheckEmailAndUsername(email string, username string) (bool, error) {

	check := true

	nb, err := q.NbEmailAndUser(email, username)
	if err != nil {
		return false, err
	}

	if nb != 0 {
		check = false
	}

	return check, err
}
