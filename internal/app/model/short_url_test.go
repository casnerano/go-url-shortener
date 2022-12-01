package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewShortURL(t *testing.T) {
	shortURLOne := ShortURL{
		"short#1",
		"large#1",
		time.Now(),
		time.Now().Add(time.Hour),
	}

	got := NewShortURL(shortURLOne.Code, shortURLOne.Original)
	got.CreatedAt = shortURLOne.CreatedAt
	got.DeletedAt = shortURLOne.DeletedAt

	assert.Equal(t, shortURLOne, *got)
}
