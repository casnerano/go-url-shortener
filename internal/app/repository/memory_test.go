package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

func TestMemory_Add(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	m := NewMemory()
	err := m.Add(context.Background(), *shortURLOne)
	require.NoError(t, err)

	got, err := m.GetByCode(context.Background(), shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)
}

func TestMemory_GetByCode(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	m := NewMemory()
	err := m.Add(context.Background(), *shortURLOne)
	require.NoError(t, err)

	got, err := m.GetByCode(context.Background(), shortURLOne.Code)
	require.NoError(t, err)

	assert.Equal(t, shortURLOne, got)

	_, err = m.GetByCode(context.Background(), "non-existent-code")
	assert.Error(t, err)
}

func TestNewMemory(t *testing.T) {
	assert.Equal(t, Memory{make(ShortURLDataStorage)}, *NewMemory())
}

func TestMemory_DeleteByCode(t *testing.T) {
	shortURLOne := model.NewShortURL("short#1", "large#1", time.Second)

	m := NewMemory()
	err := m.Add(context.Background(), *shortURLOne)
	require.NoError(t, err)

	err = m.DeleteByCode(context.Background(), shortURLOne.Code)
	require.NoError(t, err)
	assert.Equal(t, 0, len(m.store))
}
