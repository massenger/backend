package requests

//requests is formatting for the file service
type Requests struct {
	Ok           string   `json:"ok"`
	FileList     []string `json:"filelist"`
	FileName     string   `json:"filename"`
	FileContents string   `json:"filecontents"`
	ID           int      `json:"id"`
}

//New ...
func Create() {
	return Requests{}
}

//CreatePointer ...
func CreatePointer() {
	return &Requests{}
}
