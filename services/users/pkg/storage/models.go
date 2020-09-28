package storage

import "gorm.io/gorm"

//User is the dba struct for files
type User struct {
	gorm.Model
	UserName     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
}
