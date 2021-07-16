package comment

import (
	"github.com/QMDAKA/comment-mock/domain/model"
)

type CreateCommentIn struct {
	Content string `json:"content" binding:"required,min=10,max=100"`
}

type UpdateCommentIn struct {
	Content string `json:"content" binding:"required,min=10,max=100"`
}

func (c *CreateCommentIn) convert(postID uint64) *model.Comment {
	return &model.Comment{
		PostID:  postID,
		Content: c.Content,
	}
}

func (c *UpdateCommentIn) convert(postID uint64) *model.Comment {
	return &model.Comment{
		PostID:  postID,
		Content: c.Content,
	}
}