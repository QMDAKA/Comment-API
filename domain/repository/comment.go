package repository

import (
	"context"

	"github.com/QMDAKA/comment-mock/domain/model"
)

type Comment interface {
	GetByID()
	Create(ctx context.Context, comment *model.Comment) error
}
