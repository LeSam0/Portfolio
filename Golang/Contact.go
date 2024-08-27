package Blog_Ynov

import (
	"net/http"
	"text/template"
)

func Lanch_Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/contact.html"))
		templates.Execute(w, "")
	}
}

func Lanch_Contact_Recu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/message-reçu.html"))
		templates.Execute(w, "")
	}
}

func Lanch_Contact_Error(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/message-error.html"))
		templates.Execute(w, "")
	}
}