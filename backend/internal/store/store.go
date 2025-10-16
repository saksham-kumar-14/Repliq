package store

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound          = errors.New("no document found")
	ErrAlreadyExists     = errors.New("resource already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicateUsername = errors.New("username already exists")

	QueryTimeout = time.Second * 5
)

type Storage struct {
	User interface {
		Create(context.Context, *User) error
		GetByID(context.Context, uint) (*User, error)
		VerifyUser(context.Context, string, string) (*User, error)
	}
}

func NewDbStorage(db *gorm.DB) Storage {
	return Storage{
		User: &UserStore{db: db},
	}
}
