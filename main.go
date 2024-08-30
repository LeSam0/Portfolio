package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Portfolio struct {
	Name  string `json:"Name,omitempty"`
	Id    int    `json:"Id,omitempty"`
	Type  string `json:"Type,omitempty"`
	Image string `json:"Image,omitempty"`
}

type Article struct {
	Id      int    `json:"Id,omitempty"`
	Name    string `json:"Name,omitempty"`
	Image   string `json:"Image,omitempty"`
	Date    string `json:"Date,omitempty"`
	Context string `json:"Context,omitempty"`
	Text    string `json:"Text,omitempty"`
}

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

type Projet struct {
	Name  string `json:"Name,omitempty"`
	Id    int    `json:"Id,omitempty"`
	Type  string `json:"Type,omitempty"`
	Image string `json:"Image,omitempty"`
}

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Lanch_A_Propos)
	http.HandleFunc("/pas_trouve", Lanch_404)
	http.HandleFunc("/a_propos", Lanch_A_Propos)
	http.HandleFunc("/a_propos/data", SendDataPortfolio)
	http.HandleFunc("/a_propos/projet", Lanch_Projet)
	http.HandleFunc("/a_propos/projet/data", SendDataProjet)
	http.HandleFunc("/contact", Lanch_Contact)
	http.HandleFunc("/contact/recu", Lanch_Contact_Recu)
	http.HandleFunc("/contact/error", Lanch_Contact_Error)
	http.HandleFunc("/blog", Lanch_Blog)
	http.HandleFunc("/blog/data", SendDataBlog)
	http.HandleFunc("/blog/article", Lanch_Article)
	http.HandleFunc("/blog/article/data", SendDataArticle)
	http.HandleFunc("/mes_experience_professionel", Lanch_Mes_Experience_Pro)
	http.HandleFunc("/mes_experience_professionel/data", SendData)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Lanch_404(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("Page/404.html"))
	templates.Execute(w, "")
}

