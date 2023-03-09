package repository

import "errors"

// Repository error list.
var (
	ErrURLAlreadyExist    = errors.New("url already exists")
	ErrURLNotFound        = errors.New("url not found")
	ErrURLMarkedForDelete = errors.New("url marked for delete")
)

// Store interface for entity repositories.
type Store interface {
	URL() URLRepository
}
