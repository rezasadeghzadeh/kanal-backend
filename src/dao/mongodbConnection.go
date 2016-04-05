package dao

import (
	"gopkg.in/mgo.v2"
)
var session *mgo.Session
var err error

const  DATABASE_NAME ="kanal"

func init()  {
	session, err = mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
}
