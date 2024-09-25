package repository

import (
	"fmt"
	"gin_news/pkg/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (login, passhash) values ($1, $2) RETURNING id", usersTable)

	row := tx.QueryRow(createItemQuery, user.Username, user.Password)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createQuery := fmt.Sprintf("INSERT INTO %s (email, username) values ($1, $2)", personaldata)
	_, err = tx.Exec(createQuery, user.Email, user.Name)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *AuthPostgres) GetUser(username, password string) (id int, err error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND passhash=$2", usersTable)
	err = r.db.Get(&id, query, username, password)

	return id, err
}

func (r *AuthPostgres) GetStatusByID(id int) (status int, err error) {
	query := fmt.Sprintf("SELECT status FROM %s WHERE id = $1", usersTable)
	err = r.db.Get(&status, query, id)
	return status, err
}
