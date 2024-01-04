package forum

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string = "back/database/forum.db"

func InitiateDatabase() {

	// Variable qui permet de se connecter à la base de données
	var db *sql.DB
	var err error

	//La fonction "os.Stat" vérifie si la base de données existe déjà,
	// si elle existe déjà, la fonction "InitiateDatabase" s'arrête car elle ne renvoie pas d'erreur
	if _, err = os.Stat(dbPath); err == nil {
		return
	}

	//sql.Open établie la connexion à la base de données et attribut la connexion à la variable db
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données :", err)
	}

	// La méthode .Exec() exécute des queries qui modifie les données de la DB
	// on va l'utiliser pour créer les différentes tables de la DB

	//table qui va stocker les utilisateurs
	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email    TEXT,
			username TEXT,
			password TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//cette table va servir à gérer les sessions/connexion des utilisateurs
	//Cette table à 2 contraintes qui sont définit sur les 2 dernières lignes
	_, err = db.Exec(`
	CREATE TABLE user_sessions (
		session_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER UNIQUE,
		FOREIGN KEY (user_id) REFERENCES users (id)
	)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//table qui va stocker les publications
	_, err = db.Exec(`
		CREATE TABLE posts (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			title TEXT,
			content TEXT,
			img_url TEXT DEFAULT '',
			img_type TEXT DEFAULT '',
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			comments INTEGER DEFAULT 0,
			user_id INTEGER REFERENCES users(user_id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//Table pour les commentaires

	_, err = db.Exec(`
		CREATE TABLE comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			content TEXT,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			user_id REFERENCES users (user_id),
			post_id REFERENCES posts (post_id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//Table qui définit les catégories
	_, err = db.Exec(`
	CREATE TABLE categories (
		category_id INTEGER PRIMARY KEY AUTOINCREMENT,
		category_name TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//Table relationnelle (many-to-many) qui fait le lien entre les likes, les posts et les utilisateurs
	_, err = db.Exec(`
		CREATE TABLE like_post (
			like_post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER REFERENCES post(post_id),
			user_id INTEGER REFERENCES user(user_id),
			type INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//table relationnelle entre les likes de commentaires, le commentaire et l'utilisateurs

	_, err = db.Exec(`
		CREATE TABLE like_comment (
			comment_post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			comment_id INTEGER REFERENCES comment(comment_id),
			user_id INTEGER REFERENCES user(user_id),
			type INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	//table relationnelle entre les posts et les catégories
	_, err = db.Exec(`
		CREATE TABLE post_category (
			post_category_id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER REFERENCES post(post_id),
			category_id INTEGER REFERENCES category(category_id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Ne pas oubliez de fermer la connexion
	db.Close()

	//Fonction qui va peupler la data base avec des données fictives
	PopulateDatabase()

}
