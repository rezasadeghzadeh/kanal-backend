package userDao

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../dao"
	"../../utils"

	"log"
)


type User struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	MobileNumber string
	Email string
	FirstName string
	LastName string
	RegisterTimestamp int64
	Token string
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
	return c.Update(bson.M{"mobilenumber":mobileNumber},bson.M{"$set": bson.M{"token":mobileNumber + "."+ token}})
}
func getCollection() (*mgo.Collection)  {
	return dao.Session.DB(dao.DATABASE_NAME).C(USERS);
}

func GetUserIdByToken(token string) (string){
	c:= getCollection()
	user :=User{}
	err := c.Find(bson.M{"token":token}).One(&user)
	if(err != nil){
		log.Printf("Error in searching user by token %s",token)
		log.Printf("userDao:GetUserIdByToken %v",err)
		return ""
	}else{
		log.Printf("userDao: User founded, Hex-Id:%s ",user.Id.Hex())
		return user.Id.Hex()
	}
}