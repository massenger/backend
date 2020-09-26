package services

import (
	"log"
)

//Get ...
func Get() string {
	if true {
		log.Println("files\\get\\success")
		return "success!"
	}
	return "invalid"

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
