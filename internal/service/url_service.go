package service

import (
	"math/rand"
	"time"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *URLService) Create(original string) model.URL {
	code := generateCode()

	url := model.URL{
		Code:      code,
		Original:  original,
		Clicks:    0,
		CreatedAt: time.Now(),
	}

	s.repo.Save(url)
	return url
}

func (s *URLService) GetByCode(code string) (model.URL, bool) {
	return s.repo.FindByCode(code)
}
