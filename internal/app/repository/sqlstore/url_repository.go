package sqlstore

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

// URLRepository structure for url repository with sql store.
type URLRepository struct {
	store *Store
}

// Adding entity.
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
		return repository.ErrURLAlreadyExist
	}

	return err
}

// Batch adding entities.
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

// Get entity by short code.
func (rep *URLRepository) GetByCode(ctx context.Context, code string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT id, code, original, user_uuid, created_at, deleted FROM short_url WHERE code = $1",
		code,
	).Scan(
		&url.ID,
		&url.Code,
		&url.Original,
		&url.UserUUID,
		&url.CreatedAt,
		&url.Deleted,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrURLNotFound
	}

	if url.Deleted {
		return nil, repository.ErrURLMarkedForDelete
	}

	return
}

// Get entity by user uuid and original url.
func (rep *URLRepository) GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (url *model.ShortURL, err error) {
	url = &model.ShortURL{}
	err = rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT id, code, original, user_uuid, created_at, deleted FROM short_url WHERE user_uuid = $1 and original = $2",
		uuid,
		original,
	).Scan(
		&url.ID,
		&url.Code,
		&url.Original,
		&url.UserUUID,
		&url.CreatedAt,
		&url.Deleted,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrURLNotFound
	}

	if url.Deleted {
		return nil, repository.ErrURLMarkedForDelete
	}

	return
}

// Find entities by user uuid.
func (rep *URLRepository) FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error) {
	collection := []*model.ShortURL{}

	rows, err := rep.store.pgxpool.Query(
		ctx,
		"SELECT id, code, original, user_uuid, created_at FROM short_url WHERE user_uuid = $1 and deleted is false",
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

// Delete entity by short code.
func (rep *URLRepository) DeleteByCode(ctx context.Context, code string, uuid string) error {
	res, _ := rep.store.pgxpool.Exec(
		ctx,
		"update short_url set deleted = true where code = $1 and user_uuid = $2",
		code,
		uuid,
	)

	if res.RowsAffected() > 0 {
		return nil
	}

	return repository.ErrURLNotFound
}

// Batch delete entities by codes.
func (rep *URLRepository) DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error {
	_, err := rep.store.pgxpool.Exec(
		ctx,
		"update short_url set deleted = true where user_uuid = $1 and code = any($2::text[])",
		uuid,
		pq.Array(codes),
	)
	return err
}

// Delete older entities for duration.
func (rep *URLRepository) DeleteOlderRows(ctx context.Context, d time.Duration) error {
	_, err := rep.store.pgxpool.Exec(
		ctx,
		"update short_url set deleted = true where created_at > $1",
		time.Now().Add(d),
	)
	return err
}

// Get total Short URL count
func (rep *URLRepository) GetTotalURLCount(ctx context.Context) (int, error) {
	count := 0
	err := rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT count(id) FROM short_url",
	).Scan(
		&count,
	)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get total User count
func (rep *URLRepository) GetTotalUserCount(ctx context.Context) (int, error) {
	count := 0
	err := rep.store.pgxpool.QueryRow(
		ctx,
		"SELECT DISTINCT count(user_uuid) FROM short_url",
	).Scan(
		&count,
	)
	if err != nil {
		return 0, err
	}

	return count, nil
}
