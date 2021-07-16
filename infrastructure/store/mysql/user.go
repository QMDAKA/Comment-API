package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/common/apperr"
	"github.com/QMDAKA/comment-mock/domain/model"
)

type User struct {
	db *gorm.DB
}

func ProvideUserRepo(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) GetByUUID(ctx context.Context, uuid string) (*model.User, error) {
	var user model.User
	if err := u.db.WithContext(ctx).Model(model.User{}).Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return nil, apperr.New_(apperr.Database, apperr.OptCltMsg(err.Error()))
	}
	return &user, nil
}
