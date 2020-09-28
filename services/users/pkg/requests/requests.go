package requests

import "gorm.io/gorm"

//Request is formatting for the user service
type Request struct {
	gorm.Model
	UserName     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
}

//Create ...
func Create() Request {
	return Request{}
}

//CreatePointer ...
func CreatePointer() *Request {
	return &Request{}
}
