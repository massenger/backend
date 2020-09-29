package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/massenger/backend/services/users/pkg/service"
	storage "github.com/massenger/backend/services/users/pkg/storage"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = storage.ConnectDB("zane", "52455245", "localhost", "massenger")
	log.Println("Users Service stated @ 8002")
	fileRouter := mux.NewRouter().StrictSlash(false)
	methods := []string{
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodGet,
	}
	for _, method := range methods {
		func(method string) {
			fileRouter.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
				userService(method, w, r)
			}).Methods(method)
		}(method)
	}
	panic(http.ListenAndServe("localhost:8002", fileRouter))
}

func userService(method string, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch method {
	case http.MethodGet:
		io.WriteString(w, service.Get(r, r.URL.Query().Get("id"), db))
	case http.MethodPost:
		io.WriteString(w, service.Post(r, db))
	case http.MethodPut:
		io.WriteString(w, service.Put())
	case http.MethodDelete:
		io.WriteString(w, service.Delete())
	default:
		io.WriteString(w, service.InvalidMethod(method))
	}

}
