package userDao

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../dao"
	"../../utils"

)


type User struct {
	MobileNumber string
	Email string
	FirstName string
	LastName string
	RegisterTimestamp int64
}
const USERS ="user"

func Upsert(mobileNumber string,email string,firstName string,lastName string) (error){
	c:= getCollection()
	user:= User{
		MobileNumber:mobileNumber,
		Email:email,
		FirstName:firstName,
		LastName:lastName,
		RegisterTimestamp: utils.CurrentMilis(),
	}
	_,err := c.Upsert(bson.M{"mobilenumber": mobileNumber},user)
	return err
}

func UpdateToken(mobileNumber string, token string) error{
	c:= getCollection()
	return c.Update(bson.M{"mobilenumber":mobileNumber},bson.M{"$set": bson.M{"token": token}})
}
func getCollection() (*mgo.Collection)  {
	return dao.Session.DB(dao.DATABASE_NAME).C(USERS);
}