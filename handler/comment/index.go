package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/service/comment"
)

type Index struct {
	Path           string
	Method         string
	commentService comment.CommentServicer
}

func NewCommentIndex(commentService comment.CommentServicer) *Index {
	return &Index{
		Path:           "/comments",
		Method:         http.MethodGet,
		commentService: commentService,
	}
}

func (i *Index) API(router *gin.RouterGroup) {
	router.Handle(i.Method, i.Path, func(c *gin.Context) {

	})
}

func (i *Index) GetKey() string {
	return i.Method + i.Path
}
