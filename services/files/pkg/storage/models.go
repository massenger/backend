package storage

import "gorm.io/gorm"

//File is the dba struct for files
type File struct {
	ID int `json:"id"`
	gorm.Model
	Owner        string
	Permission   int    `json:"permission"` // 0 unprotected, 1 anyone in org, 2 special only, 3 owner only
	FileName     string `json:"filename"`
	FileContents string `json:"filecontents"`
}
