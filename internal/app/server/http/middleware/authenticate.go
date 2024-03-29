package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/service/crypter"
)

// ContextUserUUIDType for context keys.
type ContextUserUUIDType string

// Cookie key for user uuid.
const CookieUserUUIDKey = "SEC_USER_UUID"

// Context key for user uuid.
const ContextUserUUIDKey ContextUserUUIDType = "user_uuid"

// Authenticate middleware for user authentication by cookie.
// The cookie key is the value from the `CookieUserUUIDKey` constant,
// which specifies the user's UUID.
//
// Values in cookies are stored in encrypted form.
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

	stUUID, err := crypter.DecryptString(encryptUUIDCookie.Value, key)
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

	encryptUUID, err := crypter.EncryptString(gUUID.String(), key)
	if err != nil {
		return nil, err
	}

	user := model.NewUser(gUUID.String())
	http.SetCookie(w, &http.Cookie{Name: CookieUserUUIDKey, Value: encryptUUID, Path: "/"})

	return user, nil
}
