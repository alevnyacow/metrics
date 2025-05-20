package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignedSHA256(data []byte, key string) (string, error) {
	mac := hmac.New(sha256.New, []byte(key))
	_, err := mac.Write([]byte("hello"))
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	_, hashError := hash.Write(data)
	if hashError != nil {
		return "", hashError
	}
	result := mac.Sum(nil)
	return hex.EncodeToString(result), nil
}
