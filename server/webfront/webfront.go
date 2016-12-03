package main

import (
	"log"
	"net/http"
	"regexp"
)

// match all file types uploaded to the server for the single page app, all other requests would be webpage loads i.e /About | /blog/post/123
var pattern = regexp.MustCompile("(.js|.css|.map|.html)")

func main() {

	http.HandleFunc("/", route)
	log.Fatal(http.ListenAndServe(":8002", nil))

}

func route(w http.ResponseWriter, r *http.Request) {
	if !pattern.Match([]byte(r.URL.Path)) {
		// Route webpages to use the homepage url. i.e if coming from external web link to a specific webpage
		// to automatically still load the single page app which will then read the url and respond accordingly.
		r.URL.Path = "/"
	}
	http.FileServer(http.Dir("/go/src/zanven42/webfront/www")).ServeHTTP(w, r)
}
