package repository

import "errors"

var (
	ErrURLExist    = errors.New("url already exists")
	ErrURLNotFound = errors.New("url not found")
)

type Store interface {
	URL() URLRepository
}
