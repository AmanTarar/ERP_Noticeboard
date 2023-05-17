package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
		
	}
}

func (server *Server) Run(addr string) error {
	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}
