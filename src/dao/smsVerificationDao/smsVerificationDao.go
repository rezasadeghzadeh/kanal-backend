package smsVerificationDao

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../dao"
	"../../utils"
	"log"
)

type SmsVerification struct {
	MobileNumber string
	SmsNumber int
	Verified int
	SentTimestamp int64
}

const SMS_VERIFICATION ="smsVerification"

func Insert(mobileNumber string,smsNumber int) (error){
	smsVerification := SmsVerification{
		MobileNumber:mobileNumber,
		SmsNumber:smsNumber,
		SentTimestamp: utils.CurrentMilis(),
	}
	c := getCollection()
	return c.Insert(smsVerification)
}

func Upsert(mobileNumber string,smsNumber int) (error){
	smsVerification := SmsVerification{
		MobileNumber:mobileNumber,
		SmsNumber:smsNumber,
		SentTimestamp: utils.CurrentMilis(),
	}
	c := getCollection()
	_,err := c.Upsert(bson.M{"mobilenumber":mobileNumber},smsVerification)
	return err
}

func getCollection() (*mgo.Collection) {
	c := dao.Session.DB(dao.DATABASE_NAME).C(SMS_VERIFICATION)

	// Index
	index := mgo.Index{
		Key:        []string{"mobilenumber"},
		Unique:     true,
	}

	err := c.EnsureIndex(index)
	if(err != nil){
		log.Printf("%v",err)
	}

	return c
}

func FindByMobileNumber(mobileNumber string) (*SmsVerification,error) {
	c:= getCollection()
	var verificationSms SmsVerification
	err := c.Find(bson.M{ "mobilenumber" : mobileNumber}).One(&verificationSms)
	log.Printf("Verification sms : %v",verificationSms)
	if(err != nil){
		return nil,err
	}else {
		return &verificationSms,nil
	}
}
