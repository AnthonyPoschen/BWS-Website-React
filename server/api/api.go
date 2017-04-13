package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"bytes"

	"io/ioutil"
)

func main() {
	
	http.HandleFunc("/api/addblog", addBlog)
	http.HandleFunc("/api/editblog", editBlog)
	http.HandleFunc("/api/deleteblog", deleteBlog)
	http.HandleFunc("/api/getblog", getBlog)
	http.HandleFunc("/api/getbloglist", getBlogList)
	http.HandleFunc("/api", apiOverview)
	http.HandleFunc("/api/", apiOverview)

	log.Fatal(http.ListenAndServe(":8001", nil))
	log.Println("Shutting Down API Server")
}

func apiOverview(w http.ResponseWriter, r *http.Request) {
	// return the api layout (maybe viaswagger instead of manually handling this crap)
	w.Write([]byte("Hit Backend with | " + r.URL.Path))
	log.Println("| Request responded")
}

// addBlog is a api endpoint to add a new blog to the database.
func addBlog(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Build request to talk to database
	req, err := http.NewRequest("POST", "http://localhost:8010/addblog", bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	req.Header = r.Header

	client := http.Client{Timeout: time.Second * 5}
	// send request (should add a timeout)
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	w.WriteHeader(res.StatusCode)
	resbody, err := ioutil.ReadAll(res.Body)
	w.Write(resbody)
}

// editblog updates a pre exsisting blog with new data.
func editBlog(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	req, err := http.NewRequest("POST", "http://localhost:8010/editblog", bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	req.Header.Set("token", r.Header.Get("token"))
	if req.Header.Get("token") == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
	}
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Cant Read Response:", err)
		return
	}
	w.WriteHeader(res.StatusCode)
	w.Write(resbody)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read Incoming Message"))
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8010/deleteblog", bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	req.Header.Set("token", r.Header.Get("token"))
	if req.Header.Get("token") == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No 'token' header found"))
		return
	}

	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Cant Read Response:", err)
	}
	w.WriteHeader(res.StatusCode)
	w.Write(resbody)
}

func getBlogList(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read Incoming Message"))
		return
	}
	req, err := http.NewRequest("GET", "http://localhost:8010/getbloglist", bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.WriteHeader(res.StatusCode)
	w.Write(resbody)
}

func getBlog(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read Incoming Message"))
		return
	}
	req, err := http.NewRequest("GET", "http://localhost:8010/getblog", bytes.NewBuffer(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	defer res.Body.Close()
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.WriteHeader(res.StatusCode)
	w.Write(resbody)
}
