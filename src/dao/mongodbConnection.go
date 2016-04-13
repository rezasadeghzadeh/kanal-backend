package dao

import (
	"gopkg.in/mgo.v2"
    "flag"
    "fmt"
)
var Session *mgo.Session
var err error

var  DATABASE_NAME = ""

func init()  {
    f := flag.String("config", "config.json", "config file")
    flag.Parse()
    c, err := NewConfigFile(*f)
    if err != nil {
        fmt.Println("failed to config file, started with default values:", err)
    }

    DATABASE_NAME = c.DatabaseName

	Session, err = mgo.Dial(c.DatabaseServer)
	if err != nil{
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
}
