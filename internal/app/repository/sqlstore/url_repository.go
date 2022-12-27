package sqlstore

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(ctx context.Context, url *model.ShortURL) error {
	err := rep.store.db.QueryRow(
		ctx,
		"insert into short_url(code, original) values($1, $2) returning id, created_at",
		url.Code,
		url.Original,
	).Scan(
		&url.ID,
		&url.CreatedAt,
	)
	return err
}

func (rep *URLRepository) GetByCode(ctx context.Context, code string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.db.QueryRow(
		ctx,
		"SELECT id, code, original, user_id, created_at FROM short_url WHERE code = $1",
		code,
	).Scan(
		&url.ID,
		&url.Code,
		&url.Original,
		&url.UserID,
		&url.CreatedAt,
	)
	return
}

func (rep *URLRepository) FindByUser(ctx context.Context, uid model.UserID) ([]*model.ShortURL, error) {
	collection := make([]*model.ShortURL, 10)

	rows, err := rep.store.db.Query(
		ctx,
		"SELECT id, code, original, user_id, created_at FROM short_url WHERE user_id = $1",
		uid,
	)

	if err != nil {
		return collection, err
	}

	for rows.Next() {
		url := &model.ShortURL{}
		err = rows.Scan(
			&url.ID,
			&url.Code,
			&url.Original,
			&url.UserID,
			&url.CreatedAt,
		)
		if err == nil {
			collection = append(collection, url)
		}
	}

	return collection, nil
}

func (rep *URLRepository) DeleteByCode(ctx context.Context, code string) error {
	_, err := rep.store.db.Exec(
		ctx,
		"delete from short_url where code = $1",
		code,
	)
	return err
}

func (rep *URLRepository) DeleteOlderRows(ctx context.Context, d time.Duration) error {
	_, err := rep.store.db.Exec(
		ctx,
		"delete from short_url where created_at > $1",
		time.Now().Add(d),
	)
	return err
}
