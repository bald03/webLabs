package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title          string
	Subtitle       string
	FeaturedPosts  []featuredPostData
	MostRecentPost []mostRecentPostData
}

type createPostRequest struct {
	Title           string `json:"postTitle"`
	Subtitle        string `json:"postShortDescr"`
	ImgModifier     string `json:"image_url"`
	Autor           string `json:"postAuthorName"`
	AutorImg        string `json:"author_url"`
	PublishDate     string `json:"postPublishDate"`
	Content         string `json:"postContent"`
	BigImageName    string `json:"postBigImageName"`
	SmallImageName  string `json:"postSmallImageName"`
	AuthorPhotoName string `json:"postAdminPhotoName"`
	BigImage        string `json:"postBigImage"`
	SmallImage      string `json:"postSmallImage"`
	AuthorPhoto     string `json:"postAdminPhoto"`
}

type postPage struct {
	Title       string `db:"title"`
	PostId      string `db:"post_id"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AutorImg    string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	Content     string `db:"content"`
}

type featuredPostData struct {
	Title       string `db:"title"`
	PostId      string `db:"post_id"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	Content     string `db:"content"`
}

type mostRecentPostData struct {
	Title       string `db:"title"`
	PostId      string `db:"post_id"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	Content     string `db:"content"`
}

type adminPage struct {
	Title string
}

type adminPostPage struct {
	Title string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}
		mostRecentPost, err := mostRecentPost(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 404) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		data := indexPage{
			Title:          "Escape",
			FeaturedPosts:  featuredPosts,
			MostRecentPost: mostRecentPost,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 403)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func admin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin-login.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 404) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		data := adminPage{
			Title: "admin-page",
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 403)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func adminPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin-post.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 404) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		data := adminPostPage{
			Title: "admin-post",
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 403)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest
		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Please, download all images", 401)
			log.Println(err.Error())
			return
		}

		authorImg, err := base64.StdEncoding.DecodeString(req.AuthorPhoto[strings.IndexByte(req.AuthorPhoto, ',')+1:])
		if err != nil {
			http.Error(w, "Please, download author image", 401)
			log.Println(err.Error())
			return
		}

		fileAuthor, err := os.Create("static/img/autor_avatars/" + req.AuthorPhotoName)
		if err != nil {
			http.Error(w, "Error, download to server", 500)
			log.Println(err.Error())
			return
		}
		defer fileAuthor.Close()

		_, err = fileAuthor.Write(authorImg)
		if err != nil {
			http.Error(w, "Error, download to server", 500)
			log.Println(err.Error())
			return
		}

		bigImg, err := base64.StdEncoding.DecodeString(req.BigImage[strings.IndexByte(req.BigImage, ',')+1:])
		if err != nil {
			http.Error(w, "Please, download article preview", 401)
			log.Println(err.Error())
			return
		}

		fileBig, err := os.Create("static/img/articles_img/" + req.BigImageName)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println(err.Error())
			return
		}
		defer fileBig.Close()

		_, err = fileBig.Write(bigImg)
		if err != nil {
			http.Error(w, "error", 500)
			log.Println(err.Error())
			return
		}

		smallImg, err := base64.StdEncoding.DecodeString(req.SmallImage[strings.IndexByte(req.SmallImage, ',')+1:])
		if err != nil {
			http.Error(w, "Please, download article preview", 401)
			log.Println(err.Error())
			return
		}

		fileSmall, err := os.Create("static/img/articles_img/" + req.SmallImageName)
		if err != nil {
			http.Error(w, "error", 500)
			log.Println(err.Error())
			return
		}
		defer fileSmall.Close()

		_, err = fileSmall.Write(smallImg)
		if err != nil {
			http.Error(w, "error", 500)
			log.Println(err.Error())
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "bd", 500)
			log.Println(err.Error())
			return
		}

	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content
		)
		VALUES
		(
			?,
			?,
			?,
			CONCAT('/static/img/autor_avatars/', ?),
			?,
			CONCAT('/static/img/articles_img/', ?),
			?
		)
	`

	_, err := db.Exec(query, req.Title, req.Subtitle, req.Autor, req.AuthorPhotoName, req.PublishDate, req.BigImageName, req.Content)
	// _, err := db.Exec(query, req.Title, req.Subtitle, req.Autor, req.PublishDate, req.Content)
	return err
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			post_id,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content
		FROM
				post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func mostRecentPost(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
			title,
			post_id,
			subtitle,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []mostRecentPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] // Получаем orderID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку orderID в число
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "post not found", 404)
				http.Error(w, "post not found", postID)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func postByID(db *sqlx.DB, postID int) (postPage, error) { //функция для получения информации об отдельном посте блога по его уникальному идентификатору postID.
	const query = `
		SELECT
			title,
			subtitle,
			content,
			image_url
		FROM
			` + "`post`" +
		`WHERE
			post_id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var order postPage

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&order, query, postID)
	if err != nil {
		return postPage{}, err
	}

	return order, nil
}
