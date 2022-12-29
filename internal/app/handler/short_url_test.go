package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/middleware"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository/memstore"
	"github.com/casnerano/go-url-shortener/internal/app/service"
	"github.com/casnerano/go-url-shortener/internal/app/service/hasher"
	"github.com/casnerano/go-url-shortener/pkg/crypter"
)

func testRequest(t *testing.T, r *http.Request) (int, string) {
	client := resty.New()
	baseURL := fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host)

	getDisableRedirectPolity := func() resty.RedirectPolicyFunc {
		return func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	request := client.
		SetBaseURL(baseURL).
		SetRedirectPolicy(getDisableRedirectPolity()).
		R()

	response, err := request.
		SetBody(r.Body).
		SetCookies(r.Cookies()).
		Execute(r.Method, r.URL.Path)

	require.NoError(t, err)

	return response.StatusCode(), string(response.Body())
}

func TestNewShortURL(t *testing.T) {
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)

	cfg := config.New()
	shortener := NewShortURL(cfg, shortURLService)

	assert.Equal(t, ShortURL{cfg, shortURLService}, *shortener)
}

func TestShortURL_GetOriginalURL(t *testing.T) {
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)

	cfg := config.New()
	shortURLHandlerGroup := NewShortURL(cfg, shortURLService)

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

		err := URLRepository.Add(context.Background(), shortURLOne)
		require.NoError(t, err)

		requestTarget := fmt.Sprintf("%s/%s", testServer.URL, shortURLOne.Code)

		request, _ := http.NewRequest(http.MethodGet, requestTarget, nil)
		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusTemporaryRedirect, statusCode)
	})
}

func TestShortURL_GetUserURLHistory(t *testing.T) {
	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)

	cfg := config.New()
	shortURLHandlerGroup := NewShortURL(cfg, shortURLService)

	key := []byte("#easy_secret_key")

	router := chi.NewRouter()
	router.Use(middleware.Authenticate(key))
	router.Get("/api/user/urls", shortURLHandlerGroup.GetUserURLHistory)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	shortURLOne := &model.ShortURL{Code: "short1", Original: "large1", UserUUID: "test_uuid"}
	shortURLTwo := &model.ShortURL{Code: "short2", Original: "large2", UserUUID: "test_uuid"}

	err := URLRepository.Add(context.Background(), shortURLOne)
	require.NoError(t, err)

	err = URLRepository.Add(context.Background(), shortURLTwo)
	require.NoError(t, err)

	AES256GCM := crypter.NewAES256GCM(key)

	t.Run("get user positive url history", func(t *testing.T) {
		cipherUUID, _ := AES256GCM.Encrypt([]byte("test_uuid"))
		encryptUUID := base64.StdEncoding.EncodeToString(cipherUUID)

		cookie := http.Cookie{Name: middleware.CookieUserUUIDKey, Value: encryptUUID, Path: "/"}

		request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/api/user/urls", nil)
		request.AddCookie(&cookie)

		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusOK, statusCode)
	})

	t.Run("get user negative url history", func(t *testing.T) {
		cipherUUID, _ := AES256GCM.Encrypt([]byte("non_exist_uuid"))
		encryptUUID := base64.StdEncoding.EncodeToString(cipherUUID)

		cookie := http.Cookie{Name: middleware.CookieUserUUIDKey, Value: encryptUUID, Path: "/"}

		request, _ := http.NewRequest(http.MethodGet, testServer.URL+"/api/user/urls", nil)
		request.AddCookie(&cookie)

		statusCode, _ := testRequest(t, request)
		require.Equal(t, http.StatusNoContent, statusCode)
	})
}

func TestShortURL_PostText(t *testing.T) {
	const regexpHTTP = "^https?://"

	URLRepository := memstore.NewStore().URL()
	randHashService, _ := hasher.NewRandom(1, 1)
	shortURLService := service.NewURL(URLRepository, randHashService)

	cfg := config.New()
	shortURLHandlerGroup := NewShortURL(cfg, shortURLService)

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

	cfg := config.New()
	shortURLHandlerGroup := NewShortURL(cfg, shortURLService)

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
