package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/handler/common"
	"github.com/QMDAKA/comment-mock/service/comment"
)

type Create struct {
	Path           string
	Method         string
	commentService comment.CommentServicer
}

func NewCommentCreate(commentService comment.CommentServicer) *Create {
	return &Create{
		Path:           "/post/:id/comments",
		Method:         http.MethodPost,
		commentService: commentService,
	}
}

func (c *Create) API(router *gin.RouterGroup) {
	router.Handle(c.Method, c.Path, func(ginCtx *gin.Context) {
		id, err := strconv.ParseUint(ginCtx.Param("id"), 10, 64)
		if err != nil {
			ginCtx.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		var in CreateCommentIn
		if err := ginCtx.Bind(&in); err != nil {
			ginCtx.JSON(400, gin.H{
				"error": err,
			})
		}
		commentIn := in.convert(id)
		if err := c.commentService.Create(ginCtx, commentIn); err != nil {
			ginCtx.JSON(400, gin.H{
				"error": err,
			})
		}
		ginCtx.JSON(http.StatusOK, common.ID{ID: commentIn.ID})
	})
}

func (c *Create) GetKey() string {
	return c.Method + c.Path
}
