package handler

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/go-chi/chi/v5"

    "github.com/casnerano/go-url-shortener/internal/app/service"
)

type ShortURL struct {
    urlService *service.URL
}

func NewShortURL(urlService *service.URL) *ShortURL {
    return &ShortURL{urlService}
}

func (s *ShortURL) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
    shortCode := chi.URLParam(r, "shortCode")

    if shortCode == "" {
        http.Error(w, "bad request", http.StatusBadRequest)
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
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    shortURLModel, err := s.urlService.Create(urlOriginal)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, s.buildAbsoluteShortURL(r, shortURLModel.Code))
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
        http.Error(w, s.createErrJSON("bad request"), http.StatusBadRequest)
        return
    }

    shortURLModel, err := s.urlService.Create(bodyObj.URL)
    if err != nil {
        http.Error(w, s.createErrJSON(err.Error()), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)

    response := struct {
        Result string `json:"result"`
    }{
        s.buildAbsoluteShortURL(r, shortURLModel.Code),
    }

    rb, _ := json.Marshal(response)
    fmt.Fprint(w, string(rb))
}

func (s *ShortURL) buildAbsoluteShortURL(r *http.Request, shortCode string) string {
    scheme := "http"
    if r.TLS != nil {
        scheme = "https"
    }
    return fmt.Sprintf("%s://%s/%s", scheme, r.Host, shortCode)
}

func (s *ShortURL) createErrJSON(err string) string {
    result, _ := json.Marshal(struct {
        Error string `json:"error"`
    }{err})
    return string(result)
}
