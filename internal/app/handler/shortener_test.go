package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

func testRequest(t *testing.T, r *http.Request) (int, string) {
	client := resty.New()
	baseURL := fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host)

	getDisableRedirectPolity := func() resty.RedirectPolicyFunc {
		return func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	resp, err := client.
		SetBaseURL(baseURL).
		SetRedirectPolicy(getDisableRedirectPolity()).
		R().SetBody(r.Body).Execute(r.Method, r.URL.Path)

	require.NoError(t, err)

	return resp.StatusCode(), string(resp.Body())
}

func TestNewShortener(t *testing.T) {
	shortURLRepository := repository.NewShortURL(storage.NewMemory())
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	assert.Equal(t, Shortener{shortURLRepository, randHashService}, *shortener)
}

func TestShortener_addShortURL(t *testing.T) {
	shortURLRepository := repository.NewShortURL(storage.NewMemory())
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	code, err := shortener.addShortURL("large#1", time.Minute)
	require.NoError(t, err)

	_, err = shortURLRepository.GetURLByCode(code)
	assert.NoError(t, err)
}

func TestShortener_URLGetHandler(t *testing.T) {
	store := storage.NewMemory()
	shortURLRepository := repository.NewShortURL(store)
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	router := chi.NewRouter()
	router.Get("/{shortCode}", shortener.URLGetHandler)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	t.Run("get non-existent url", func(t *testing.T) {
		store.Reset()

		request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/non-existent-code", nil)
		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("get existing url", func(t *testing.T) {
		store.Reset()
		shortURLOne := model.NewShortURL("short", "large", time.Hour)

		err := store.AddURL(*shortURLOne)
		require.NoError(t, err)

		requestTarget := fmt.Sprintf("%s/%s", testServer.URL, shortURLOne.Code)

		request, _ := http.NewRequest(http.MethodGet, requestTarget, nil)
		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusTemporaryRedirect, statusCode)
	})
}

func TestShortener_URLPostHandler(t *testing.T) {
	const regexpHTTP = "^https?://"

	store := storage.NewMemory()
	shortURLRepository := repository.NewShortURL(store)
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(shortURLRepository, randHashService)

	router := chi.NewRouter()
	router.Post("/", shortener.URLPostHandler)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	t.Run("post url for shorten", func(t *testing.T) {
		store.Reset()

		body := []byte(`https://ya.ru`)
		request, _ := http.NewRequest(http.MethodPost, testServer.URL+"/", bytes.NewBuffer(body))
		statusCode, payload := testRequest(t, request)

		require.Equal(t, http.StatusCreated, statusCode)
		assert.Regexp(t, regexpHTTP, payload)
	})
}
