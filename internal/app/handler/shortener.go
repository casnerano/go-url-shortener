package handler

import (
	"errors"
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
	rep  *repository.ShortURL
	hash hash.Hash
}

func NewShortener(r *repository.ShortURL, h hash.Hash) *Shortener {
	return &Shortener{rep: r, hash: h}
}

func (s *Shortener) URLGetHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")

	if shortCode == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	shortURL, err := s.rep.GetURLByCode(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if shortURL.CreatedAt.Add(shortURL.LifeTime).Before(time.Now()) {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, "url lifetime is expired")
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

	code, err := s.addShortURL(originalURL, 1*time.Minute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s://%s/%s", scheme, r.Host, code)
}

func (s *Shortener) addShortURL(url string, lifeTime time.Duration) (string, error) {
	h := s.hash.Generate(url)
	err := s.rep.AddURL(*model.NewShortURL(h, url, lifeTime))
	if err != nil {
		return "", errors.New("url adding error")
	}
	return h, nil
}
