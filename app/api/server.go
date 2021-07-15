package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/QMDAKA/comment-mock/handler"
)

type Server struct {
	// アプリケーションが使用するDBのクライアント
	db     *gorm.DB
	router *gin.Engine
	// handler集
	handlers    HandlerCollection
	handlersMap map[string]handler.APIHandler
}

// NewServer .
func NewServer(db *gorm.DB, handlers HandlerCollection) Server {
	router := gin.New()
	s := Server{
		db:       db,
		router:   router,
		handlersMap: make(map[string]handler.APIHandler),
		handlers: handlers,
	}
	return s
}

// Close はサーバの終了時処理
func (s *Server) Close() error {
	// DBコネクションのクローズ処理
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Start() {
	s.RegisterRouter()
	if err := s.router.Run(); err != nil {
		fmt.Printf("server run error: %s", err)
	}
}

func (s *Server) Handler() {
	for _, h := range s.handlers {
		s.Use(h)
	}
}

func (s *Server) Use(handler handler.APIHandler) {
	s.handlersMap[handler.GetKey()] = handler
}

func (s *Server) RegisterRouter() *gin.Engine {

	for _, handler := range s.handlersMap {
		apiRouter := s.router.Group("")
		handler.API(apiRouter)
	}
	return s.router
}
