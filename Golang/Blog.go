package Blog_Ynov

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"text/template"
)

type Article struct {
	Id         int    `json:"Id,omitempty"`
	Image      string `json:"Image,omitempty"`
	Titre      string `json:"TItre,omitempty"`
	Sous_Titre string `json:"Sous-Titre,omitempty"`
}

func Lanch_Blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/blog.html"))
		templates.Execute(w, "")
	}
}

func SendDataBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetDataBlog()
		json.NewEncoder(w).Encode(AllData)
	}
}

func GetDataBlog() []Article {
	var articles []Article
	db, err := sql.Open("sqlite3", "./Blog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.Image, &article.Titre, &article.Sous_Titre)
		if err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return articles
}
