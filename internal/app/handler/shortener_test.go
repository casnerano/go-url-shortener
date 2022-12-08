package handler

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/service/hash"
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
	conf := &config.Config{}
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(conf, URLRepository, randHashService)

	assert.Equal(t, Shortener{conf, URLRepository, randHashService}, *shortener)
}

func TestShortener_URLGetHandler(t *testing.T) {
	conf := &config.Config{}
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(conf, URLRepository, randHashService)

	router := chi.NewRouter()
	router.Get("/{shortCode}", shortener.URLGetHandler)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	t.Run("get non-existent url", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/non-existent-code", nil)
		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusNotFound, statusCode)
	})

	t.Run("get existing url", func(t *testing.T) {
		shortURLOne := model.NewShortURL("short", "large")

		err := URLRepository.Add(context.Background(), *shortURLOne)
		require.NoError(t, err)

		requestTarget := fmt.Sprintf("%s/%s", testServer.URL, shortURLOne.Code)

		request, _ := http.NewRequest(http.MethodGet, requestTarget, nil)
		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusTemporaryRedirect, statusCode)
	})
}

func TestShortener_URLPostHandler(t *testing.T) {
	const regexpHTTP = "^https?://"

	conf := &config.Config{}
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hash.NewRandom(1, 1)
	shortener := NewShortener(conf, URLRepository, randHashService)

	router := chi.NewRouter()
	router.Post("/", shortener.URLPostHandler)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	t.Run("post url for shorten", func(t *testing.T) {
		body := []byte(`https://ya.ru`)
		request, _ := http.NewRequest(http.MethodPost, testServer.URL+"/", bytes.NewBuffer(body))
		statusCode, payload := testRequest(t, request)

		require.Equal(t, http.StatusCreated, statusCode)
		assert.Regexp(t, regexpHTTP, payload)
	})
}
