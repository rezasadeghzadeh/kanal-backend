package utils

import (
	"time"
	"io"
	"encoding/base64"
	"crypto/rand"
	"strings"
	"log"
	"net/http"


)

import m "math/rand"

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

func GetRandomFileName(n int) (string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[m.Intn(len(letters))]
	}
	return string(b)
}

func GetFileExtention(fileName string) (string){
	dotPosition := strings.LastIndex(fileName,".")
	ext := fileName[dotPosition:len(fileName)]
	log.Printf("FileName: %s  Ext: %s",fileName,ext)
	return ext
}

func GetUserTokenFromReq(r *http.Request) (string){
	return r.Header.Get("X-Auth-Token")
}