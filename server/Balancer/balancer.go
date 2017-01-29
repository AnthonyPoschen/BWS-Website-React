package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var bLocalDev = false
var api = "api"
var webfront = "webfront"

// Init the balancer and start the server listener.
func main() {

	flag.BoolVar(&bLocalDev, "local", false, "If Present Uses localhost instead of service names")
	flag.Parse()
	if bLocalDev {
		api = "localhost"
		webfront = "localhost"
	}
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/", webpageHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// forward this to the api backend

	resp, err := http.Get("http://" + api + ":8001" + r.URL.String())
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}

func webpageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://" + webfront + ":8002" + r.URL.Path)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		// return a 404 later?
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}
