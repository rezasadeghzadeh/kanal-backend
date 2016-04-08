package channelDao

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"log"
	"../../dao"
)

type Channel struct {
	Name string
	Title string
	Description string
	ChannelType string
	OwnerId string
	RegisterTimestamp time.Time
	Images channelImage
}

type channelImage struct  {
	OriginalPath string
	ThumbnailPath string
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
		Images: channelImage{
			OriginalPath:"http://www.yarancenter.com/sites/default/files/gallery/smod/DSC02797_0.JPG",
			ThumbnailPath:"http://www.yarancenter.com/sites/default/files/gallery/smod/DSC02797_0.JPG",
		},
	}
	Insert(channel1)
}

func SaveNewChannel(channelName string, channelTitle string, channelDesc string, channelType string) (int){
	if channelName  == "" || channelTitle == "" {
		return 0;
	}

	newChannel  := Channel{
		Name:channelName,
		Title: channelTitle,
		ChannelType:channelType,
		Description:channelDesc,
		Images :channelImage{
			OriginalPath:"http://www.yarancenter.com/sites/default/files/gallery/smod/DSC02797_0.JPG",
			ThumbnailPath:"http://www.yarancenter.com/sites/default/files/gallery/smod/DSC02797_0.JPG",
		},
	}
	err := Insert(newChannel)
	if(err != nil){
		log.Printf("%v",err)
		return  0;
	}
	return 1;
}







