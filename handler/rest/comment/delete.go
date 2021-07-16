package comment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/handler/common"
	"github.com/QMDAKA/comment-mock/service/comment"
)

type Delete struct {
	Path           string
	Method         string
	commentService comment.CommentServicer
}

func NewCommentDelete(
	commentService comment.CommentServicer,
) *Delete {
	return &Delete{
		Path:           "/comments/:id",
		Method:         http.MethodDelete,
		commentService: commentService,
	}
}

func (d *Delete) API(router *gin.RouterGroup) {
	router.Handle(d.Method, d.Path, common.NewHandler(func(ginCtx *gin.Context, ctx context.Context) {
		id, err := strconv.ParseUint(ginCtx.Param("id"), 10, 64)
		if err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		if err := d.commentService.Delete(ctx, id); err != nil {
			common.SetErrorResponse(ginCtx, err)
		}
		ginCtx.JSON(http.StatusOK, common.ID{ID: id})
	}))
}

func (d *Delete) GetKey() string {
	return d.Method + d.Path
}

func (d *Delete) LoginRequire() bool {
	return true
}
