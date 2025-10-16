package store

import "gorm.io/gorm"

type CommentStore struct {
	db gorm.DB
}

type Comment struct {
	gorm.Model
	Id        int
	ParentId  int
	Text      string
	Upvotes   int
	CreatedAt string
	UserId    string
}
