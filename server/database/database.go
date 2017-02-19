package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"time"

	"encoding/json"

	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/justinas/alice"
)

var dev *bool

func main() {
	// Check tables exsist and if they dont create them.
	dev = flag.Bool("dev", false, "If present the server expects a local hosted dynamodb connection")
	flag.Parse()
	startup(*dev)
	ct := contextTimeout{duration: time.Second * 3, parent: context.Background()}

	common := alice.New(ct.ContextTimeout)
	admin := alice.New(ct.ContextTimeout)

	http.Handle("/", common.ThenFunc(status))
	http.Handle("/addblog", admin.ThenFunc(addBlog))
	http.Handle("/editblog", admin.ThenFunc(editBlog))
	http.Handle("/deleteblog", admin.ThenFunc(deleteBlog))
	http.Handle("/getblog", common.ThenFunc(fetchBlog))
	http.Handle("/getbloglist", common.ThenFunc(fetchBlogPage))

	log.Println(http.ListenAndServe(":8010", nil))
}

func status(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	defer r.Body.Close()
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	db, err := getdbSession(*dev)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error Getting DB Session:", err)
		return
	}
	var b blog
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error Decoding Request:", err)
		return
	}
	u, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error Generating UUID:", err)
		return
	}
	b.ID = u.String()
	b.PubDate = time.Now().UTC().Format(time.RFC3339)
	item, err := dynamodbattribute.MarshalMap(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error Marshaling Request:", err)
	}
	fmt.Println("Item:", b)
	fmt.Println(item)
	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TableBlog),
		Item:      item,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("Error Putting Item into DB:", err)
	}
}

func editBlog(w http.ResponseWriter, r *http.Request) {
	db, err := getdbSession(*dev)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var b blog
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := dynamodbattribute.MarshalMap(b)
	db.Query(&dynamodb.QueryInput{
		TableName: aws.String(TableBlog),
	})
	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TableBlog),
		Item:      item,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	db, err := getdbSession(*dev)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var b blog
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := dynamodbattribute.MarshalMap(b)
	_, err = db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(TableBlog),
		Key:       item,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

// fetchblog requires knowledge of the Key to be able to be fetched.
func fetchBlog(w http.ResponseWriter, r *http.Request) {
	db, err := getdbSession(*dev)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var b blog
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := dynamodbattribute.MarshalMap(b)

	output, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableBlog),
		Key:       item,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = dynamodbattribute.UnmarshalMap(output.Item, &b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func fetchBlogPage(w http.ResponseWriter, r *http.Request) {
	db, err := getdbSession(*dev)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//var blogs []blog
	result, err := db.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(TableBlog),
		FilterExpression: aws.String(AttributeBlogPublished + " = :published"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":published": {
				BOOL: aws.Bool(true),
			},
		},
		ProjectionExpression: aws.String(
			AttributeBlogID + "," +
				AttributeBlogTittle + "," +
				AttributeBlogAuthor + "," +
				AttributeBlogPubDate + "," +
				AttributeBlogCategory),
	})
	if err != nil {
		fmt.Println("error doing scan: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var blogs []blog
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &blogs)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to convert database result to object"))
		return
	}
	//fmt.Println(len(blogs))
	//fmt.Println(blogs)
	data, err := json.Marshal(blogs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Converting data to returnable information: " + err.Error()))
		return
	}
	w.Write(data)
	//dynamodb.ScanInput{}
}

func fetchallAuthors(w http.ResponseWriter, r *http.Request) {

}

func fetchAllCategorys(w http.ResponseWriter, r *http.Request) {

}

func fetchAuthor(w http.ResponseWriter, r *http.Request) {

}

func fetchCategory(w http.ResponseWriter, r *http.Request) {

}
