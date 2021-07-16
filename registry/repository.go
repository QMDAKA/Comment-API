// +build wireinject

package registry

import (
	"github.com/google/wire"

	"github.com/QMDAKA/comment-mock/domain/repository"
	"github.com/QMDAKA/comment-mock/infrastructure/store/mysql"
)

var RepositorySet = wire.NewSet(
	mysql.ProvideCommentRepo,
	wire.Bind(new(repository.Comment), new(*mysql.Comment)),
	mysql.ProvidePostRepo,
	wire.Bind(new(repository.Post), new(*mysql.Post)),
	mysql.ProvideUserRepo,
	wire.Bind(new(repository.User), new(*mysql.User)),
	mysql.ProvideTransaction,
	wire.Bind(new(repository.Tx), new(*mysql.Transaction)),
)
