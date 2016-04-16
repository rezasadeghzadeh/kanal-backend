package channelPostController

import (
	"net/http"
	"encoding/json"
	"log"
	"../../dao/channelPostDao"
	"../../utils"
)

func SaveTextPost(w http.ResponseWriter, r * http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var textPost channelPostDao.TextPost
	err := decoder.Decode(&textPost)
	if(err != nil){
		log.Printf("%v",err)
	}
	e := channelPostDao.SaveTextPost(textPost)
	result := 1
	if(e != nil){
		result=0
	}
	byteValue,_ := json.Marshal(result)
	utils.WriteJsonToResponse(&w,byteValue)
}

func GetTextPostsByChannelId(w http.ResponseWriter, r * http.Request)  {
	channelId := r.URL.Query().Get("channelId")
	result := channelPostDao.GetTextPostsByChannelId(channelId)
	byteValue,_ := json.Marshal(result)
	utils.WriteJsonToResponse(&w,byteValue)
}
