package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var bLocalDev = false
var api = "localhost"
var webfront = "localhost"

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
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8001"+r.URL.String(), bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Request Error:", err.Error())
	}
	req.Header = r.Header
	client := http.Client{Timeout: time.Second * 5}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
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
	if err != nil {
		log.Println(err)
		// return a 404 later?
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(data)
}
