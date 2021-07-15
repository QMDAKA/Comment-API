// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/QMDAKA/comment-mock/service/comment"
)

var ServiceSet = wire.NewSet(
	comment.NewComment,
	wire.Bind(new(comment.CommentServicer), new(*comment.Comment)),
)
