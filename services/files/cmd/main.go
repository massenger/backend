package main

import (
	"log"
	"net/http"

	"../pkg/service"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("File Service stated @ 8001")
	globalRouter := mux.NewRouter()
	serviceRouter := globalRouter.PathPrefix("/service").Subrouter()
	fileRouter := serviceRouter.PathPrefix("/files").Subrouter()
	fileRouter.HandleFunc("", fileService)
}

func fileService(w http.ResponseWriter, r *http.Request) {
	service.Get()
}
