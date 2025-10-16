package store

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *UserStore) GetByID(ctx context.Context, userID uint) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var user User
	err := s.db.WithContext(ctx).
		Select("id", "username", "email", "avatar", "created_at").
		First(&user, userID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) VerifyUser(ctx context.Context, email string, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var user User
	err := s.db.WithContext(ctx).
		Where("email = ?", email).
		Select("id", "username", "email", "avatar", "password", "created_at").
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	log.Println("ASdfASD", user.Password)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if s.db == nil {
		return errors.New("database connection is nil")
	}

	err := s.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate key value") {
			if strings.Contains(err.Error(), "users_email_key") {
				return ErrDuplicateEmail
			}
			if strings.Contains(err.Error(), "users_username_key") {
				return ErrDuplicateUsername
			}
		}
		return err
	}

	return nil
}
