package Blog_Ynov

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func SuppToDB(id string) {
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
}

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
	// Create table article
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS articles (id INTEGER PRIMARY KEY AUTOINCREMENT, ArticleName TEXT, ArticleImage TEXT, ArticleDate TEXT, ArticleContext TEXT, ArticleText TEXT)")
	if err != nil {
		panic(err)
	}
	log.Println("Table article created successfully")
}

func CreateCategorieCourse() {
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO articles (ArticleName, ArticleImage, ArticleDate, ArticleContext, ArticleText) VALUES (?, ?, ?, ?, ?)", "Bilan de stage", "stage.png", "26/07/2024", "Bilan de mon stage, de 6 semaines, de fin de 2e année d'étude en formation informatique, au sein de l'entreprise Shinken Solution.<br><br>Entreprise créatrice de la solution shinken, Solution de surveillance de systèmes et réseaux en interface web.", "Durant mon stage au sein de l'entreprise, j'ai pu contribuer à l'amélioration de la solution, en réalisant plusieurs nouvelles fonctionnalités pour le site, fonctionnalités qui ont été ensuite ajoutées à la nouvelle mis à jour. J'ai donc pu mettre à profit mes connaissances vu en cours, et changer ma méthode de travail pour fonctionner en méthode agile dans une teams de développement.<br><br>Ce stage était passionnant et l'ambiance au sein de l'entreprise est incroyable, j'ai été très bien accueilli et vite intégrer.J'ai travaillé en collaboration avec plusieurs autre employés, ce qui m'a permis de me sentir utile et chaque personne au sein de l'entreprise a beaucoup à m'apprendre.<br><br>Malgré tous, c'est bon point, certaines choses pourraient être améliorées, comme l'utilisation de Teams à la place de Skype ou encore ne pas utiliser un iframe pour afficher le html.")
	if err != nil {
		panic(err)
	}
	log.Println("Categorie course created successfully")
}