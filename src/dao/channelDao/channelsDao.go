package channelDao

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"log"
	"../../dao"
)

type Channel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Title string
	Description string
	ChannelType string
	OwnerId string
	RegisterTimestamp time.Time
	ImageUrl string
	ThumbnailUrl string
}

const CHANNELS  = "channels"

func GetList() ([]Channel)  {
	c := getCollection()
	var channelsList []Channel
	c.Find(nil).All(&channelsList)
	return channelsList
}

func AlreadyExistsChannelName(channelName string) int{
	c:= getCollection()
	count,err :=  c.Find(bson.M{"name":channelName}).Count()
	if(err != nil){
		log.Printf("%v",err)
		return -1
	}
	return count
}

func getCollection() (*mgo.Collection)  {
	return dao.Session.DB(dao.DATABASE_NAME).C(CHANNELS)
}

func Insert(channel Channel) (error)  {
	c := getCollection()
	return c.Insert(channel)
}

func InsertSample(title string) {
	channel1 := Channel{
		Name:title,
		Title: title,
		ChannelType:"1",
		Description:"Desc " + title,
	}
	Insert(channel1)
}

func SaveNewChannel(channelName string, channelTitle string, channelDesc string, channelType string, channelImageUrl string,userId string) (int){
	if channelName  == "" || channelTitle == "" || userId  == ""{
		return 0;
	}

	newChannel  := Channel{
		Name:channelName,
		Title: channelTitle,
		ChannelType:channelType,
		Description:channelDesc,
		ImageUrl:channelImageUrl,
		ThumbnailUrl : channelImageUrl,
		OwnerId:userId,
	}
	err := Insert(newChannel)
	if(err != nil){
		log.Printf("%v",err)
		return  0;
	}
	return 1;
}







