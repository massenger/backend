package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/massenger/backend/services/files/pkg/service"
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
				fileService(method, w, r, mux.Vars(r))
			}).Methods(method)
		}(method)
	}
	panic(http.ListenAndServe("localhost:8001", fileRouter))
}

func fileService(method string, w http.ResponseWriter, r *http.Request, urlParams map[string]string) {

	if method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, service.Get(r))
	} else if method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, service.Post())
	} else if method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, service.Put())
	} else if method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, service.Delete())
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, service.InvalidMethod(method))
	}
}
