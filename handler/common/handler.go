package common

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/common/apperr"
)

func NewHandler(f func(gCtx *gin.Context, ctx context.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// コンテキストの生成
		ctx := createContext(c)
		// ctx := context.Background()
		f(c, ctx)
	}
}

type ID struct {
	ID uint64 `json:"id"`
}

type Error struct {
	Errors []ErrorDetail `json:"error"`
}

// ErrorDetail エラー詳細
type ErrorDetail struct {
	Message string `json:"message"`
}

func SetErrorResponse(ginCtx *gin.Context, err error) {
	if aErr, ok := apperr.AsAppError(err); ok {
		ginCtx.JSON(apperr.ToHTTPStatus(err), toRespFromAppErr(aErr))
		return
	}
	ginCtx.JSON(http.StatusInternalServerError, Error{
		Errors: []ErrorDetail{
			{
				Message: "internal error",
			},
		},
	})

}

func toRespFromAppErr(aErr *apperr.AppError) Error {
	return Error{
		Errors: []ErrorDetail{
			{
				Message: aErr.ClientMessage(),
			},
		},
	}
}
