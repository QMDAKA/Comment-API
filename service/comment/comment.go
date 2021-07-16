package comment

import (
	"context"

	"github.com/QMDAKA/comment-mock/auth"
	"github.com/QMDAKA/comment-mock/domain/model"
	"github.com/QMDAKA/comment-mock/domain/repository"
)

type CommentServicer interface {
	GetAll()
	Create(ctx context.Context, comment *model.Comment) error
	Update(ctx context.Context, comment *model.Comment) error
}

type Comment struct {
	commentRepo repository.Comment
	tx          repository.Tx
	loginUser   auth.LoginUser
}

func NewComment(
	commentRepo repository.Comment,
	tx repository.Tx,
	loginUser auth.LoginUser,
) *Comment {
	return &Comment{
		commentRepo: commentRepo,
		tx:          tx,
		loginUser:   loginUser,
	}
}

func (c *Comment) GetAll() {

}

func (c *Comment) Create(ctx context.Context, comment *model.Comment) error {
	user, err := c.loginUser.CurrentUser(ctx)
	if err != nil {
		return err
	}
	comment.UserID = user.ID
	return c.tx.Transaction(ctx, func(ctx context.Context) error {
		// TODO: 1個レコードだけ作成するなら、Tx不要
		return c.commentRepo.Create(ctx, comment)
	})
}

func (c *Comment) Update(ctx context.Context, comment *model.Comment) error {
	return nil
}
