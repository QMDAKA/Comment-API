package common

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	LoggedInUser = "logged_in_user"
)

// createContext context生成
func createContext(ginCtx *gin.Context) context.Context {
	ctx := context.Background()
	fmt.Println(ginCtx.Get(LoggedInUser))
	if user, ok := ginCtx.Get(LoggedInUser); ok {
		ctx = WithCurrentUser(ctx, user)
	}
	return ctx
}

// WithCurrentUser はContextにログイン中のユーザを埋め込む
func WithCurrentUser(ctx context.Context, user interface{}) context.Context {
	return context.WithValue(ctx, LoggedInUser, user)
}

// GetCurrentUser はContextからログイン中ユーザを取り出す
func GetCurrentUser(ctx context.Context) interface{} {
	return ctx.Value(LoggedInUser)
}
