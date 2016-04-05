package controller

import (
	"../dao"
	"encoding/json"
	"log"
)

func GetAllChannels() ([]byte){
	content,error :=json.MarshalIndent(dao.GetList(),"","    ")
	if(error != nil){
		log.Printf("%v",error)
	}
	return content
}

func AlreadyExistsChannelName(channelName string) int  {
	return dao.AlreadyExistsChannelName(channelName);
}

func SaveNewChannel(channelName string, channelTitle string, channelDesc string, channelType string) (int) {
	return dao.SaveNewChannel(channelName, channelTitle, channelDesc, channelType)
}


