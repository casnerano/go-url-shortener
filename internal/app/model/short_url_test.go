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
		time.Now(),
	}

	got := NewShortURL(shortURLOne.Code, shortURLOne.Original)
	got.ID = shortURLOne.ID
	got.CreatedAt = shortURLOne.CreatedAt

	assert.Equal(t, shortURLOne, *got)
}
