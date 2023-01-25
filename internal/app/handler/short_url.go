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
	"github.com/casnerano/go-url-shortener/internal/app/repository"
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
		if errors.Is(err, repository.ErrURLMarkedForDelete) {
			http.Error(w, err.Error(), http.StatusGone)
			return
		}
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

	ctxUUID := r.Context().Value(middleware.ContextUserUUIDKey)

	uuid, ok := ctxUUID.(string)
	if !ok {
		s.httpJSONError(w, errUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	userURLHistory, err := s.urlService.FindByUserUUID(uuid)
	if err != nil {
		s.httpJSONError(w, errServerInternal.Error(), http.StatusInternalServerError)
		return
	}

	if len(userURLHistory) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resultHistory := make([]resultURLItem, 0, len(userURLHistory))
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

	ctxUUID := r.Context().Value(middleware.ContextUserUUIDKey)
	uuid, ok := ctxUUID.(string)
	if !ok {
		uuid = ""
	}

	shortURLModel, err := s.urlService.Create(urlOriginal, uuid)
	if err != nil {
		if errors.Is(err, repository.ErrURLAlreadyExist) {
			w.WriteHeader(http.StatusConflict)
			shortURLModel, err = s.urlService.GetByUserUUIDAndOriginal(uuid, urlOriginal)
			if err != nil {
				s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

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

	ctxUUID := r.Context().Value(middleware.ContextUserUUIDKey)
	uuid, ok := ctxUUID.(string)
	if !ok {
		uuid = ""
	}

	shortURLModel, err := s.urlService.Create(bodyObj.URL, uuid)
	if err != nil {
		if errors.Is(err, repository.ErrURLAlreadyExist) {
			w.WriteHeader(http.StatusConflict)
			shortURLModel, err = s.urlService.GetByUserUUIDAndOriginal(uuid, bodyObj.URL)
			if err != nil {
				s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	response := struct {
		Result string `json:"result"`
	}{
		s.buildAbsoluteShortURL(shortURLModel.Code),
	}

	rb, _ := json.Marshal(response)
	fmt.Fprint(w, string(rb))
}

func (s *ShortURL) PostBatchJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var batchRequest []*model.ShortURLBatchRequest
	err = json.Unmarshal(body, &batchRequest)

	if err != nil {
		s.httpJSONError(w, errBadRequest.Error(), http.StatusBadRequest)
		return
	}

	ctxUUID := r.Context().Value(middleware.ContextUserUUIDKey)
	uuid, ok := ctxUUID.(string)
	if !ok {
		uuid = ""
	}

	response, err := s.urlService.CreateBatch(batchRequest, uuid)
	if err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k := range response {
		response[k].ShortURL = s.buildAbsoluteShortURL(response[k].ShortURL)
	}

	w.WriteHeader(http.StatusCreated)

	rb, _ := json.Marshal(response)
	fmt.Fprint(w, string(rb))
}

func (s *ShortURL) DeleteBatchJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		s.httpJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var codes []string
	err = json.Unmarshal(body, &codes)

	if err != nil {
		s.httpJSONError(w, errBadRequest.Error(), http.StatusBadRequest)
		return
	}

	ctxUUID := r.Context().Value(middleware.ContextUserUUIDKey)
	uuid, ok := ctxUUID.(string)
	if !ok {
		uuid = ""
	}

	go func() {
		_ = s.urlService.DeleteBatch(codes, uuid)
	}()

	w.WriteHeader(http.StatusAccepted)
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
