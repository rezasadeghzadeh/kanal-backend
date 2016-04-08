package smsVerificationController

import (
	"math/rand"
	"../../dao/smsVerificationDao"
	"log"
	"strconv"
)
func SendSmsVerification(mobileNumber string)  (int){
	smsNumber  := rand.Intn(99999)
	//send by sms

	//save in db
	err:= smsVerificationDao.Upsert(mobileNumber,smsNumber)
	if err==nil{
		return 1
	}else {
		return 0
	}
}

func VerifySentSms(mobileNumber string,enteredText string) (int){
	smsVerification,err := smsVerificationDao.FindByMobileNumber(mobileNumber)
	log.Printf("Veryfing for mobile %s and enteredText %s",mobileNumber,enteredText)
	if(err != nil){
		log.Printf("sms verification: %v",err)
		return 0
	}
	enteredNumberic,err  := strconv.Atoi(enteredText)
	if(err != nil){
		log.Printf("sms verification: %v",err)
		return 0
	}

	if (smsVerification.SmsNumber == enteredNumberic){
		log.Printf("sms verified successfully,: %s %s",mobileNumber,enteredText)
		return  1
	}
	return 0
}



