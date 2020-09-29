package responses

import "gorm.io/gorm"

//Response is formatting for the file service
type Response struct {
	gorm.Model
	Ok           bool   `json:"ok"`
	UserName     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
}

//Create ...
func Create() Response {
	return Response{}
}

//CreatePointer ...
func CreatePointer() *Response {
	return &Response{}
}
