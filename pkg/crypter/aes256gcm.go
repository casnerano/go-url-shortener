package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type AES256GCM struct {
	key []byte
}

func NewAES256GCM(key []byte) *AES256GCM {
	hKey := sha256.Sum256(key)
	return &AES256GCM{key: hKey[:]}
}

func (aes256gcm *AES256GCM) Encrypt(src []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(aes256gcm.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonce, err := aes256gcm.getRandomBytes(aesgcm.NonceSize())
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nonce, nonce, src, nil), nil
}

func (aes256gcm *AES256GCM) Decrypt(dst []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(aes256gcm.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	nonce, dst := dst[:nonceSize], dst[nonceSize:]

	return aesgcm.Open(nil, nonce, dst, nil)
}

func (aes256gcm *AES256GCM) getRandomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, b)
	return b, err
}
