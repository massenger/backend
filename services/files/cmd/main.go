package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/massenger/backend/services/files/pkg/service"
)

func main() {
	log.Println("File Service stated @ 8001")
	fileRouter := mux.NewRouter()
	methods := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodGet,
	}
	for _, method := range methods {
		func(method string) {
			fileRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fileService(method, w, r)
			}).Methods(method)
		}(method)
	}
	panic(http.ListenAndServe("localhost:8001", fileRouter))
}

func fileService(method string, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch method {
	case http.MethodGet:
		io.WriteString(w, service.Get(r, r.URL.Query().Get("id")))
	case http.MethodPost:
		io.WriteString(w, service.Post(r))
	case http.MethodPut:
		io.WriteString(w, service.Put())
	case http.MethodDelete:
		io.WriteString(w, service.Delete())
	default:
		io.WriteString(w, service.InvalidMethod(method))
	}

}
