package handler

import (
	"github.com/gin-gonic/gin"
)

// APIHandler .
type APIHandler interface {
	API(router *gin.RouterGroup)
	GetKey() string
	LoginRequire() bool
}
