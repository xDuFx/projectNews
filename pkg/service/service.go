package service

import (
	"gin_news/pkg/models"
	"gin_news/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckAccess (id int) (bool, error)
}

type Newslist interface {
	Create(news models.News) (int, error)
	GetAll() ([]models.News, error)
	GetByIdNews(id int) (models.News, error)
	DeleteNews(id int) error
	UpdateNews(id int, input models.UpdateNews) error
}

type Service struct {
	Authorization
	Newslist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Newslist:      NewNewsService(repos.Newslist),
	}
}