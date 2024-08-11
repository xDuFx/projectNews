package api

import (
	"maratproject/pkg/repository"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) FillEndpoints() {
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	api.r.Handle("/static", http.NotFoundHandler())
	api.r.Handle("/static/", http.StripPrefix("/static", fileServer))
	api.r.HandleFunc("/api/news", api.news)
	api.r.HandleFunc("/api/addnews/{token}", api.addNews)
	api.r.HandleFunc("/api/auth", func(w http.ResponseWriter, r *http.Request){api.login_page(w,"")})
	api.r.HandleFunc("/api/loggin_check", api.logins)
	api.r.HandleFunc("/api/del_news", api.delete_news).Queries("id", "{id:[0-9]+}")
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
