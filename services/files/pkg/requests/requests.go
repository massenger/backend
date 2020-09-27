package requests

//Requests is formatting for the file service
type Requests struct {
	FileName     string `json:"filename"`
	FileContents string `json:"filecontents"`
}

//Create ...
func Create() Requests {
	return Requests{}
}

//CreatePointer ...
func CreatePointer() *Requests {
	return &Requests{}
}
