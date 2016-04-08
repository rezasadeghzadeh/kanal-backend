package dao

import (
	"gopkg.in/mgo.v2"
)
var Session *mgo.Session
var err error

const  DATABASE_NAME ="kanal"

func init()  {
	Session, err = mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
}
