package store

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CommentStore struct {
	db *gorm.DB
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ParentId  int       `json:"parent_id"`
	Text      string    `json:"text"`
	Upvotes   int       `json:"upvotes"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *CommentStore) GetAll(ctx context.Context) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var comments []*Comment
	err := s.db.WithContext(ctx).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentStore) GetByID(ctx context.Context, commentID uint) (*Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var comment Comment
	err := s.db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &comment, nil
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if s.db == nil {
		return errors.New("database connection is nil")
	}

	err := s.db.WithContext(ctx).Create(comment).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) DeleteByID(ctx context.Context, commentID uint) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	err := s.db.WithContext(ctx).Delete(&Comment{}, commentID).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) PatchByID(ctx context.Context, commentID uint, updates map[string]interface{}) (*Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var comment Comment
	err := s.db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	err = s.db.WithContext(ctx).Model(&comment).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
