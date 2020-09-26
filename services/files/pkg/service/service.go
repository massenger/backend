package services

import (
	"encoding/json"
	"log"
	"net/http"

	responses "github.com/massenger/backend/services/files/pkg/service/requests"
	responses "github.com/massenger/backend/services/files/pkg/service/responses"
)

//Get ...
func Get(r *http.Request) string {
	request := requests.CreatePointer()
	err = json.NewDecoder(r.Body).Decode(&request)
	response := responses.New()
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		return responseString
	}
	//search request.ID
	//find or dont find
	if false {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		return responseString
	}
	response.FileName = "hello.txt"
	response.FileContents = "asdghbvdse" //open hello.txt and return val
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	return responseString
}

//Post ...
func Post() string {
	if true {
		log.Println("files\\post\\success")
		return "success!"
	}
	return "invalid"

}

//Put ...
func Put() string {
	if true {
		log.Println("files\\put\\success")
		return "success!"
	}
	return "invalid"

}

//Delete ...
func Delete() string {
	if true {
		log.Println("files\\delete\\success")
		return "success!"
	}
	return "invalid"

}

//InvalidMethod ...
func InvalidMethod(method string) string {
	return "Invalid method " + method
}
