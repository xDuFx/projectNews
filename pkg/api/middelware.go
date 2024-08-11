package api

import (
	"log"
	"maratproject/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *api) MidCheck(check_handle func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		lg, ok := vars["tokken"]
		if !ok || lg == "" {
			http.Redirect(w, r, "/api/auth", http.StatusSeeOther)
			return
		}
		data := models.UserDataLogin{Login: "", Hashpass: lg}
		if flag, err := api.db.Authentication(data); flag{
			if err != nil {
				log.Printf("Ошибка аутенфикации: %s", err.Error())
				http.Redirect(w, r, "/api/auth", http.StatusSeeOther)
				return
			}
			http.Redirect(w, r, "/api/auth", http.StatusSeeOther)
			return
		}

		check_handle(w,r)
	}
}