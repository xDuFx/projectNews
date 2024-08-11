package repository

import (
	"context"
	"maratproject/pkg/models"
	"maratproject/pkg/optional"
)

func (repo *PGRepo) GetNews(page, limit int) ([]models.News, error) {
	page = (page - 1) * limit
	rows, err := repo.pool.Query(context.Background(), `
	SELECT title, body, image, mark, reliz 
	FROM Newstable WHERE n > $2 
	ORDER BY n ASC
	LIMIT ($1);`, limit, page,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var datas []models.News
	for rows.Next() {
		var data models.News
		err := rows.Scan(
			&data.Title,
			&data.Body,
			&data.Image,
			&data.Mark,
			&data.Reliz,
		)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (repo *PGRepo) NewNews(item models.News) error {
	_, err := repo.pool.Exec(context.Background(), `
	INSERT INTO Newstable(title, body, image, mark, reliz)
    VALUES ($1, $2, $3, $4, $5);`,
		item.Title,
		item.Body,
		item.Image,
		item.Mark,
		item.Reliz,
	)

	return err
}

func (repo *PGRepo) DelNews(id int) error {
	_, err := repo.pool.Exec(context.Background(), `
	DELETE FROM Newstable
    WHERE id = ($1);`,
		id,
	)
	return err
}

func (repo *PGRepo) Authentication (data models.UserDataLogin) (bool, error) {
	row := repo.pool.QueryRow(context.Background(), `
	SELECT passhash 
	FROM UsersData
	WHERE login = $1`,
		data.Login,
	)
	data_check := models.UserDataLogin{}
	err := row.Scan(
		&data_check.Hashpass,
	)
	if err != nil {
		return false, err
	}
	if optional.CheckHash(data_check.Hashpass, data.Hashpass) {
		return true, nil
	}
	return false, nil
}

// func (repo *PGRepo) Register(data models.UserDataLogin) (bool, error) {
// 	var id int
// 	err := repo.pool.QueryRow(context.Background(), `
// 		`,
// 		data.Login,
// 		data.Login,
// 		111,
// 		data.Login,
// 		data.Login,
// 		data.Login,
// 		data.Login,
// 	).Scan(&id)

// 	if err != nil {
// 		return false, err
// 	}
// 	if id == 1 {
// 		return true, nil
// 	}
// 	return false, nil

// }
