package main

import (
	"io/ioutil"
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
	w.Write([]byte("test Backend\n"))
	resp, err := http.Get("http://api:8001/api")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(data)
}

func webpageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test WebPage"))
}
