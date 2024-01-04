package forum_test

import (
	q "forum/back/func/queries"
	"testing"

	"github.com/mattn/go-sqlite3"
)

// func TestAddUser(t *testing.T) {

// 	_ = sqlite3.SQLiteDriver{}
// 	// Testez le cas où l'e-mail et le nom d'utilisateur sont valides
// 	validEmail := "test@example.com"
// 	validUsername := "testuser"
// 	result, err := f.CheckEmailAndUsername(validEmail, validUsername)
// 	if err != nil {
// 		t.Errorf("Erreur inattendue : %v", err)
// 	}
// 	if !result {
// 		t.Error("Attendu : vrai, Obtenu : faux")
// 	}

// 	// Testez le cas où l'e-mail ou le nom d'utilisateur sont déjà utilisés
// 	// (assurez-vous d'adapter cela à votre logique)
// 	usedEmail := "sophie@sophie.com"
// 	usedUsername := "sophie"
// 	result, err = f.CheckEmailAndUsername(usedEmail, usedUsername)
// 	if err != nil {
// 		t.Errorf("Erreur inattendue : %v", err)
// 	}
// 	if result {
// 		t.Error("Attendu : faux, Obtenu : vrai")
// 	}
// }

func TestAddNewUser(t *testing.T) {
	// Initialisez le pilote SQLite3
	_ = sqlite3.SQLiteDriver{}

	// Utilisez la fonction depuis le package forum/back/func
	email := "nouvellemail@email.com"
	username := "nouvelutilisateur"
	password := "motdepasse"

	err := q.AddNewUser(email, username, password)

	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}

}
