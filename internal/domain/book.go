package domain

import (
	"errors"
	"time"
)

var (
	ErrBookNotFound = errors.New("book not found")
)

type Book struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	Rating      int       `json:"rating,omitempty"`
}

type UpdateBookInput struct {
	Title       *string    `json:"title,omitempty"`
	Author      *string    `json:"author,omitempty"`
	PublishDate *time.Time `json:"publish_date,omitempty"`
	Rating      *int       `json:"rating,omitempty"`
}
