package handler

import (
	"bytes"
	"fmt"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func TestNewShortener(t *testing.T) {
	shortURLRepository := repository.NewShortURL(storage.NewInMemory())
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	assert.Equal(t, Shortener{shortURLRepository, randHashService}, *shortener)
}

func TestShortener_URLHandler(t *testing.T) {
	const regexpHTTP = "^https?://"

	store := storage.NewInMemory()
	shortURLRepository := repository.NewShortURL(store)
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	t.Run("post url for shorten", func(t *testing.T) {
		store.Reset()

		body := []byte(`https://ya.ru`)
		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(shortener.URLHandler)

		handler.ServeHTTP(recorder, request)
		response := recorder.Result()

		require.Equal(t, http.StatusCreated, response.StatusCode)

		defer response.Body.Close()
		payload, err := io.ReadAll(response.Body)

		require.NoError(t, err)

		assert.Regexp(t, regexpHTTP, string(payload))
	})

	t.Run("get non-existent url", func(t *testing.T) {
		store.Reset()

		request := httptest.NewRequest(http.MethodGet, "/non-existent-code", nil)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(shortener.URLHandler)

		handler.ServeHTTP(recorder, request)
		response := recorder.Result()
		defer response.Body.Close()

		require.Equal(t, http.StatusNotFound, response.StatusCode)
	})

	t.Run("get existing url", func(t *testing.T) {
		store.Reset()
		shortURLOne := model.NewShortURL("short", "large", time.Second)

		err := store.AddURL(*shortURLOne)
		require.NoError(t, err)

		requestTarget := fmt.Sprintf("/%s", shortURLOne.Code)
		request := httptest.NewRequest(http.MethodGet, requestTarget, nil)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(shortener.URLHandler)

		handler.ServeHTTP(recorder, request)
		response := recorder.Result()
		defer response.Body.Close()

		require.Equal(t, http.StatusTemporaryRedirect, response.StatusCode)

		location := response.Header.Get("Location")
		require.NotEmpty(t, location)
	})
}

func TestShortener_addShortURL(t *testing.T) {
	shortURLRepository := repository.NewShortURL(storage.NewInMemory())
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	code, err := shortener.addShortURL("large#1", time.Minute)
	require.NoError(t, err)

	_, err = shortURLRepository.GetURLByCode(code)
	assert.NoError(t, err)
}

func TestShortener_extractLastPart(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"empty path", "", ""},
		{"one level path", "/one", "one"},
		{"one level path with slash", "/one/", ""},
		{"two level path", "/one/two", "two"},
		{"two level path with slash", "/one/two/", ""},
	}

	shortURLRepository := repository.NewShortURL(storage.NewInMemory())
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shortener.extractLastPart(tt.path))
		})
	}
}