func Lanch_Projet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/projet.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func SendDataProjet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("numero")
		AllData := GetDataProjet(id)
		json.NewEncoder(w).Encode(AllData)
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_A_Propos(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/a_propos" && r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/profil.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func SendDataPortfolio(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetDataPortfolio()
		json.NewEncoder(w).Encode(AllData)
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Article(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var templates = template.Must(template.ParseFiles("Page/article.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func SendDataArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("numero")
		AllData := GetDataArticle(id)
		json.NewEncoder(w).Encode(AllData)
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Blog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/blog.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func SendDataBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetDataBlog()
		json.NewEncoder(w).Encode(AllData)
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/contact.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Contact_Recu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/message-reçu.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Contact_Error(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/message-error.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func Lanch_Mes_Experience_Pro(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var templates = template.Must(template.ParseFiles("Page/entreprise.html"))
		templates.Execute(w, "")
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func SendData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		AllData := GetData()
		json.NewEncoder(w).Encode(AllData)
	} else {
		http.Redirect(w, r, "/pas_trouve", http.StatusSeeOther)
	}
}

func GetData() []Entreprise {
	entreprises := []Entreprise{
		Entreprise{
			Id:           1,
			Name:         "Shinken Solutions",
			Image:        "shinken.png",
			Contract:     "Stage de 6 semaines",
			Lien:         "https://www.shinken-enterprise.com/",
			Creation:     "2013",
			Creator:      "Jean Gabès et Jean-Paul Harnisch",
			Dirigeant:    "Anis Zouaoui",
			Localisation: "33370 Artigues-près-Bordeaux",
			Activite:     "Surveillance système et réseau",
			Description:  "Shinken est une application de surveillance système et réseau. Elle surveille les hôtes et services spécifiés, alertant lorsque les systèmes vont mal et quand ils vont mieux. Elle a pour but d'apporter une supervision distribuée et hautement disponible, facile à mettre en place.<br><br>Shinken Solutions a été créée en 2013 par Jean Gabès et Jean-Paul Harnisch. L'entreprise développe une solution de supervision des différents systèmes informatiques d'une organisation. La solution Shinken permet de superviser la disponibilité des systèmes informatiques en continu et de détecter en temps réel les problèmes, tout en aidant à leurs résolutions avant qu'ils n'impactent les utilisateurs.<br><br>En 2023, Shinken Solutions, l'entreprise qui développe Shinken, a été acquise par Adservio Group. Adservio Group, représentée par AD Holding et son gérant Anis ZOUAOUI, est désormais le président de Shinken Solutions.",
			Experience:   "Mon expérience dans cette entreprise est du a un stage de 6 semaines, de juin 2024 a juillet 2024, pour valider ma 2e année de licence informatique au sein d'Ynov. <br><br>Durant mon stage de 6 semaines, j'ai réalisé plusieurs nouvelles fonctionnalités pour le site web, fonctionnalités qui ont été ajouter à la nouvelle mis à jour. Ce qui m'a permis de contribuer à l'amélioration de la solution et de comprendre le fonctionnement d'une entreprise de développement web, qui fonctionne en méthode agile.<br><br>De plus, l'ambiance au sein de l'entreprise est excellente, j'ai été très bien accueilli et vite intégrer. J'ai travaillé en collaboration avec plusieurs autre employé, ce qui m'a permis de me sentir utile dans l'entreprise. ",
		},
		Entreprise{
			Id:           2,
			Name:         "Centre des finances public",
			Image:        "DGFiP.png",
			Contract:     "Vacataire de 6 semaines",
			Lien:         "http://economie.gouv.fr",
			Creation:     "1518",
			Creator:      "Jacques de Beaune",
			Dirigeant:    "Bruno Le Maire",
			Localisation: "33000 Melun",
			Activite:     "Service Amendes",
			Description:  "Le ministère des finances est l'institution gouvernementale chargée de gérer les finances publiques, il s'occupe de collecter les impôts, de payer les dépenses de l'État, d'établir le budget annuel et de s'assurer de son équilibre. De définir la politique économique, il joue un rôle central dans la mise en place des orientations économiques d'un pays, en régulant l'activité économique, en contrôlant l'inflation et en favorisant la croissance. De surveiller les marchés financiers, il veille au bon fonctionnement des marchés financiers et assure la stabilité du système bancaire.<br><br>En résumé, le ministère des finances est le gardien des finances publiques et le moteur de la politique économique d'un pays.<br><br>Un centre des amendes est un organisme qui centralise le traitement des infractions et délits pour lesquelles une amende est due. Il s'occupe de réceptionner les infractions relevées par les forces de l'ordre (police, gendarmerie) ou par des systèmes automatisés (radars, caméras). De notifier les amendes aux contrevenants. De gérer les paiements des amendes. Et informer les contrevenants sur leurs droits et sur les modalités de paiement.",
			Experience:   "Mon expérience au sein de cette administration publique dura 6 semaines, de Mai 2023 à Juin 2023, en contrat vacataire pour un job d'été.<br><br>Durant ces 6 semaines, j'ai réalisé plusieurs tâches afin d'aider les fonctionnaires du centre des amendes de Melun. J'ai pu me rentrer utile et comprendre le fonctionnement d'une administration publique. <br><br>Pour ce qui s'agit de l'ambiance au sein de ce centre, tout le monde s'entendait parfaitement, une 'réunion' autour d'un café était organisée tous les matins afin de discuter et de garder le lien entre les fonctionnaires.",
		},
	}

	return entreprises
}

func GetDataPortfolio() []Portfolio {
	projets := []Portfolio{
		Portfolio{
			Name:  "LifeManager",
			Id:    1,
			Type:  "Application",
			Image: "LifeManager.png",
		},
	}
	return projets
}

func GetDataBlog() []Article {
	articles := []Article{
		Article{
			Id:      1,
			Name:    "Certification Web",
			Date:    "14/06/2024",
			Image:   "certification.png",
			Context: "Validation de la certification 'Intégrer les règles et le vocabulaire assurance qualité web dans sa pratique professionnelle', certification passé dans le cadre de mes études au sein d'ynov sur la plateforme Opquast.",
			Text:    "Durant ma deuxième année, j'ai eu l'occasion de passer un certification web sur la plateforme Opquast. Après plusieurs mois d'entrainement, j'ai réaliser l'examen final et l'ai obtenue le 14 juin 2024 avec une note de 750/1000. L'examen final était totalement différent des entrainement et bien plus tricky.<br><br>Cette certification ne sers pas a grand chose pour le moment puisque je souhaite m'orianter en cybersécurité, mais elle reste une bonne expérience afin de pouvoir en passer d'autre en cybersécurité.",
		},
		Article{
			Id:      2,
			Name:    "Bilan de stage",
			Date:    "26/07/2024",
			Image:   "stage.png",
			Context: "Bilan de mon stage, de 6 semaines, de fin de 2e années d'étude en formation informatique, au sein de l'entreprise Shinken Solution.<br><br>Entreprise créatrice de la solution shinken, Solution de surveillance de systèmes et réseaux en interface web.",
			Text:    "Durant mon stage au sein de l'entreprise, j'ai pu contribuer à l'amélioration de la solution, en réalisant plusieurs nouvelles fonctionnalités pour le site, fonctionnalités qui ont été ensuite ajouter à la nouvelle mis à jour. J'ai donc pu mettre a profit mes connaissance vu en cours, et changer ma méthode de travail pour fonctionner en méthode agile dans une teams de développement.<br><br>Ce stage été passionnant et l'ambience au sein de l'entreprise est incroyable, j'ai été très bien accueilli et vite intégrer.J'ai travaillé en collaboration avec plusieurs autre employés, ce qui m'a permis de me sentir utile et chaque personne au sein de l'entreprise a beaucoup a m'apprendre.<br><br>Malgré tous c'est bon point, certaine chose pourrait être améliorer, comme l'utilisation de Teams à la place de Skype ou encore ne pas utiliser un iframe pour afficher le html.",
		},
	}
	return articles
}

func GetDataArticle(id string) Article {
	articles := []Article{
		Article{
			Id:      1,
			Name:    "Certification Web",
			Date:    "14/06/2024",
			Image:   "certification.png",
			Context: "Validation de la certification 'Intégrer les règles et le vocabulaire assurance qualité web dans sa pratique professionnelle', certification passé dans le cadre de mes études au sein d'ynov sur la plateforme Opquast.",
			Text:    "Durant ma deuxième année, j'ai eu l'occasion de passer un certification web sur la plateforme Opquast. Après plusieurs mois d'entrainement, j'ai réaliser l'examen final et l'ai obtenue le 14 juin 2024 avec une note de 750/1000. L'examen final était totalement différent des entrainement et bien plus tricky.<br><br>Cette certification ne sers pas a grand chose pour le moment puisque je souhaite m'orianter en cybersécurité, mais elle reste une bonne expérience afin de pouvoir en passer d'autre en cybersécurité.",
		},
		Article{
			Id:      2,
			Name:    "Bilan de stage",
			Date:    "26/07/2024",
			Image:   "stage.png",
			Context: "Bilan de mon stage, de 6 semaines, de fin de 2e années d'étude en formation informatique, au sein de l'entreprise Shinken Solution.<br><br>Entreprise créatrice de la solution shinken, Solution de surveillance de systèmes et réseaux en interface web.",
			Text:    "Durant mon stage au sein de l'entreprise, j'ai pu contribuer à l'amélioration de la solution, en réalisant plusieurs nouvelles fonctionnalités pour le site, fonctionnalités qui ont été ensuite ajouter à la nouvelle mis à jour. J'ai donc pu mettre a profit mes connaissance vu en cours, et changer ma méthode de travail pour fonctionner en méthode agile dans une teams de développement.<br><br>Ce stage été passionnant et l'ambience au sein de l'entreprise est incroyable, j'ai été très bien accueilli et vite intégrer.J'ai travaillé en collaboration avec plusieurs autre employés, ce qui m'a permis de me sentir utile et chaque personne au sein de l'entreprise a beaucoup a m'apprendre.<br><br>Malgré tous c'est bon point, certaine chose pourrait être améliorer, comme l'utilisation de Teams à la place de Skype ou encore ne pas utiliser un iframe pour afficher le html.",
		},
	}
	ids, _ := strconv.Atoi(id)
	return articles[ids-1]
}

func GetDataProjet(id string) Projet {
	projet := []Projet{
		Projet{
			Name:  "LifeManager",
			Id:    1,
			Type:  "Application",
			Image: "LifeManager.png",
		},
	}
	ids, _ := strconv.Atoi(id)
	return projet[ids]
}

// func GetData() []Entreprise {
// 	var entreprises []Entreprise
// 	db, err := sql.Open("sqlite3", "./Blog.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM entreprise")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var entreprise Entreprise
// 		err = rows.Scan(&entreprise.Id, &entreprise.Name, &entreprise.Image, &entreprise.Contract, &entreprise.Lien, &entreprise.Creation, &entreprise.Creator, &entreprise.Dirigeant, &entreprise.Localisation, &entreprise.Activite, &entreprise.Description, &entreprise.Experience)
// 		if err != nil {
// 			panic(err)
// 		}
// 		entreprises = append(entreprises, entreprise)
// 	}
// 	if err = rows.Err(); err != nil {
// 		panic(err)
// 	}
// 	return entreprises
// }

// func GetDataPortfolio() []Portfolio {
// 	var projets []Portfolio
// 	db, err := sql.Open("sqlite3", "./Blog.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM portfolio")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var projet Portfolio
// 		err = rows.Scan(&projet.Id, &projet.Name, &projet.Type, &projet.Image)
// 		if err != nil {
// 			panic(err)
// 		}
// 		projets = append(projets, projet)
// 	}
// 	if err = rows.Err(); err != nil {
// 		panic(err)
// 	}
// 	return projets
// }

// func GetDataBlog() []Article {
// 	var articles []Article
// 	db, err := sql.Open("sqlite3", "./Blog.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM articles")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var article Article
// 		err = rows.Scan(&article.Id, &article.Name, &article.Image, &article.Date, &article.Context, &article.Text)
// 		if err != nil {
// 			panic(err)
// 		}
// 		articles = append(articles, article)
// 	}
// 	if err = rows.Err(); err != nil {
// 		panic(err)
// 	}
// 	return articles
// }

// func GetDataArticle(id string) Article {
// 	var article Article
// 	db, err := sql.Open("sqlite3", "./Blog.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM articles WHERE id = ?", id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err = rows.Scan(&article.Id, &article.Name, &article.Image, &article.Date, &article.Context, &article.Text)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	if err = rows.Err(); err != nil {
// 		panic(err)
// 	}
// 	return article
// }

// func GetDataProjet(id string) Projet {
// 	var projet Projet
// 	db, err := sql.Open("sqlite3", "./Blog.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM portfolio WHERE id = ?", id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err = rows.Scan(&projet.Id, &projet.Name, &projet.Image)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	if err = rows.Err(); err != nil {
// 		panic(err)
// 	}
// 	return projet
// }
