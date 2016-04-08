package userController

import (
	"../../dao/userDao"
	"log"
	"../../utils"
)

func Upsert(mobileNumber string,email string,firstName string,lastName string) (int){
	err := userDao.Upsert(mobileNumber,email,firstName,lastName)
	if(err != nil){
		log.Printf("%v",err)
		return 0
	}else {
		return 1
	}
}

func UpsertAndGetToken(mobileNumber string,email string,firstName string,lastName string) (string){
	result := Upsert(mobileNumber,email,firstName,lastName)
	token:="0"
	if(result == 1){
		token = utils.GetToken()
		userDao.UpdateToken(mobileNumber,token)
	}
	return token
}
