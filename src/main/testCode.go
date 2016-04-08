package main

import (
	"fmt"
	"encoding/base64"
	"io"
	"crypto/rand"

)
func main1()  {
	token := make([]byte, 128)
	_, err := io.ReadFull(rand.Reader, token);
	if  err != nil {
	}
	s:= base64.URLEncoding.EncodeToString(token)
	fmt.Printf("%s",s)
}
