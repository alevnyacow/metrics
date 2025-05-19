package hash

import (
	"crypto/hmac"
	"crypto/sha256"
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

func SameSHA256(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
