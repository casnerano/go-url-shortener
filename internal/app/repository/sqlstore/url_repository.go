package sqlstore

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(ctx context.Context, url *model.ShortURL) error {
	err := rep.store.pgxpool.QueryRow(
		ctx,
		"insert into short_url(code, original, user_uuid) values($1, $2, $3) returning id, created_at",
		url.Code,
		url.Original,
		url.UserUUID,
	).Scan(
		&url.ID,
		&url.CreatedAt,
	)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return repository.ErrURLExist
	}

	return err
}

func (rep *URLRepository) AddBatch(ctx context.Context, urls []*model.ShortURL) error {
	batch := &pgx.Batch{}

	for _, url := range urls {
		batch.Queue(
			"insert into short_url(code, original, user_uuid) values($1, $2, $3)",
			url.Code,
			url.Original,
			url.UserUUID,
		)
	}

	br := rep.store.pgxpool.SendBatch(ctx, batch)
	defer br.Close()

	_, err := br.Exec()

	return err
}

func (rep *URLRepository) GetByCode(ctx context.Context, code string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT id, code, original, user_uuid, created_at FROM short_url WHERE code = $1",
		code,
	).Scan(
		&url.ID,
		&url.Code,
		&url.Original,
		&url.UserUUID,
		&url.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrURLNotFound
	}

	return
}

func (rep *URLRepository) GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT id, code, original, user_uuid, created_at FROM short_url WHERE user_uuid = $1 and original = $2",
		uuid,
		original,
	).Scan(
		&url.ID,
		&url.Code,
		&url.Original,
		&url.UserUUID,
		&url.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrURLNotFound
	}

	return
}

func (rep *URLRepository) FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error) {
	collection := []*model.ShortURL{}

	rows, err := rep.store.pgxpool.Query(
		ctx,
		"SELECT id, code, original, user_uuid, created_at FROM short_url WHERE user_uuid = $1",
		uuid,
	)

	if err != nil {
		return collection, err
	}

	defer rows.Close()

	for rows.Next() {
		url := &model.ShortURL{}
		err = rows.Scan(
			&url.ID,
			&url.Code,
			&url.Original,
			&url.UserUUID,
			&url.CreatedAt,
		)
		if err == nil {
			collection = append(collection, url)
		}
	}

	return collection, nil
}

func (rep *URLRepository) DeleteByCode(ctx context.Context, code string) error {
	_, err := rep.store.pgxpool.Exec(
		ctx,
		"delete from short_url where code = $1",
		code,
	)
	return err
}

func (rep *URLRepository) DeleteOlderRows(ctx context.Context, d time.Duration) error {
	_, err := rep.store.pgxpool.Exec(
		ctx,
		"delete from short_url where created_at > $1",
		time.Now().Add(d),
	)
	return err
}
