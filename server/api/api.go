package main

import (
	"log"
	"net/http"

	"time"

	"strconv"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

func main() {

	http.HandleFunc("/api/addblog", addBlog)
	http.HandleFunc("/api/editblog", editBlog)
	http.HandleFunc("/api/deleteblog", deleteBlog)
	http.HandleFunc("/api/getblog", getBlog)
	http.HandleFunc("/api/getbloglist", getBlogList)
	http.HandleFunc("/api", apiOverview)
	http.HandleFunc("/api/", apiOverview)

	mainctx := context.Background()
	dsc, err := datastore.NewClient(mainctx, "brainwave-studios")
	defer dsc.Close()
	if err != nil {
		// Wait 10 seconds just incase the service is only just coming up now.
		<-time.After(time.Second * 10)
		dsc, err = datastore.NewClient(mainctx, "brainwave-studios")
		if err != nil {
			log.Fatalln("Failed to Connect to datastore:", err)
		}
	}
	log.Fatal(http.ListenAndServe(":8001", nil))
	log.Println("Shutting Down API Server")
}

func apiOverview(w http.ResponseWriter, r *http.Request) {
	// return the api layout (maybe viaswagger instead of manually handling this crap)
	w.Write([]byte("Hit Backend"))
	log.Println("| Request responded")
}

// addBlog is a api endpoint to add a new blog to the database.
// TODO: AUTH req comes from an admin
func addBlog(w http.ResponseWriter, r *http.Request) {
	// Adds a blog to the database
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	dsc, err := datastore.NewClient(ctx, "brainwave-studios")
	if err != nil {
		log.Println("Failed to create DataStore Client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer dsc.Close()
	var b blog
	var id int
	_, err = dsc.Run(ctx, datastore.NewQuery(blogType).Ancestor(rootBlogKey()).Order("-ID").Limit(1)).Next(&b)
	// if err is something other than nothing found
	if err != nil {
		id = 0

	} else {
		id = b.ID + 1
	}

	err = decodeJSON(&b, r.Body)

	if err != nil {
		log.Println("Failed to recieve a valid blog on decode:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// Have the new blog obtained now to set its new ID and dates and push it.
	tn := time.Now()
	b.CreationDate = tn
	if b.Published == true {
		b.PublishedDate = tn
	}
	b.ID = id

	// post the blog to the blogs.
	_, err = dsc.Put(ctx, datastore.IncompleteKey(blogType, rootBlogKey()), &b)
	if err != nil {
		log.Println("Failed to store new blog, Err:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println("New Blog Added with ID:", id)
}

func editBlog(w http.ResponseWriter, r *http.Request) {
	// edits a existing blog in the database
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	dsc, err := datastore.NewClient(ctx, "brainwave-studios")
	if err != nil {
		log.Println("Failed to create DataStore Client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var b blog
	err = decodeJSON(&b, r.Body)
	if err != nil {
		log.Println("Error Decoding Incoming editBlog Request", err)
	}

	key, err := dsc.Run(ctx, datastore.NewQuery(blogType).Ancestor(rootBlogKey()).Filter("ID =", b.ID).KeysOnly()).Next(nil)
	if err != nil {
		// somehow cant find the blog we are magically editing :/
	}
	// put the edited blog in the datastore
	dsc.Put(ctx, key, &b)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	// deletes a blog in the database
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	dsc, err := datastore.NewClient(ctx, "brainwave-studios")
	if err != nil {
		log.Println("Failed to create DataStore Client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var b blog
	err = decodeJSON(&b, r.Body)
	if err != nil {
		log.Println("Error Decoding Incoming deleteBlog Request", err)
	}
	key, err := dsc.Run(ctx, datastore.NewQuery(blogType).Ancestor(rootBlogKey()).Filter("ID =", b.ID).KeysOnly()).Next(nil)
	if err != nil {
		log.Println("Cant find Blog to Delete, Err:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// delete the matching blog
	dsc.Delete(ctx, key)
}

func getBlogList(w http.ResponseWriter, r *http.Request) {
	// gets a list of blogs within a range and returns the titles and hyperlinks
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	dsc, err := datastore.NewClient(ctx, "brainwave-studios")
	if err != nil {
		log.Println("Failed to create DataStore Client")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i := dsc.Run(ctx, datastore.NewQuery(blogType).Ancestor(rootBlogKey()).Order("-ID"))
	var blogs []blog
	counter := 0
	for {
		counter++
		var b blog
		_, err = i.Next(&b)

		if err != nil {
			break
		}
		if b.Published == false {
			continue
		}
		blogs = append(blogs, b)
	}

	var bloglist []blogDesc
	for _, b := range blogs {
		bloglist = append(bloglist, blogDesc{ID: b.ID, Author: b.Author, Tittle: b.Tittle, PublishedDate: b.PublishedDate})
	}

	jsonEncode(bloglist, w)
}

func getBlog(w http.ResponseWriter, r *http.Request) {
	// returns the details of a specific blog.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url variable 'id' missing"))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	dsc, err := datastore.NewClient(ctx, "brainwave-studios")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error Connecting to DataStore:", err)
		return
	}
	var b blog
	_, err = dsc.Run(ctx, datastore.NewQuery(blogType).Ancestor(rootBlogKey()).Filter("ID =", id)).Next(&b)
	if err != nil {
		log.Printf("Failed to fetch Blog with ID %v: %v \n", id, err)
		return
	}
	jsonEncode(b, w)
}
