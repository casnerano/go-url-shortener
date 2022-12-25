package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/config"
	"github.com/casnerano/go-url-shortener/internal/app/middleware"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/service"
)

var (
	errBadRequest   = errors.New("bad request")
	errUnauthorized = errors.New("unauthorized")
)

type ShortURL struct {
	cfg        *config.Config
	urlService *service.URL
}

func NewShortURL(cfg *config.Config, urlService *service.URL) *ShortURL {
	return &ShortURL{cfg, urlService}
}

func (s *ShortURL) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	if shortCode == "" {
		http.Error(w, errBadRequest.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := s.urlService.GetByCode(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, shortURL.Original, http.StatusTemporaryRedirect)
}

func (s *ShortURL) PostText(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urlOriginal := string(body)

	if urlOriginal == "" {
		http.Error(w, errBadRequest.Error(), http.StatusBadRequest)
		return
	}

	ctxUID := r.Context().Value(middleware.ContextUserIDKey)
	uid, ok := ctxUID.(model.UserID)
	if !ok {
		http.Error(w, errUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	shortURLModel, err := s.urlService.Create(urlOriginal, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, s.buildAbsoluteShortURL(shortURLModel.Code))
}

func (s *ShortURL) PostJSON(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, s.createErrJSON(err.Error()), http.StatusInternalServerError)
		return
	}

	bodyObj := struct {
		URL string `json:"url"`
	}{}
	err = json.Unmarshal(body, &bodyObj)

	if err != nil || bodyObj.URL == "" {
		http.Error(w, s.createErrJSON(errBadRequest.Error()), http.StatusBadRequest)
		return
	}

	ctxUID := r.Context().Value(middleware.ContextUserIDKey)
	uid, ok := ctxUID.(model.UserID)
	if !ok {
		http.Error(w, errUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	shortURLModel, err := s.urlService.Create(bodyObj.URL, uid)
	if err != nil {
		http.Error(w, s.createErrJSON(err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	response := struct {
		Result string `json:"result"`
	}{
		s.buildAbsoluteShortURL(shortURLModel.Code),
	}

	rb, _ := json.Marshal(response)
	fmt.Fprint(w, string(rb))
}

func (s *ShortURL) buildAbsoluteShortURL(shortCode string) string {
	return fmt.Sprintf("%s/%s", s.cfg.Server.BaseURL, shortCode)
}

func (s *ShortURL) createErrJSON(err string) string {
	result, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{err})
	return string(result)
}
