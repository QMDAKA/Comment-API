package api

import (
	"github.com/QMDAKA/comment-mock/handler"
	"github.com/QMDAKA/comment-mock/handler/comment"
)

// HandlerCollection Handler集
type HandlerCollection []handler.APIHandler

func NewHandlerCollection(
	commentIndex *comment.Index,
) HandlerCollection {
	return []handler.APIHandler{
		commentIndex,
	}
}
