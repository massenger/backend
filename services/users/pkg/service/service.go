package services

import (
	"log"
	"net/http"

	"github.com/massenger/backend/services/users/pkg/requests"
	"gorm.io/gorm"
)

//Get ...
func Get(r *http.Request, id string, db *gorm.DB) string {
	if true {
		log.Println("files\\get\\fail")
		return "fail!"
	}
	return ""
}

//Post ...
func Post(r *http.Request, db *gorm.DB) string {
	if true {
		log.Println("files\\post\\fail")
		return "fail!"
	}
	return ""
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
