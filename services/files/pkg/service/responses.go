package responses

//Response is formatting for the file service
type Response struct {
	Ok           bool     `json:"ok"`
	FileList     []string `json:"filelist"`
	FileName     string   `json:"filename"`
	FileContents string   `json:"filecontents"`
}

//New ...
func New() {
	return Response{}
}

//New ...
func NewPointer() {
	return &Response{}
}
