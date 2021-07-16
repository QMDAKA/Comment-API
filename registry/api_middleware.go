// +build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/QMDAKA/comment-mock/middleware"
)

var MiddlewareSet = wire.NewSet(
	middleware.ProvideAuth,
)
