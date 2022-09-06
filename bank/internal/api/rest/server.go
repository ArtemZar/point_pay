package rest

import (
	"bank/internal/account"
	"bank/internal/config"
	"bank/internal/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: gin.New(),
	}
}

func (s *Server) Start() error {
	handler := account.NewHandler(s.config)
	handler.Register(s.router)

	utils.Logger.Info("The Server is starting")

	return s.router.Run(s.config.SrvConf.BindAddr)
}
