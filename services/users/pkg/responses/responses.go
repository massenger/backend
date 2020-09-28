package responses

import "gorm.io/gorm"

//Response is formatting for the file service
type Response struct {
	gorm.Model
	UserName     `json:"username"`
	PasswordHash `json:"passwordhash"`
}

//Create ...
func Create() Response {
	return Requests{}
}

//CreatePointer ...
func CreatePointer() *Response {
	return &Requests{}
}
