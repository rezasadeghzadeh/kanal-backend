package uploadController

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"log"
	"../../utils"
)

type Response struct {
	FileName string
}
func HandleUpload(r *http.Request) (Response,error){
	//copy file to tmp directory
	res := Response{}
	r.ParseMultipartForm(32 << 20)
	log.Printf("%v",r)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return res,err
	}
	defer file.Close()
	log.Printf("Processing uploaded file : %s",handler.Filename)
	fileExt  := utils.GetFileExtention(handler.Filename)
	userToken := utils.GetUserTokenFromReq(r)
	newFileName :=  userToken + "_" + utils.GetRandomFileName(10) + fileExt
	f, err := os.OpenFile("/home/reza/kanal/kanal-backend/src/main/static/" + newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return res,err
	}
	defer f.Close()
	io.Copy(f, file)
	res.FileName=newFileName
	log.Printf("HandleUpload newFileName: %v",res)

	return res,nil
}
