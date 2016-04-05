package main

import (
	"log"
	"time"
	"net/http"
	"../controller"
	"strconv"

	"encoding/json"
	"../dao"
)

func InitHttpHandlers() {
	log.Printf("Start Http Initalization\n")
	//channels list handler
	channelsListHandler := func(w http.ResponseWriter, r *http.Request) {
		start  := time.Now().UnixNano()
		b := controller.GetAllChannels()
		w.Header().Set("Content-Type","application/json")
		w.Write(b)
		end := time.Now().UnixNano()
		duration :=  end - start
		log.Printf("start:%d end:%d      %d ns",start,end,duration)
	}
	http.Handle("/channels/list",http.HandlerFunc(channelsListHandler))

	alreadyExistsChannelNameHandler := func(w http.ResponseWriter, r *http.Request) {
		start  := time.Now().UnixNano()
		channelName  := r.URL.Query().Get("name");
		b := controller.AlreadyExistsChannelName(channelName)
		w.Header().Set("Content-Type","application/json")
		byteValue := []byte(strconv.Itoa(b))
		w.Write(byteValue)
		end := time.Now().UnixNano()
		duration :=  end - start
		log.Printf("alreadyExistsChannelNameHandler :   %d ns",duration)
	}
	http.Handle("/channels/alreadyExistsChannelName",http.HandlerFunc(alreadyExistsChannelNameHandler))

	saveNewChannel := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var channel dao.Channel
		err := decoder.Decode(&channel)
		if(err != nil){
			log.Printf("%v",err)
		}
		start  := time.Now().UnixNano()

		b := controller.SaveNewChannel(channel.Name,channel.Title,channel.Description,channel.ChannelType)
		w.Header().Set("Content-Type","application/json")
		byteValue := []byte(strconv.Itoa(b))
		w.Write(byteValue)
		end := time.Now().UnixNano()
		duration :=  end - start
		log.Printf("saveNewChannel:   %d ns ",duration)
	}
	http.Handle("/channels/saveNewChannel",http.HandlerFunc(saveNewChannel))


}
