package requests

//Requests is formatting for the file service
type Requests struct {
	Ok           string   `json:"ok"`
	FileList     []string `json:"filelist"`
	FileName     string   `json:"filename"`
	FileContents string   `json:"filecontents"`
	ID           int      `json:"id"`
}

//Create ...
func Create() Requests {
	return Requests{}
}

//CreatePointer ...
func CreatePointer() *Requests {
	return &Requests{}
}
