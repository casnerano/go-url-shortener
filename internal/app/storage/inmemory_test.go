package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

func TestInMemory_AddURL(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	m := NewInMemory()
	err := m.AddURL(*shortURLOne)
	require.NoError(t, err)

	got, err := m.GetURL(shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)
}

func TestInMemory_GetURL(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	m := NewInMemory()
	err := m.AddURL(*shortURLOne)
	require.NoError(t, err)

	got, err := m.GetURL(shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)

	_, err = m.GetURL("non-existent-code")
	assert.Error(t, err)
}

func TestNewInMemory(t *testing.T) {
	assert.Equal(t, InMemory{make(ShortURLDataStorage)}, *NewInMemory())
}
