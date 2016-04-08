package controller

import (
	"../../dao/channelDao"
	"encoding/json"
	"log"
)

func GetAllChannels() ([]byte){
	content,error :=json.MarshalIndent(channelDao.GetList(),"","    ")
	if(error != nil){
		log.Printf("%v",error)
	}
	return content
}

func AlreadyExistsChannelName(channelName string) int  {
	return channelDao.AlreadyExistsChannelName(channelName);
}

func SaveNewChannel(channelName string, channelTitle string, channelDesc string, channelType string) (int) {
	return channelDao.SaveNewChannel(channelName, channelTitle, channelDesc, channelType)
}


