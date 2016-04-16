package channelPostDao

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"../../dao"
)

const CHANNEL_POST = "channelPost"

type TextPost struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	ChannelId string
	Text string
	RegisterTimestamp time.Time
}

func GetTextPostsByChannelId(channelId string) ([]TextPost)  {
	c:= getCollection()
	var channelPost []TextPost
	c.Find(bson.M{"channelid":channelId}).All(&channelPost)
	return channelPost
}

func SaveTextPost(textPost TextPost) (error){
	c:= getCollection()
	return c.Insert(textPost)
}

func getCollection() (*mgo.Collection) {
	return dao.Session.DB(dao.DATABASE_NAME).C(CHANNEL_POST)
}




