package utils

import (
	"time"
	"io"
	"encoding/base64"
	"crypto/rand"
)

func CurrentMilis() (int64)  {
	return time.Now().UnixNano() / 1000000
}

func GetToken() (string) {
	token := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, token);
	if  err != nil {
	}
	return base64.URLEncoding.EncodeToString(token)
}

