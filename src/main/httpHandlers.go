package main

import (
	"log"
	"net/http"
	"../controller/channelsController"
	"../controller/channelPostController"
	"../controller/smsVerificationController"
	"strconv"
	"encoding/json"
	"../dao/channelDao"
	"../controller/userController"
	"../controller/uploadController"
	"../utils"
	"../dao/userDao"
)

func InitHttpHandlers() {
	log.Printf("Start Http Initalization\n")
	channelsBinding()
	channelPostBinding()
	userBinding()
	smsVerificationBinding()
	fileUploadBinding()
	serveImagesBinding()
}



type response struct {
	status string
	data string
}

func channelPostBinding()  {
	http.Handle("/channelPosts/postTextMessage",http.HandlerFunc(channelPostController.SaveTextPost ))
	http.Handle("/channelPosts",http.HandlerFunc(channelPostController.GetTextPostsByChannelId ))
}
func serveImagesBinding()  {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving file %s",r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:])
	})
}
func fileUploadBinding() {
	fileUploadHandler  := func(w http.ResponseWriter,r *http.Request){
		res,_ := uploadController.HandleUpload(r)
		json,_ := json.Marshal(res)
		w.Write(json)
	}
	http.Handle("/file/upload",http.HandlerFunc(fileUploadHandler))
}

func smsVerificationBinding()  {
	sendSmsVerificationHandler := func(w http.ResponseWriter, r *http.Request) {
		mobileNumber := r.URL.Query().Get("mobileNumber")
		log.Printf("Mobilenumber : %s",mobileNumber)
		b,_ := json.Marshal(strconv.Itoa(smsVerificationController.SendSmsVerification(mobileNumber)))
		utils.WriteJsonToResponse(&w,b)
	}
	http.Handle("/smsVerification/send",http.HandlerFunc(sendSmsVerificationHandler))

	verifySentSmsHandler := func(w http.ResponseWriter, r *http.Request) {
		mobileNumber := r.URL.Query().Get("mobileNumber")
		enteredText  := r.URL.Query().Get("enteredText")
		b := []byte(strconv.Itoa(smsVerificationController.VerifySentSms(mobileNumber,enteredText)))
		utils.WriteJsonToResponse(&w,b)
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
		utils.WriteJsonToResponse(&w,b)
	}
	http.Handle("/user/updateAndGetToken",http.HandlerFunc(userUpdateHandler))

}
func channelsBinding()  {
	//channels list handler
	channelsListHandler := func(w http.ResponseWriter, r *http.Request) {
		b := controller.GetAllChannels()
		utils.WriteJsonToResponse(&w,b)
	}
	http.Handle("/channels/list",http.HandlerFunc(channelsListHandler))

	alreadyExistsChannelNameHandler := func(w http.ResponseWriter, r *http.Request) {
		channelName  := r.URL.Query().Get("name");
		b := controller.AlreadyExistsChannelName(channelName)
		byteValue := []byte(strconv.Itoa(b))
		utils.WriteJsonToResponse(&w,byteValue)
	}
	http.Handle("/channels/alreadyExistsChannelName",http.HandlerFunc(alreadyExistsChannelNameHandler))

	saveNewChannel := func(w http.ResponseWriter, r *http.Request) {
		token := utils.GetUserTokenFromReq(r)
		userId := userDao.GetUserIdByToken(token)
		log.Printf("User Token: %s UserId:%s",token,userId)

		decoder := json.NewDecoder(r.Body)
		var channel channelDao.Channel
		err := decoder.Decode(&channel)
		if(err != nil){
			log.Printf("%v",err)
		}
		b := controller.SaveNewChannel(channel.Name,channel.Title,channel.Description,channel.ChannelType,channel.ImageUrl,userId)
		byteValue := []byte(strconv.Itoa(b))
		utils.WriteJsonToResponse(&w,byteValue)
	}
	http.Handle("/channels/saveNewChannel",http.HandlerFunc(saveNewChannel))
}

