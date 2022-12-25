package middleware

import (
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"net/http"
)

type userID = uint64

const (
	CookieUserIDKey  = "SEC_UID"
	ContextUserIDKey = "uid"
)

func Authenticate(secretEncryptKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var currentUserID userID

			if encryptUserIDCookie, err := r.Cookie(CookieUserIDKey); err == nil {
				if uID, err := decrypt(encryptUserIDCookie.Value, secretEncryptKey); err == nil {
					currentUserID = uID
				}
			}

			if currentUserID == 0 {

				currentUserID = getRandomUserID()
				encryptUserID, err := encrypt(currentUserID, secretEncryptKey)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.SetCookie(w, &http.Cookie{Name: CookieUserIDKey, Value: encryptUserID})
			}

			ctx := context.WithValue(r.Context(), ContextUserIDKey, currentUserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getRandomUserID() userID {
	return rand.Uint64()
}

func encrypt(uID userID, key string) (string, error) {
	aesblock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	encryptUID := make([]byte, aes.BlockSize)
	bytesUID := make([]byte, aes.BlockSize)

	binary.LittleEndian.PutUint64(bytesUID, uID)
	aesblock.Encrypt(encryptUID, bytesUID)

	return base64.StdEncoding.EncodeToString(encryptUID), nil
}

func decrypt(encryptUID string, key string) (userID, error) {
	aesblock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return 0, err
	}

	decryptUID := make([]byte, aes.BlockSize)

	bEncryptUID, err := base64.StdEncoding.DecodeString(encryptUID)
	if err != nil {
		return 0, nil
	}

	aesblock.Decrypt(decryptUID, bEncryptUID)
	uID := binary.LittleEndian.Uint64(decryptUID)

	return uID, nil
}
