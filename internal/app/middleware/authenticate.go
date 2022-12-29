package middleware

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/google/uuid"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/pkg/crypter"
)

type ContextUserUUIDType string

const (
	CookieUserUUIDKey                      = "SEC_USER_UUID"
	ContextUserUUIDKey ContextUserUUIDType = "user_uuid"
)

func Authenticate(key []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := wakeCookieUser(key, r)

			if err != nil {
				user, err = createCookieUser(key, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), ContextUserUUIDKey, user.UUID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func wakeCookieUser(key []byte, r *http.Request) (*model.User, error) {
	encryptUUIDCookie, err := r.Cookie(CookieUserUUIDKey)
	if err != nil {
		return nil, err
	}

	stUUID, err := decrypt(encryptUUIDCookie.Value, key)
	if err != nil {
		return nil, err
	}

	return model.NewUser(stUUID), nil
}

func createCookieUser(key []byte, w http.ResponseWriter) (*model.User, error) {
	gUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	encryptUUID, err := encrypt(gUUID.String(), key)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(gUUID.String())
	http.SetCookie(w, &http.Cookie{Name: CookieUserUUIDKey, Value: encryptUUID, Path: "/"})

	return user, nil
}

func encrypt(uuid string, key []byte) (string, error) {
	AES256GCM := crypter.NewAES256GCM(key)
	cipherUUID, err := AES256GCM.Encrypt([]byte(uuid))
	return base64.StdEncoding.EncodeToString(cipherUUID), err
}

func decrypt(cipher string, key []byte) (string, error) {
	AES256GCM := crypter.NewAES256GCM(key)

	bCipher, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return "", err
	}

	bUUID, err := AES256GCM.Decrypt(bCipher)
	if err != nil {
		return "", err
	}

	return string(bUUID), err
}
