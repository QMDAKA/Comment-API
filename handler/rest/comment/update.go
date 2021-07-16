package comment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/handler/common"
	"github.com/QMDAKA/comment-mock/service/comment"
)

type Update struct {
	Path           string
	Method         string
	commentService comment.CommentServicer
}

func NewCommentUpdate(
	commentService comment.CommentServicer,
) *Update {
	return &Update{
		Path:           "/post/:id/comments",
		Method:         http.MethodPatch,
		commentService: commentService,
	}
}

func (u *Update) API(router *gin.RouterGroup) {
	router.Handle(u.Method, u.Path, common.NewHandler(func(ginCtx *gin.Context, ctx context.Context) {
		id, err := strconv.ParseUint(ginCtx.Param("id"), 10, 64)
		if err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		var in UpdateCommentIn
		if err := ginCtx.Bind(&in); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		commentIn := in.convert(id)
		if err := u.commentService.Create(ctx, commentIn); err != nil {
			common.SetErrorResponse(ginCtx, err)
		}
		ginCtx.JSON(http.StatusOK, common.ID{ID: commentIn.ID})
	}))
}

func (u *Update) GetKey() string {
	return u.Method + u.Path
}

func (u *Update) LoginRequire() bool {
	return true
}
