package api

import (
	"github.com/QMDAKA/comment-mock/handler"
	"github.com/QMDAKA/comment-mock/handler/rest/comment"
)

// HandlerCollection Handler集
type HandlerCollection []handler.APIHandler

func NewHandlerCollection(
	commentIndex *comment.Index,
	commentCreate *comment.Create,
	commentDelete *comment.Delete,
	commentUpdate *comment.Update,
) HandlerCollection {
	return []handler.APIHandler{
		commentIndex,
		commentCreate,
		commentDelete,
		commentUpdate,
	}
}
