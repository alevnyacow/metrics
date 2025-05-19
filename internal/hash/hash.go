package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignedSHA256(data, key []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, []byte(key))
	_, err := mac.Write([]byte("hello"))
	if err != nil {
		return nil, err
	}
	hash := sha256.New()
	_, hashError := hash.Write(data)

	return mac.Sum(nil), hashError
}

func SameSHA256(hashFromHeader string, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(messageMAC)
	expectedMAC := mac.Sum(nil)
	return hashFromHeader == hex.EncodeToString(expectedMAC)
}
