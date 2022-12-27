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
	errBadRequest     = errors.New("bad request")
	errUnauthorized   = errors.New("unauthorized")
	errServerInternal = errors.New("server internal error")
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

func (s *ShortURL) GetUserURLHistory(w http.ResponseWriter, r *http.Request) {
	type resultURLItem struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	ctxUID := r.Context().Value(middleware.ContextUserIDKey)
	uid, ok := ctxUID.(model.UserID)
	if !ok {
		s.httpJSONError(w, errUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	userURLHistory, err := s.urlService.FindByUser(uid)
	if err != nil {
		s.httpJSONError(w, errServerInternal.Error(), http.StatusInternalServerError)
		return
	}

	if len(userURLHistory) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resultHistory := []resultURLItem{}
	for _, userURL := range userURLHistory {
		item := resultURLItem{s.buildAbsoluteShortURL(userURL.Code), userURL.Original}
		resultHistory = append(resultHistory, item)
	}

	if err = json.NewEncoder(w).Encode(resultHistory); err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bodyObj := struct {
		URL string `json:"url"`
	}{}
	err = json.Unmarshal(body, &bodyObj)

	if err != nil || bodyObj.URL == "" {
		s.httpJSONError(w, errBadRequest.Error(), http.StatusBadRequest)
		return
	}

	ctxUID := r.Context().Value(middleware.ContextUserIDKey)
	uid, ok := ctxUID.(model.UserID)
	if !ok {
		s.httpJSONError(w, errUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	shortURLModel, err := s.urlService.Create(bodyObj.URL, uid)
	if err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
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

func (s *ShortURL) httpJSONError(w http.ResponseWriter, error string, code int) {
	jsonError, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{error})
	http.Error(w, string(jsonError), code)
}
