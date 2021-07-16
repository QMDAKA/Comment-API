// +build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/QMDAKA/comment-mock/app/api"
	"github.com/QMDAKA/comment-mock/handler/rest/comment"
)

var HandlerSet = wire.NewSet(
	api.NewHandlerCollection,
	// core
	comment.NewCommentIndex,
	comment.NewCommentCreate,
	comment.NewCommentDelete,
	comment.NewCommentUpdate,
)
