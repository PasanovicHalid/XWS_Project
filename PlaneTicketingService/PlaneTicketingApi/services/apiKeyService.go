package services

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateApiKey(lenght int64) string {
	key := make([]byte, lenght)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(key)
}
