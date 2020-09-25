package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/massenger/backend/services/files/pkg/service"
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
