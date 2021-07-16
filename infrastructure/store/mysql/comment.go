package mysql

import (
	"context"

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

func (c *Comment) GetByID() {

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
