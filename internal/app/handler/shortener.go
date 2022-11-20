package handler

import (
	"errors"
	"fmt"
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
	"github.com/casnerano/go-url-shortener/internal/app/service/url/hash"
)

type Shortener struct {
	rep  *repository.ShortUrl
	hash hash.Hash
}

func (s *Shortener) UrlHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		lastPart := s.extractLastPart(r.URL.Path)

		shortUrl, err := s.rep.GetUrlByCode(lastPart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if shortUrl.CreatedAt.Add(shortUrl.LifeTime).Before(time.Now()) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "url lifetime is expired")
			return
		}

		http.Redirect(w, r, shortUrl.Original, http.StatusTemporaryRedirect)

	case http.MethodPost:

		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		originalUrl := string(body)

		if originalUrl == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		code, err := s.addShortUrl(originalUrl, 1*time.Minute)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s/%s", r.Host, code)
	}
}

func (s *Shortener) addShortUrl(url string, lifeTime time.Duration) (string, error) {
	h := s.hash.Generate(url)
	err := s.rep.AddUrl(*model.NewShortUrl(h, url, lifeTime))
	if err != nil {
		return "", errors.New("url adding error")
	}
	return h, nil
}

func (s *Shortener) extractLastPart(path string) string {
	if index := strings.LastIndex(path, "/"); index >= 0 {
		return path[index+1:]
	}
	return path
}

func NewShortener(r *repository.ShortUrl, h hash.Hash) *Shortener {
	return &Shortener{rep: r, hash: h}
}
