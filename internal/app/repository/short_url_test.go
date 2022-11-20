package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func TestNewShortURL(t *testing.T) {
	store := storage.NewInMemory()
	assert.Equal(t, ShortURL{storage: store}, *NewShortURL(store))
}

func TestShortURL_AddURL(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	store := storage.NewInMemory()
	rep := NewShortURL(store)

	err := rep.AddURL(*shortURLOne)
	require.NoError(t, err)

	got, err := rep.GetURLByCode(shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)
}

func TestShortURL_GetURLByCode(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	store := storage.NewInMemory()
	rep := NewShortURL(store)

	err := rep.AddURL(*shortURLOne)
	require.NoError(t, err)

	got, err := rep.GetURLByCode(shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)

	_, err = rep.GetURLByCode("non-existent-code")
	assert.Error(t, err)
}
