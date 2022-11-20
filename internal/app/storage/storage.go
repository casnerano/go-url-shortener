package storage

import "github.com/casnerano/go-url-shortener/internal/app/model"

type Storage interface {
	AddURL(url model.ShortURL) error
	GetURL(string) (*model.ShortURL, error)
}
