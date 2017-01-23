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
	resp, err := http.Get("http://api:8001/api")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(data)
}

func webpageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://webfront:8002" + r.URL.Path)
	if err != nil {
		log.Println(err)
		// return a 404 later?
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(data)
}
