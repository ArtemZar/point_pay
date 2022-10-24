package rest

import (
	"bank/internal/config"
	"bank/internal/service"
	"bank/internal/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewHTTPServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: gin.New(),
	}
}

func (s *Server) Start() error {
	srv := service.NewService(s.config)
	handler := NewHandlers(srv)
	handler.Register(s.router)

	utils.Logger.Info("The Server is starting")

	return s.router.Run(s.config.SrvConf.BindAddr)
}
