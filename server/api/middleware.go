package main

import (
	"io"

	"encoding/json"

	"cloud.google.com/go/datastore"
)

func isAdmin() {

}

func rootBlogKey() *datastore.Key {
	return datastore.NameKey("blog", "BlogRoot", nil)
}

func decodeJSON(t interface{}, req io.Reader) error {
	decoder := json.NewDecoder(req)
	return decoder.Decode(&t)
}

func jsonEncode(d interface{}, w io.Writer) {
	encoder := json.NewEncoder(w)
	encoder.Encode(d)
}
