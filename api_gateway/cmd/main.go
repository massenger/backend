package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("API Gateway started")
	routes := map[string]string{
		"/files":                   ":8001",
		"/websocket-load-balancer": ":8002",
		"/users":                   ":8003",
	}
	router := mux.NewRouter()
	for in, direction := range routes {
		router.HandleFunc(
			in, func(w http.ResponseWriter, r *http.Request) {
				_, err := io.WriteString(w, direct(w, r, direction+"/service"))
				if err != nil {
					io.WriteString(w, "invalid url")
				}
			},
		)
	}
	http.ListenAndServe(":8000", router)
}

func direct(w http.ResponseWriter, r *http.Request, direction string) string {

	resp, err := http.Get("http://0.0.0.0" + direction)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	return bodyString

}
