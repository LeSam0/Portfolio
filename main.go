package main

import (
	Page "Blog_Ynov/Golang"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Page.Lanch_A_Propos)
	http.HandleFunc("/a_propos/data", Page.SendDataPortfolio)
	http.HandleFunc("/contact", Page.Lanch_Contact)
	http.HandleFunc("/contact/recu", Page.Lanch_Contact_Recu)
	http.HandleFunc("/contact/error", Page.Lanch_Contact_Error)
	http.HandleFunc("/blog", Page.Lanch_Blog)
	http.HandleFunc("/blog/article", Page.Lanch_Article)
	http.HandleFunc("/mes_experience_professionel", Page.Lanch_Mes_Experience_Pro)
	http.HandleFunc("/mes_experience_professionel/data", Page.SendData)
	Page.Create()
	log.Fatal(http.ListenAndServe(":8000", nil))
}