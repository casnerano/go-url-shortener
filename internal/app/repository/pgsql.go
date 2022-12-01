package repository

import (
	"context"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type PgSQL struct {
}

func (m *PgSQL) Add(_ context.Context, url model.ShortURL) error {
	return nil
}

func (m *PgSQL) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	return nil, nil
}

func (m *PgSQL) DeleteByCode(_ context.Context, code string) error {
	return nil
}

func NewPgSQL() *PgSQL {
	return &PgSQL{}
}
