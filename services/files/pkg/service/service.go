package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/massenger/backend/services/files/pkg/requests"
	"github.com/massenger/backend/services/files/pkg/responses"
)

//Get ...
func Get(r *http.Request, id string) string {
	idAsInt, err := strconv.ParseInt(id, 0, 64)
	response := responses.New()
	if err != nil {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\get\\fail")
		return responseString
	}
	if idAsInt == 4 {

		response.Ok = true
		response.FileName = "hello4.mp4"
		response.FileContents = "somebody once told me the world was gonna roll me"
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		return responseString
		log.Println("files\\get\\success")
		return responseString
	}
	response.Ok = false
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("files\\get\\fail")
	return responseString
}

//Post ...
func Post(r *http.Request) string {
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
	if false {
		response.Ok = false
		responseByte, _ := json.Marshal(response)
		responseString := string(responseByte)
		log.Println("files\\post\\fail")
		return responseString
	}
	response.FileName = "hello.txt"
	response.Ok = true
	response.FileContents = "asdghbvdse" //open hello.txt and return val
	responseByte, _ := json.Marshal(response)
	responseString := string(responseByte)
	log.Println("files\\post\\success")
	return responseString
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
