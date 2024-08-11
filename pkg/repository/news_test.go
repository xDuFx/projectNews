package repository

import (
	"log"
	"maratproject/pkg/models"
	"testing"
)

const connStr = "postgres://postgres:123@localhost:5432/postgres"

func Test_NewsCRUD(t *testing.T) {
	test_data := models.UserDataLogin{Login: "asd", Hashpass: "asd"}

	db, err := New("postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Register(test_data)

	if err != nil {
		log.Fatal(err.Error())
	}

	
}
