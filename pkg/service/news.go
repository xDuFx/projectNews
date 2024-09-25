package service

import (
	"gin_news/pkg/models"
	"gin_news/pkg/repository"
)

type NewsService struct {
	repo     repository.Newslist
}


func NewNewsService(repo repository.Newslist) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) Create(item models.News) (int, error) {
	return s.repo.Create(item)
}

func (s *NewsService) GetAll() ([]models.News, error) {
	return s.repo.GetAll()
}

func (s *NewsService) GetByIdNews(itemId int) (models.News, error) {
	return s.repo.GetByIdNews(itemId)
}

func (s *NewsService) DeleteNews(itemId int) error {
	return s.repo.DeleteNews(itemId)
}

func (s *NewsService) UpdateNews(itemId int, input models.UpdateNews) error {
	return s.repo.UpdateNews(itemId, input)
}