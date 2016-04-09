package main

import (
	"log"
	"net/http"
	"../controller/channelsController"
	"../controller/smsVerificationController"
	"strconv"
	"encoding/json"
	"../dao/channelDao"
	"../controller/userController"
)

func InitHttpHandlers() {
	log.Printf("Start Http Initalization\n")
	channelsBinding()
	userBinding()
	smsVerificationBinding()
}

type response struct {
	status string
	data string
}

func smsVerificationBinding()  {
	sendSmsVerificationHandler := func(w http.ResponseWriter, r *http.Request) {
		mobileNumber := r.URL.Query().Get("mobileNumber")
		log.Printf("Mobilenumber : %s",mobileNumber)
		b,_ := json.Marshal(strconv.Itoa(smsVerificationController.SendSmsVerification(mobileNumber)))
		writeJsonToResponse(&w,b)
	}
	http.Handle("/smsVerification/send",http.HandlerFunc(sendSmsVerificationHandler))

	verifySentSmsHandler := func(w http.ResponseWriter, r *http.Request) {
		mobileNumber := r.URL.Query().Get("mobileNumber")
		enteredText  := r.URL.Query().Get("enteredText")
		b := []byte(strconv.Itoa(smsVerificationController.VerifySentSms(mobileNumber,enteredText)))
		writeJsonToResponse(&w,b)
	}
	http.Handle("/smsVerification/verify",http.HandlerFunc(verifySentSmsHandler))
}


func userBinding()  {
	userUpdateHandler := func(w http.ResponseWriter, r *http.Request) {
		mobileNumber := r.URL.Query().Get("mobileNumber")
		email := r.URL.Query().Get("email")
		firstName := r.URL.Query().Get("firstName")
		lastName := r.URL.Query().Get("lastName")
		b,_ := json.Marshal(userController.UpsertAndGetToken(mobileNumber,email,firstName,lastName))
		writeJsonToResponse(&w,b)
	}
	http.Handle("/user/updateAndGetToken",http.HandlerFunc(userUpdateHandler))

}
func channelsBinding()  {
	//channels list handler
	channelsListHandler := func(w http.ResponseWriter, r *http.Request) {
		b := controller.GetAllChannels()
		writeJsonToResponse(&w,b)
	}
	http.Handle("/channels/list",http.HandlerFunc(channelsListHandler))

	alreadyExistsChannelNameHandler := func(w http.ResponseWriter, r *http.Request) {
		channelName  := r.URL.Query().Get("name");
		b := controller.AlreadyExistsChannelName(channelName)
		byteValue := []byte(strconv.Itoa(b))
		writeJsonToResponse(&w,byteValue)
	}
	http.Handle("/channels/alreadyExistsChannelName",http.HandlerFunc(alreadyExistsChannelNameHandler))

	saveNewChannel := func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var channel channelDao.Channel
		err := decoder.Decode(&channel)
		if(err != nil){
			log.Printf("%v",err)
		}
		b := controller.SaveNewChannel(channel.Name,channel.Title,channel.Description,channel.ChannelType)
		byteValue := []byte(strconv.Itoa(b))
		writeJsonToResponse(&w,byteValue)
	}
	http.Handle("/channels/saveNewChannel",http.HandlerFunc(saveNewChannel))
}
func writeJsonToResponse(w *http.ResponseWriter,b []byte){
	(*w).Header().Set("Content-Type","application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Auth-Token")
	(*w).Write(b)
}
