package auth

import (
	"context"

	"github.com/QMDAKA/comment-mock/common/apperr"
	"github.com/QMDAKA/comment-mock/domain/model"
	"github.com/QMDAKA/comment-mock/domain/repository"
	"github.com/QMDAKA/comment-mock/handler/common"
)

type LoginUser interface {
	CurrentUser(ctx context.Context) (*model.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (*model.User, error)
}

// Auth .
type Auth struct {
	// リポジトリ
	userRepo repository.User
}

// NewAuth .
func NewAuth(
	userRepo repository.User,
) *Auth {
	return &Auth{
		userRepo: userRepo,
	}
}

func (a *Auth) CurrentUser(ctx context.Context) (*model.User, error) {
	user, ok := common.GetCurrentUser(ctx).(*model.User)
	if !ok {
		return user, apperr.New_(apperr.Unauthorized, apperr.OptCltMsg("unauthorized"))
	}
	return user, nil
}

func (a *Auth) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	user, err := a.userRepo.GetByUUID(ctx, uuid)
	if err != nil {
		return user, apperr.New_(apperr.Unauthorized, apperr.OptCltMsg(err.Error()))
	}
	return user, nil
}
