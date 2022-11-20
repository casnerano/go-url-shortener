package storage

import "github.com/casnerano/go-url-shortener/internal/app/model"

type Storage interface {
	AddUrl(url model.ShortUrl) error
	GetUrl(string) (*model.ShortUrl, error)
}
