package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maratproject/pkg/models"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func (api *api) news(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pglm := mux.Vars(r)
		pg, ok := pglm["page"]
		if !ok {
			http.Error(w, "No parameter", http.StatusInternalServerError)
			return
		}
		lm, ok := pglm["limit"]
		if !ok {
			http.Error(w, "No parameter", http.StatusInternalServerError)
			return
		}
		page, err := strconv.Atoi(pg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		limit, err := strconv.Atoi(lm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := api.db.GetNews(page, limit)
		if err != nil {
			http.Error(w, err.Error()+"Нет данных", http.StatusInternalServerError)
			return
		}
		files := []string{
			"./static/html/photo.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (api *api) addNews(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://127.0.0.1:8090/api/news", http.StatusSeeOther)
	var News models.News
	switch r.Method {
	case http.MethodPost:
		filepath := make(chan string, 1)
		ctxTime, _ := context.WithTimeout(context.Background(), 5*time.Second) //спорное решение
		files := []string{                                                     //чтобы долго не висло
			"./static/html/photo.html",
		}
		templ(w, files, "")
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка парсинга формы", 500)
		}
		go func() {
			file, handler, err := r.FormFile("file")
			if err.Error() == "http: no such file" {
				filepath <- "Не было картинки"
				return
			}
			if err != nil && err.Error() != "http: no such file" {
				fmt.Println("Ошибка при получении файла:", err)
				return
			}
			defer file.Close()
			// Создание нового файла на сервере
			f, err := os.OpenFile("./static/image/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("Ошибка при создании файла:", err)
				return
			}
			defer f.Close()

			// Копирование содержимого файла в созданный файл на сервере
			_, err = io.Copy(f, file)
			if err != nil {
				fmt.Println("Ошибка при копировании файла:", err)
				return
			}
			filepath <- "/static/image/" + handler.Filename
		}()

		var path string
		var pathimage models.PathFromServer
		select {
		case <-ctxTime.Done():
			fmt.Println("Не смог принять файл")
			return
		case path = <-filepath:
			pathimage.Path = path
			err := json.NewEncoder(w).Encode(pathimage)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			News.Image = path
			News.Title = r.FormValue("title")
			News.Body = r.FormValue("body")
			News.Mark = r.FormValue("mark")
			News.Reliz = r.FormValue("reliz")
			err = api.db.NewNews(News)
			if err != nil {
				fmt.Println("Ошибка при отправке новости в БД:", err)
				return
			}
		}
	}

}

func (api *api) delete_news(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		book_id, ok := vars["id"]
		if !ok {
			http.Error(w, "No parameter", http.StatusInternalServerError)
			return
		}
		id, err := strconv.Atoi(book_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.DelNews(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// err = json.NewDecoder(r.Body).Decode(&id)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
	}
}
func (api *api) login_page(w http.ResponseWriter, errorlog string) {
	files := []string{"./static/html/login.html"}
	templ(w, files, errorlog)
}

func (api *api) logins(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data_user models.UserDataLogin
		data_user.Login = r.FormValue("login")
		data_user.Hashpass = r.FormValue("password")
		if data_user.Login == "" || data_user.Hashpass == "" {
			api.login_page(w, "Необходимо указать логин и пароль")
			return
		}
		check_flag, err := api.db.Authentication(data_user)
		if err != nil && err.Error() != "no rows in result set" {
			log.Print(err.Error(), http.StatusInternalServerError)
			return
		}
		if !check_flag {
			api.login_page(w, "Неверный пароль")
			return
		}
		http.Redirect(w, r, "/api/news", http.StatusSeeOther)
	}
}

func templ(w http.ResponseWriter, files []string, message string) {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	type answer struct {
		Message string
	}
	data := answer{message}
	err = ts.ExecuteTemplate(w, "login", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
