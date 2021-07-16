package mysql

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/common/apperr"
	"github.com/QMDAKA/comment-mock/domain/model"
)

type Comment struct {
	db *gorm.DB
}

func ProvideCommentRepo(db *gorm.DB) *Comment {
	return &Comment{db: db}
}

func (c *Comment) GetByID(ctx context.Context, id uint64) (*model.Comment, error) {
	var comment model.Comment
	if err := c.db.WithContext(ctx).First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.NewWithMsg(apperr.NotFound, fmt.Sprintf("record not found: %d", id))
		}
		return nil, apperr.Wrap(apperr.Database, err)
	}
	return &comment, nil
}

func (c *Comment) Create(ctx context.Context, comment *model.Comment) error {
	db, ok := GetTx(ctx)
	if !ok {
		db = c.db
	}
	if err := db.WithContext(ctx).Create(comment).Error; err != nil {
		return apperr.New_(apperr.Database, apperr.OptCltMsg(err.Error()))
	}
	return nil
}

func (c *Comment) UpdateContentByID(ctx context.Context, commentID uint64, content string) error {
	db, ok := GetTx(ctx)
	if !ok {
		db = c.db
	}
	if err := db.WithContext(ctx).Model(model.Comment{}).Where("id = ?", commentID).
		Update("content", content).Error; err != nil {
		return apperr.New_(apperr.Database, apperr.OptCltMsg(err.Error()))
	}
	return nil
}

func (c *Comment) DeleteByID(ctx context.Context, commentID uint64) error {
	db, ok := GetTx(ctx)
	if !ok {
		db = c.db
	}
	if err := db.WithContext(ctx).Delete(&model.Comment{}, commentID).Error; err != nil {
		return apperr.New_(apperr.Database, apperr.OptCltMsg(err.Error()))
	}
	return nil
}
