package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/massenger/backend/services/files/pkg/requests"
	"github.com/massenger/backend/services/files/pkg/responses"
	"github.com/massenger/backend/services/files/pkg/storage"
	"gorm.io/gorm"
)

//Get ...
func Get(r *http.Request, id string, db *gorm.DB) string {
	idAsInt, err := strconv.ParseInt(id, 0, 64)
	response := responses.New()
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\get\\fail")
		return responseString
	}
	file := &storage.File{}
	db.First(&file, idAsInt) // find product with integer primary key
	if file.ID == 0 {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\get\\fail")
		return responseString
	}
	response.ID = file.ID
	response.FileName = file.FileName
	response.FileContents = file.FileContents
	response.Ok = true
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("files\\get\\success")
	return responseString

}

//Post ...
func Post(r *http.Request, db *gorm.DB) string {
	defer r.Body.Close()
	request := requests.CreatePointer()
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)
	err = json.Unmarshal(body, &request)
	response := responses.New()
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\post\\fail")
		return responseString
	}
	//search request.ID
	//find or dont find

	if valid(request) {
		file := &storage.File{}
		file.FileName = request.FileName
		file.FileContents = request.FileContents
		file.Owner = "anon"
		file.Permission = 0
		db.Create(file)
		response.ID = file.ID
		response.Ok = true
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\post\\success")
		return responseString
	}
	response.Ok = false
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("files\\post\\fail")
	return responseString
}

//Put ...
func Put() string { //Wait until auth
	if true {
		log.Println("files\\put\\fail")
		return "fail!"
	}
	return ""

}

//Delete ...
func Delete() string { //Wait until auth
	if true {
		log.Println("files\\delete\\fail")
		return "fail!"
	}
	return ""

}

//InvalidMethod ...
func InvalidMethod(method string) string {
	return "Invalid method " + method
}

//Need to work on this
func valid(r *requests.Requests) bool {
	retVal := true

	return retVal
}
