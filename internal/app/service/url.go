package service

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("url adding error: %w", err)
	}

	return shortURLModel, nil
}

func (urlService *URL) CreateBatch(request []*model.ShortURLBatchRequest, uuid string) ([]*model.ShortURLBatchResponse, error) {
	response := make([]*model.ShortURLBatchResponse, 0, len(request))
	batchShortURL := make([]*model.ShortURL, 0, len(request))

	for _, requestBatchItem := range request {
		shortCode := urlService.hash.Generate(requestBatchItem.OriginalURL)
		shortURLModel := model.NewShortURL(shortCode, requestBatchItem.OriginalURL)
		shortURLModel.UserUUID = uuid

		batchShortURL = append(batchShortURL, shortURLModel)

		response = append(
			response,
			&model.ShortURLBatchResponse{
				CorrelationID: requestBatchItem.CorrelationID,
				ShortURL:      shortURLModel.Code,
			},
		)
	}

	err := urlService.rep.AddBatch(context.TODO(), batchShortURL)
	if err != nil {
		response = response[:0]
	}

	return response, err
}

func (urlService *URL) GetByCode(shortCode string) (*model.ShortURL, error) {
	return urlService.rep.GetByCode(context.TODO(), shortCode)
}

func (urlService *URL) GetByUserUUIDAndOriginal(uuid string, original string) (*model.ShortURL, error) {
	return urlService.rep.GetByUserUUIDAndOriginal(context.TODO(), uuid, original)
}

func (urlService *URL) FindByUserUUID(uuid string) ([]*model.ShortURL, error) {
	return urlService.rep.FindByUserUUID(context.TODO(), uuid)
}
