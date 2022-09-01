package rest

import (
	"bank/internal/api/grpc"
	"bank/internal/config"
	"bank/internal/utils"
	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	config *config.ServerConfig
	router *gin.Engine
}

func NewServer(config *config.ServerConfig) *RESTServer {
	return &RESTServer{
		config: config,
		router: gin.Default(),
	}
}

func (s *RESTServer) Start() error {

	s.configureRouter()
	utils.Logger.Info("The RESTServer is starting")

	return s.router.Run(s.config.BindAddr)
}

func (s *RESTServer) configureRouter() {
	client := grpc.NewGRPCClient()
	s.router.GET("/test", s.handleTest(client))
	s.router.POST("/create_account", s.handleCreateAccount(client))
	s.router.GET("/get_accounts", s.handleGetAccounts())
	s.router.POST("/generate_address", s.handleGenerateAddress(client))
	s.router.POST("/deposit", s.handleDeposit(client))
	s.router.POST("/withdrawl", s.handleWithdrawl(client))

}
