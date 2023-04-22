package crypter

import (
	"encoding/base64"

	"github.com/casnerano/go-url-shortener/pkg/crypter"
)

// EncryptString helper for AES256GCM encrypt
func EncryptString(uuid string, key []byte) (string, error) {
	AES256GCM := crypter.NewCipher(key)
	cipherUUID, err := AES256GCM.Encrypt([]byte(uuid))
	return base64.StdEncoding.EncodeToString(cipherUUID), err
}

// DecryptString helper for AES256GCM decrypt
func DecryptString(cipher string, key []byte) (string, error) {
	AES256GCM := crypter.NewCipher(key)

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
