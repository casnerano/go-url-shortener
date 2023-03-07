package handler

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/casnerano/go-url-shortener/internal/app/middleware"
)

func ExampleShortURL_PostText() {
	client := resty.New()

	request := client.SetBaseURL("http://127.0.0.1:8081").
		R().
		SetCookie(
			&http.Cookie{
				Name:  middleware.CookieUserUUIDKey,
				Value: "[EncryptedUUID]",
				Path:  "/",
			},
		)

	response, err := request.
		SetBody("https://ya.ru/long-example-url").
		Post("/")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode())
	}
}

func ExampleShortURL_PostJSON() {
	client := resty.New()

	request := client.SetBaseURL("http://127.0.0.1:8081").
		R().
		SetCookie(
			&http.Cookie{
				Name:  middleware.CookieUserUUIDKey,
				Value: "[EncryptedUUID]",
				Path:  "/",
			},
		)

	response, err := request.
		SetHeader("Content-Type", "application/json").
		SetBody(`{"url": "https://ya.ru/long-example-url"}`).
		Post("/api/shorten")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode())
	}
}
