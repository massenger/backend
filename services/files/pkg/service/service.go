package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/massenger/backend/services/files/pkg/requests"
	"github.com/massenger/backend/services/files/pkg/responses"
)

//Get ...
func Get(r *http.Request) string {
	if true {
		log.Println("files\\delete\\success")
		return "success!"
	}
	return "invalid"
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
		log.Println("couldn't decode json")
		log.Println(err)
		spew.Dump(request)
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
