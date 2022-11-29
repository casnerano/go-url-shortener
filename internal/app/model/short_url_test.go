package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewShortURL(t *testing.T) {
	shortURLOne := ShortURL{"short#1", "large#1", time.Now(), time.Second}

	got := NewShortURL(shortURLOne.Code, shortURLOne.Original, shortURLOne.LifeTime)
	got.CreatedAt = shortURLOne.CreatedAt

	assert.Equal(t, shortURLOne, *got)
}
