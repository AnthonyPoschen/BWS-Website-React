package main

import (
	"time"
)

const (
	blogType   = "blog"
	authorType = "author"
)

type blog struct {
	ID        int    `json:"id"`
	Tittle    string `json:"tittle"`
	Content   string `json:"content"`
	Author    author `json:"author"`
	Published bool   `json:"published"`

	PublishedDate time.Time `json:"publishdate"`
	CreationDate  time.Time `json:"creationdate"`
	EditedDate    time.Time `json:"editeddate"`
}

type blogDesc struct {
	ID            int       `json:"id"`
	Tittle        string    `json:"tittle"`
	Author        author    `json:"author"`
	PublishedDate time.Time `json:"publishdate"`
}

type author struct {
	Name  string `json:"name"`
	email string `json:"email"`
	ID    string `json:"id"`
}
