package sqlstore

import (
	"context"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(ctx context.Context, url model.ShortURL) error {
	_, err := rep.store.db.Exec(
		ctx,
		"insert into short_url(code, original) values($1, $2)",
		url.Code,
		url.Original,
	)
	return err
}

func (rep *URLRepository) GetByCode(ctx context.Context, code string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.db.QueryRow(
		ctx,
		"SELECT code, original, created_at FROM short_url WHERE code = $1",
		code,
	).Scan(
		&url.Code,
		&url.Original,
		&url.CreatedAt,
	)
	return
}

func (rep *URLRepository) DeleteByCode(ctx context.Context, code string) error {
	_, err := rep.store.db.Exec(
		ctx,
		"delete from short_url where code = $1",
		code,
	)
	return err
}
