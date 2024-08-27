package Blog_Ynov

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"text/template"
)

type Entreprise struct {
	Id           int    `json:"Id,omitempty"`
	Name         string `json:"Name,omitempty"`
	Image        string `json:"Image,omitempty"`
	Contract     string `json:"Contract,omitempty"`
	Lien         string `json:"Lien,omitempty"`
	Creation     string `json:"Creation,omitempty"`
	Creator      string `json:"Creator,omitempty"`
	Dirigeant    string `json:"Dirigeant,omitempty"`
	Localisation string `json:"Localisation,omitempty"`
	Activite     string `json:"Activite,omitempty"`
	Description  string `json:"Description,omitempty"`
	Experience   string `json:"Experience,omitempty"`
}

func Lanch_Mes_Experience_Pro(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/entreprise.html"))
		templates.Execute(w, "")
	}
}

func SendData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetData()
		json.NewEncoder(w).Encode(AllData)
	}
}

func GetData() []Entreprise {
	var entreprises []Entreprise
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM entreprise")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var entreprise Entreprise
		err = rows.Scan(&entreprise.Id, &entreprise.Name, &entreprise.Image, &entreprise.Contract, &entreprise.Lien, &entreprise.Creation, &entreprise.Creator, &entreprise.Dirigeant, &entreprise.Localisation, &entreprise.Activite, &entreprise.Description, &entreprise.Experience)
		if err != nil {
			panic(err)
		}
		entreprises = append(entreprises, entreprise)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return entreprises
}
