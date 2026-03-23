package repository

import (
	"url-shortener/internal/model"
)

type URLRepository struct {
	data map[string]model.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		data: make(map[string]model.URL),
	}
}

func (r *URLRepository) Save(url model.URL) {
	r.data[url.Code] = url
}

func (r *URLRepository) FindByCode(code string) (model.URL, bool) {
	url, exists := r.data[code]
	return url, exists
}

func (r *URLRepository) FindAll() []model.URL {
	var urls []model.URL
	for _, u := range r.data {
		urls = append(urls, u)
	}
	return urls
}
