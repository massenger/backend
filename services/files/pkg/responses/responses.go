package responses

//Response is formatting for the file service
type Response struct {
	Ok           bool     `json:"ok"`
	FileList     []string `json:"filelist"`
	FileName     string   `json:"filename"`
	FileContents string   `json:"filecontents"`
	ID           int      `json:"id"`
}

//New ...
func New() Response {
	return Response{}
}

//NewPointer ...
func NewPointer() *Response {
	return &Response{}
}
