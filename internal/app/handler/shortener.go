package handler

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
)

type Shortener struct {
	rep  repository.URLRepository
	hash hash.Hash
}

func NewShortener(r repository.URLRepository, h hash.Hash) *Shortener {
	return &Shortener{rep: r, hash: h}
}

func (s *Shortener) URLGetHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	if shortCode == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	shortURL, err := s.rep.GetByCode(r.Context(), shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, shortURL.Original, http.StatusTemporaryRedirect)
}

func (s *Shortener) URLPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	originalURL := string(body)

	if originalURL == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	lifeTime := 1 * time.Minute
	code := s.hash.Generate(originalURL)
	err = s.rep.Add(r.Context(), *model.NewShortURL(code, originalURL, lifeTime))
	if err != nil {
		http.Error(w, "url adding error", http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s://%s/%s", scheme, r.Host, code)
}
