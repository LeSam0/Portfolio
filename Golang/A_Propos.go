package Blog_Ynov

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"text/template"
)

type Projet struct {
	Name  string
	Id    int
	Type  string
	Image string
}

func Lanch_A_Propos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/profil.html"))
		templates.Execute(w, "")
	}
}

func SendDataPortfolio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetDataPortfolio()
		json.NewEncoder(w).Encode(AllData)
	}
}

func GetDataPortfolio() []Projet {
	var projets []Projet
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM portfolio")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var projet Projet
		err = rows.Scan(&projet.Id, &projet.Name, &projet.Type, &projet.Image)
		if err != nil {
			panic(err)
		}
		projets = append(projets, projet)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return projets
}
