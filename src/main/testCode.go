package main

import (


	"strings"
	"log"
)
func main1()  {
	fileName :=  "image.jpg"
	dotPosition := strings.LastIndex(fileName,".")
	log.Printf("%d",dotPosition)
	ext := fileName[dotPosition:len(fileName)]

	log.Printf("FileName: %s  Ext: %s",fileName,ext)
}
