package Blog_Ynov

import (
	"net/http"
	"text/template"
)

func Lanch_Blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/blog.html"))
		templates.Execute(w, "")
	}
}
