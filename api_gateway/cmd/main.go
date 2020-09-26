package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var methods = []string{
	http.MethodGet,
	http.MethodPut,
	http.MethodPost,
	http.MethodDelete,
}

func main() {
	log.Println("API Gateway started")
	routes := map[string]string{
		"/files": ":8001",
	}
	router := mux.NewRouter()
	var methods = []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
	}
	for in, direction := range routes {

		for _, method := range methods {
			func(method string) {
				router.HandleFunc(in, func(w http.ResponseWriter, r *http.Request) {
					log.Println("Gateway [", in, "] to [", direction, "]")
					io.WriteString(w, direct(w, r, direction, method))
				}).Methods(method)
			}(method)
		}
	}
	http.ListenAndServe(":8000", router)
}

func direct(w http.ResponseWriter, r *http.Request, direction, method string) string {

	req, err := http.NewRequest(method, "http://localhost"+direction, nil)
	clnt := &http.Client{}
	resp, err := clnt.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	resp.Body.Close()
	return bodyString

}
