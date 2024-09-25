package repository

import ("gin_news/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (int, error)
	GetStatusByID(id int) (int, error)
}

type Newslist interface {
	Create(news models.News) (int, error)
	GetAll() ([]models.News, error)
	GetByIdNews(id int) (models.News, error)
	DeleteNews(id int) error
	UpdateNews(id int, input models.UpdateNews) error
}



type Repository struct {
	Authorization
	Newslist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Newslist:      NewNewsPostgres(db),
	}
}