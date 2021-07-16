package repository

import (
	"context"

	"github.com/QMDAKA/comment-mock/domain/model"
)

type Comment interface {
	GetByID(ctx context.Context, id uint64) (*model.Comment, error)
	Create(ctx context.Context, comment *model.Comment) error
	UpdateContentByID(ctx context.Context, commentID uint64, content string) error
	DeleteByID(ctx context.Context, commentID uint64) error
}
