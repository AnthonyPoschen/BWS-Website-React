package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Information for DynamoDB table data
const (
	// TableBlog Identifier for DynamoDB Table
	TableBlog = "blog"

	AttributeBlogID        = "ID"
	AttributeBlogTittle    = "Tittle"
	AttributeBlogContent   = "Content"
	AttributeBlogAuthor    = "AuthorID"
	AttributeBlogCategory  = "CategoryID"
	AttributeBlogPubDate   = "PublishDate"
	AttributeBlogPublished = "Published"

	SIndexBlogCategory = "category-index"
	SIndexBlogAuthor   = "author-index"

	TableCategory = "Category"

	AttributeCategoryID       = "ID"
	AttributeCategoryName     = "Name"
	AttributeCategoryPriority = "Priority"
	AttributeCategoryParentID = "ParentID"
)

func startup(dev bool) {
	db, err := getdbSession(dev)
	if err != nil {
		fmt.Println("Failed to get DB Session", err)
		return
	}
	tables, err := db.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Failed Listing Tables", err)
		return
	}
	blogtable, categorytable := false, false
	for _, v := range tables.TableNames {
		if v == nil {
			continue
		}
		if *v == TableBlog {
			blogtable = true
		}
		if *v == TableCategory {
			categorytable = true
		}
	}
	if blogtable == false {
		fmt.Println("Creating Blog Tabkle")
		_, err := db.CreateTable(makeBlogTableSchema())
		if err != nil {
			fmt.Println("Error Making Blog Table:", err)
		}
	}
	if categorytable == false {

	}

}

func makeBlogTableSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(TableBlog),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(AttributeBlogID),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String(AttributeBlogAuthor),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String(AttributeBlogCategory),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String(AttributeBlogPubDate),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(AttributeBlogID),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String(AttributeBlogPubDate),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			},
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: aws.String(SIndexBlogCategory),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(AttributeBlogID),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
					{
						AttributeName: aws.String(AttributeBlogCategory),
						KeyType:       aws.String(dynamodb.KeyTypeRange),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeKeysOnly),
				},
			},
			{
				IndexName: aws.String(SIndexBlogAuthor),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(AttributeBlogID),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
					{
						AttributeName: aws.String(AttributeBlogAuthor),
						KeyType:       aws.String(dynamodb.KeyTypeRange),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeKeysOnly),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
}

func makeCategoryTableSchema() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(TableCategory),
	}
}
