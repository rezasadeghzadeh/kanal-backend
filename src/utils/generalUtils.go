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
	token := r.Header.Get("X-Auth-Token")
	log.Printf("Request sent with Token : %s \n",token)
	return token
}

func WriteJsonToResponse(w *http.ResponseWriter,b []byte){
	(*w).Header().Set("Content-Type","application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token")
	(*w).Write(b)
}