package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPageDate struct {
	Title    string
	Subtitle string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return
	}

	data := indexPageDate{
		Title:    "Blog",
		Subtitle: "Best blog",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return
	}
}
