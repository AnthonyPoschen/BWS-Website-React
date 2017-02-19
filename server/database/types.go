package main

type blog struct {
	ID           string `dynamodbav:"ID" json:"ID"`
	Tittle       string `dynamodbav:"Tittle" json:"Tittle"`
	Content      string `dynamodbav:"Content" json:"Content"`
	AuthorID     string `dynamodbav:"AuthorID" json:"AuthorID"`
	AuthorName   string `dynamodbav:"-" json:"AuthorName"`
	CategoryID   string `dynamodbav:"CategoryID" json:"CategoryID"`
	CategoryName string `dynamodbav:"-" json:"CategoryName"`
	PubDate      string `dynamodbav:"PublishDate" json:"PublishDate"`
	// newPubDate may or may not be used as the origional publishdate is needed to be able to find
	// the blog post as it is part of the key (wasnt really thinking ahead)
	NewPubDate string `dynamodbav:"-" json:"NewPublishDate"`
	Published  bool   `dynamodbav:"Published" json:"Published"`
}

type blogpage struct {
	ID         string `dynamodbav:"ID" json:"ID"`
	Tittle     string `dynamodbav:"Tittle" json:"Tittle"`
	AuthorName string `dynamodbav:"AuthorName" json:"AuthorName"`
}

type category struct {
	ID       string `dynamodbav:"ID" json:"ID"`
	Name     string `dynamodbav:"Name" json:"Name"`
	Priority int    `dynamodbav:"Priority" json:"Priority"`
	ParentID string `dynamodbav:"ParentID" json:"ParentID"`
}

type user struct {
	Username string `dynamodbav:"Username" json:"Username"`
	Email    string `dynamodbav:"Email" json:"Email"`
	Password string `dynamodbav:"Password" json:"Password"`
	Name     string `dynamodbav:"Name" json:"Name"`
}
