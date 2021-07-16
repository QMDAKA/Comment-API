package repository

import (
	"context"

	"github.com/QMDAKA/comment-mock/domain/model"
)

type User interface {
	GetByUUID(ctx context.Context, uuid string) (*model.User, error)
}
