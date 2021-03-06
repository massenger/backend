package main

import (
	"bytes"
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

var routes = map[string]string{
	"/files": ":8001",
	"/users": ":8002",
}

func main() {
	log.Println("API Gateway started")
	router := mux.NewRouter().StrictSlash(false)
	var methods = []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodDelete,
	}
	for in, direction := range routes {

		for _, method := range methods {
			in, direction := in, direction

			func(method string) {
				router.HandleFunc(in, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("Access-Control-Allow-Origin", "*")
					direction := direction
					direction += r.URL.String()
					log.Println("Gateway [", in, "] to [", direction, "]")
					io.WriteString(w, direct(w, r, direction, method))
				}).Methods(method)
			}(method)
		}
	}
	http.ListenAndServe(":8000", router)
}

func direct(w http.ResponseWriter, r *http.Request, direction, method string) string {
	jsonStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("im working here lol")
	}
	log.Println("http://localhost" + direction)
	req, err := http.NewRequest(method, "http://localhost"+direction, bytes.NewBuffer(jsonStr))
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
