package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", apiOverview)
	http.HandleFunc("/api/", apiOverview)
	// LogIn ??
	// Logout ??

	// AddBlog
	// EditBlog
	// DeleteBlog
	// GetBlogList
	// GetBlog
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func apiOverview(w http.ResponseWriter, r *http.Request) {
	// return the api layout (maybe viaswagger instead of manually handling this crap)
	w.Write([]byte("Hit Backend"))
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	// Adds a blog to the database
}

func editBlog(w http.ResponseWriter, r *http.Request) {
	// edits a existing blog in the database
}

func delteBlog(w http.ResponseWriter, r *http.Request) {
	// deletes a blog in the database
}

func getBlogList(w http.ResponseWriter, r *http.Request) {
	// gets a list of blogs within a range and returns the titles and hyperlinks
}

func getBlog(w http.ResponseWriter, r *http.Request) {
	// returns the details of a specific blog.
}
