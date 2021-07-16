package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/QMDAKA/comment-mock/auth"
	"github.com/QMDAKA/comment-mock/common/apperr"
	"github.com/QMDAKA/comment-mock/handler/common"
)

const (
	Authorization = "Authorization"
)

// Auth .
type Auth struct {
	loginUser auth.LoginUser
}

// ProvideAuth .
func ProvideAuth(loginUser auth.LoginUser) Auth {
	return Auth{loginUser: loginUser}
}

func (a *Auth) UserAuth(c *gin.Context) {
	// TODO: use jwt to extract uuid
	uuid := c.GetHeader(Authorization)
	if uuid == "" {
		common.SetErrorResponse(c, apperr.Wrap(apperr.BadRequest, errors.New("bad header")))
		c.Abort()
		return
	}
	user, err := a.loginUser.GetUserByUUID(c.Request.Context(), uuid)
	if err != nil {
		common.SetErrorResponse(c, err)
	}
	c.Set(common.LoggedInUser, user)
	c.Next()
}
