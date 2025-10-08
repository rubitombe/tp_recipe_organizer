package main

import (
	"html/template"
	"log"
	"net/http"
)

type Plates struct {
	Title       string
	Description string
	Category    string
	Difficulty  string
}

func Home(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("./index.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
