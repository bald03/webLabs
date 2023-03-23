package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title           string
	FeaturedPosts   []featuredPostData
	MostResentPosts []mostResentPostData
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
	Author      string
	AuthorImg   string
	PublishDate string
}

type mostResentPostData struct {
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
		Title:           "Escape.",
		FeaturedPosts:   featuredPosts(),
		MostResentPosts: mostResentPostData(),
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
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you’ve never been before.",
			ImgModifier: "static/img/articles_img/from-top-down.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
	}
}

func mostResentPostData() []mostResentPostData {
	return []mostResentPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you’ve never been before.",
			ImgModifier: "static/img/articles_img/from-top-down.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			ImgModifier: "static/img/articles_img/the-road-ahead-post.png",
			Author:      "Mat Vogels",
			AuthorImg:   "static/img/autor_avatars/mat_vogels.jpg",
			PublishDate: "September 25, 2015",
		},
	}
}
