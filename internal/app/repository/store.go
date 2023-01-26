package repository

import "errors"

var (
	ErrURLAlreadyExist    = errors.New("url already exists")
	ErrURLNotFound        = errors.New("url not found")
	ErrURLMarkedForDelete = errors.New("url marked for delete")
)

type Store interface {
	URL() URLRepository
}
