package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// match all file types uploaded to the server for the single page app, all other requests would be webpage loads i.e /About | /blog/post/123
var pattern = regexp.MustCompile("(.js|.css|.map|.html)")
var dirRoot = ""
var dirExtention *string
var dir string

func main() {

	dirExtention = flag.String("ext", "/www", "If present this relative directory will be used instead of './www' for the location of the webfiles")
	flag.Parse()
	// Get Current Directory to forward to file server later
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dirRoot = pwd
	fmt.Println(dirRoot)
	dir, _ = filepath.Abs(*dirExtention)
	fmt.Println("Dir:", dir)
	http.HandleFunc("/", route)
	log.Fatal(http.ListenAndServe(":8002", nil))
}

func route(w http.ResponseWriter, r *http.Request) {
	if !pattern.Match([]byte(r.URL.Path)) {
		// Route webpages to use the homepage url. i.e if coming from external web link to a specific webpage
		// to automatically still load the single page app which will then read the url and respond accordingly.
		r.URL.Path = "/"
	}
	//http.FileServer(http.Dir(path.Join(dirRoot, *dirExtention))).ServeHTTP(w, r)
	fmt.Println(dir)
	http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
}
