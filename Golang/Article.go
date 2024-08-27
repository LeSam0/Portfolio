package Blog_Ynov

import (
	"net/http"
	"text/template"
)

func Lanch_Article(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var templates = template.Must(template.ParseFiles("Page/article.html"))
		templates.Execute(w, "")
	}
}
