package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/massenger/backend/services/users/pkg/requests"
	"github.com/massenger/backend/services/users/pkg/responses"
	"github.com/massenger/backend/services/users/pkg/storage"
	"gorm.io/gorm"
)

//Get ...
func Get(r *http.Request, id string, db *gorm.DB) string {
	response := responses.Create()
	idAsInt, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("users\\get\\fail")
		return responseString
	}
	if idAsInt == 0 {
		//list in future
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("users\\get\\fail")
		return responseString
	}
	user := &storage.User{}
	db.First(&user, idAsInt)
	if user.ID == 0 {
		//not found
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("users\\get\\fail")
		return responseString
	}
	response.ID = user.ID
	response.UserName = user.UserName
	response.Ok = true
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("files\\get\\success")
	return responseString

}

//Post ...
func Post(r *http.Request, db *gorm.DB) string {
	//validation to be updated...
	defer r.Body.Close()
	request := requests.CreatePointer()
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)
	err = json.Unmarshal(body, &request)
	response := responses.Create()
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("users\\post\\fail")
		return responseString
	}
	//search request.ID
	//find or dont find

	if valid(request) {
		user := &storage.User{}
		user.UserName = request.UserName
		//more desc to come!
		db.Create(user)
		response.ID = user.ID
		response.Ok = true
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("users\\post\\success")
		return responseString
	}
	response.Ok = false
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("users\\post\\fail")
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
func valid(r *requests.Request) bool {
	retVal := true

	return retVal
}
