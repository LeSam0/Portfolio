package Blog_Ynov

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func Create() {
	//connect to database
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	Database = db
	defer db.Close()
	// Create table portfolio
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS portfolio (id INTEGER PRIMARY KEY AUTOINCREMENT, ProjetName TEXT, ProjetType TEXT, ProjetImage TEXT)")
	if err != nil {
		panic(err)
	}
	log.Println("Table portfolio created successfully")
	// Create table entreprise
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS entreprise (id INTEGER PRIMARY KEY AUTOINCREMENT, EntrepriseName TEXT, EntrepriseImage TEXT, EntrepriseContract TEXT, EntrepiseLink TEXT, EntrepriseCreation TEXT, EntrepriseCreator TEXT, EntrepriseDirigeant TEXT, EntrepriseLocalisation TEXT, EntrepriseActivite TEXT, EntrepriseDescription TEXT, EntrepriseExperience TEXT)")
	if err != nil {
		panic(err)
	}
	log.Println("Table entreprise created successfully")
}

