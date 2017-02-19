package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/nu7hatch/gouuid"
)

type blog struct {
	Id       string `dynamodbav:"id"`
	Authorid string `dynamodbav:"authorid"`
	Tittle   string `dynamodbav:"tittle"`
	Postdate string `dynamodbav:"postdate"`
	Content  string `dynamodbav:"content"`
}

func main() {

	dynamodbScanTable()
}

func dynamodbScanTable() {

	db, _ := getDBSess()
	res, err := db.Scan(&dynamodb.ScanInput{
		TableName:        aws.String("Blogs"),
		IndexName:        aws.String("authorby-index"),
		FilterExpression: aws.String("authorid = :author"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":author": &dynamodb.AttributeValue{S: aws.String("0")},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Count:", *res.Count)
	if *res.Count == 0 {
		return
	}
	for _, item := range res.Items {
		fmt.Println("----------")
		for k, v := range item {
			fmt.Println("Name:", k, "Value:", v.GoString())
		}
	}

	// now to fetch based off the id
	dynamodbQueryItem(*res.Items[0]["id"].S)

}

func dynamodbQueryItem(id string) {
	db, _ := getDBSess()
	fmt.Println("-------------")
	fmt.Println("ID:", id)
	res, err := db.Query(&dynamodb.QueryInput{
		TableName:              aws.String("Blogs"),
		KeyConditionExpression: aws.String("id = :value"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":value": &dynamodb.AttributeValue{S: &id},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Found ", *res.Count, "Items")
	for _, item := range res.Items {
		fmt.Println("-------------")
		for k, v := range item {
			fmt.Println(k, v.GoString())
		}
	}
}

func dynamodbGetItem() {
	//db, _ := getDBSess()
	//db.GetItem(&dynamodb.GetItemInput{
	//
	//	TableName: "Blogs",
	//})
}

func dynamodbAddItem() {
	db, _ := getDBSess()
	// Gen a Unique ID for the primary key so nothing gets overwritten ever.
	id, _ := uuid.NewV4()
	//Create item in a golang struct
	b := blog{Id: id.String(), Tittle: "My First Blog Post :D", Postdate: time.Now().UTC().Format(time.RFC3339), Authorid: "0", Content: "I did it i published my first post to amazon's dynamodb statically :D"}
	// MARSHEL the struct's public members
	item, err := dynamodbattribute.MarshalMap(b)
	if err != nil {
		fmt.Println("Marshal Error:", err)
		return
	}
	// Put the item in the database (simple :D)
	_, err = db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Blogs"),
		Item:      item,
	})
	if err != nil {
		fmt.Println("DB Put Error:", err)
	}
}

func dynamodbTestTable() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint := "http://localhost:8400/"
	region := "us-east-1"
	db := dynamodb.New(sess, &aws.Config{Endpoint: &endpoint, Region: &region, Credentials: credentials.NewStaticCredentials("123", "123", "")})
	res, err := createBlogsTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Table:")
	fmt.Println(res.TableDescription)
}

func createBlogsTable(db *dynamodb.DynamoDB) (*dynamodb.CreateTableOutput, error) {
	return db.CreateTable(
		&dynamodb.CreateTableInput{
			TableName: aws.String("Blogs"),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("tittle"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("postdate"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("authorid"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("postdate"),
					KeyType:       aws.String("RANGE"),
				},
			},
			LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
				{
					IndexName: aws.String("authorby-index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("id"),
							KeyType:       aws.String("HASH"),
						},
						{
							AttributeName: aws.String("authorid"),
							KeyType:       aws.String("RANGE"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("KEYS_ONLY"),
					},
				},
				{
					IndexName: aws.String("tittle-index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("id"),
							KeyType:       aws.String("HASH"),
						},
						{
							AttributeName: aws.String("tittle"),
							KeyType:       aws.String("RANGE"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("KEYS_ONLY"),
					},
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(2),
			},
		},
	)
}

func getDBSess() (*dynamodb.DynamoDB, error) {
	// Create Sessions
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Connect to Database
	endpoint := "http://localhost:8400/"
	region := "us-east-1"
	return dynamodb.New(sess, &aws.Config{Endpoint: &endpoint, Region: &region, Credentials: credentials.NewStaticCredentials("123", "123", "")}), nil
}
