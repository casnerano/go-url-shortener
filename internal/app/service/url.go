package service

import (
	"context"
	"errors"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
)

type URL struct {
	rep  repository.URLRepository
	hash hasher.Hash
}

func NewURL(rep repository.URLRepository, hash hasher.Hash) *URL {
	return &URL{rep, hash}
}

func (urlService *URL) Create(urlOriginal string, uuid string) (*model.ShortURL, error) {
	shortCode := urlService.hash.Generate(urlOriginal)
	shortURLModel := model.NewShortURL(shortCode, urlOriginal)
	shortURLModel.UserUUID = uuid

	err := urlService.rep.Add(context.TODO(), shortURLModel)
	if err != nil {
		return nil, errors.New("url adding error")
	}

	return shortURLModel, nil
}

func (urlService *URL) GetByCode(shortCode string) (*model.ShortURL, error) {
	shortURLModel, err := urlService.rep.GetByCode(context.TODO(), shortCode)
	if err != nil {
		return nil, err
	}

	return shortURLModel, nil
}

func (urlService *URL) FindByUserUUID(uuid string) ([]*model.ShortURL, error) {
	return urlService.rep.FindByUserUUID(context.TODO(), uuid)
}
