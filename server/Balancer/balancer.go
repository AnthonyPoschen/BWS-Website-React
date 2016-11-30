package main

import (
	"log"
	"net/http"
)

// Init the balancer and start the server listener.
func main() {

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/", webpageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// forward this to the api backend
	w.Write([]byte("test Backend"))
	resp, err := http.Get("localhost:8001/api")

}

func webpageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test WebPage"))
}
