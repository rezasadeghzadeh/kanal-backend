package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	InitHttpHandlers()
	InitHttpServer()
}

func InitHttpServer()  {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
