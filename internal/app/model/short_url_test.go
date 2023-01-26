package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewShortURL(t *testing.T) {
	shortURLOne := ShortURL{
		1,
		"short#1",
		"large#1",
		"uuid",
		time.Now(),
		false,
	}

	got := NewShortURL(shortURLOne.Code, shortURLOne.Original)
	got.ID = shortURLOne.ID
	got.UserUUID = shortURLOne.UserUUID
	got.CreatedAt = shortURLOne.CreatedAt
	got.Deleted = shortURLOne.Deleted

	assert.Equal(t, shortURLOne, *got)
}
