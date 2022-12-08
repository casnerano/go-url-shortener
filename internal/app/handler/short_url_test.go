package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/service"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
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

func TestNewShortURL(t *testing.T) {
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)
	shortener := NewShortURL(shortURLService)

	assert.Equal(t, ShortURL{shortURLService}, *shortener)
}

func TestShortURL_GetOriginalURL(t *testing.T) {
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)
	shortURLHandlerGroup := NewShortURL(shortURLService)

	router := chi.NewRouter()
	router.Get("/{shortCode}", shortURLHandlerGroup.GetOriginalURL)

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

func TestShortURL_PostText(t *testing.T) {
	const regexpHTTP = "^https?://"

	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)
	shortURLHandlerGroup := NewShortURL(shortURLService)

	router := chi.NewRouter()
	router.Post("/", shortURLHandlerGroup.PostText)

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

func TestShortURL_PostJSON(t *testing.T) {
	const regexpHTTP = "^https?://"

	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)
	shortURLHandlerGroup := NewShortURL(shortURLService)

	router := chi.NewRouter()
	router.Post("/api/shorten", shortURLHandlerGroup.PostJSON)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	t.Run("post url for shorten", func(t *testing.T) {
		body := []byte(`{"url": "http://ya.ru"}`)
		request, _ := http.NewRequest(http.MethodPost, testServer.URL+"/api/shorten", bytes.NewBuffer(body))
		statusCode, payload := testRequest(t, request)

		require.Equal(t, http.StatusCreated, statusCode)

		resultObj := struct {
			Result string `json:"result"`
		}{payload}
		err := json.Unmarshal([]byte(payload), &resultObj)
		require.NoError(t, err)

		assert.Regexp(t, regexpHTTP, resultObj.Result)
	})
}
