package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title          string
	FeaturedPosts  []featuredPostData
	MostRecentPost []mostRecentPostData
}

type postPage struct {
	Title    string
	Subtitle string
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

type mostRecentPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}
	data := indexPage{
		Title:          "Escape",
		FeaturedPosts:  featuredPosts(),
		MostRecentPost: mostRecentPost(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/from-top-down.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
	}
}

func mostRecentPost() []mostRecentPostData {
	return []mostRecentPostData{
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			ImgModifier: "static/img/articles_img/still_standing_tall_img.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/william_wong.jpg",
			PublishDate: "September 25, 2015",
		},
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}
	data := postPage{
		Title:    "The Road Ahead",
		Subtitle: "The road ahead might be paved - it might not be.",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
