package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

// match all file types uploaded to the server for the single page app, all other requests would be webpage loads i.e /About | /blog/post/123
var pattern = regexp.MustCompile("(.js|.css|.map|.html)")
var dirRoot = ""

func main() {
	// Get Current Directory to forward to file server later
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dirRoot = pwd
	http.HandleFunc("/", route)
	log.Fatal(http.ListenAndServe(":8002", nil))

}

func route(w http.ResponseWriter, r *http.Request) {
	if !pattern.Match([]byte(r.URL.Path)) {
		// Route webpages to use the homepage url. i.e if coming from external web link to a specific webpage
		// to automatically still load the single page app which will then read the url and respond accordingly.
		r.URL.Path = "/"
	}
	http.FileServer(http.Dir(dirRoot+"/www")).ServeHTTP(w, r)
}
