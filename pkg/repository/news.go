package repository

import (
	"fmt"
	"gin_news/pkg/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

//	func (r *NewsPostgres) GetNews(page, limit int) ([]models.News, error) {
//		page = (page - 1) * limit
//		query := fmt.Sprintf(`SELECT title, body, image, mark, reliz FROM %s WHERE n > %s ORDER BY n ASC LIMIT (%s);`, NewsTable, page, limit)
//		rows, err := r.db.Get( `
//		SELECT title, body, image, mark, reliz
//		FROM Newstable WHERE n > $2
//		ORDER BY n ASC
//		LIMIT ($1);`, limit, page,
//		)
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//		var datas []models.News
//		for rows.Next() {
//			var data models.News
//			err := rows.Scan(
//				&data.Title,
//				&data.Body,
//				&data.Image,
//				&data.Mark,
//				&data.Reliz,
//			)
//			if err != nil {
//				return nil, err
//			}
//			datas = append(datas, data)
//		}
//		return datas, nil
//	}
func (r *NewsPostgres) Create(item models.News) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s(title, body, image, mark, reliz) VALUES ($1, $2, $3, $4, $5) RETURNING id;`, NewsTable)
	var id int
	row := r.db.QueryRow(query, item.Title, item.Body, item.Image, item.Mark, item.Reliz)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *NewsPostgres) GetAll() (news []models.News, err error) {
	query := fmt.Sprintf(`SELECT id, title, body, image, mark, reliz FROM %s;`, NewsTable)
	err = r.db.Select(&news, query)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *NewsPostgres) GetByIdNews(id int) (models.News, error) {
	query := fmt.Sprintf(`SELECT id, title, body, image, mark, reliz FROM %s WHERE id = $1;`, NewsTable)
	news := models.News{}
	err := r.db.Get(&news, query, id)
	if err != nil {
		return news, err
	}
	return news, nil
}

func (r *NewsPostgres) DeleteNews(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = ($1);`, NewsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *NewsPostgres) UpdateNews(id int, input models.UpdateNews) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argId))
		args = append(args, *input.Body)
		argId++
	}

	if input.Image != nil {
		setValues = append(setValues, fmt.Sprintf("image=$%d", argId))
		args = append(args, *input.Image)
		argId++
	}

	if input.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argId))
		args = append(args, *input.Mark)
		argId++
	}

	if input.Reliz != nil {
		setValues = append(setValues, fmt.Sprintf("reliz=$%d", argId))
		args = append(args, *input.Reliz)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s 
									WHERE id = $%d`,
		NewsTable, setQuery, argId)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
