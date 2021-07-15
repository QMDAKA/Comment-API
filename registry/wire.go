// +build wireinject

package registry

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/app/api"
)

func InitializeServer(db *gorm.DB) (api.Server, error) {
	wire.Build(
		ServiceSet,
		HandlerSet,
		api.NewServer,
	)
	return api.Server{}, nil
}
