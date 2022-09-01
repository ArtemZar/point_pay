package rest

import (
	pb "bank/internal/api/gen/proto"
	"bank/internal/api/grpc"
	"bank/internal/model"
	"bank/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *RESTServer) handleTest(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, _ := g.Client.Test(context.Background(), &pb.Request{
			X: 3,
			Y: 5,
		})
		//c.IndentedJSON(http.StatusOK, resp.Z)
		c.JSON(http.StatusOK, gin.H{
			"Response": resp.Z,
		})
	}
}

func (s *RESTServer) handleCreateAccount(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.CreateAccountDTO{}
		c.Bind(&requestBody)
		utils.Logger.Info(requestBody.Email)
		resp, _ := g.Client.CreateAccount(context.Background(), &pb.NewUserRequest{
			Email: requestBody.Email,
		})

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": resp.Id,
		})
	}
}

func (s *RESTServer) handleGetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (s *RESTServer) handleGenerateAddress(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody)
		resp, _ := g.Client.GenerateAddress(context.Background(), &pb.NewWalletRequest{
			Id: requestBody.ID,
		})

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Wallet_ID":  resp.WalletId,
		})
	}
}

func (s *RESTServer) handleDeposit(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody)

		resp, _ := g.Client.Deposit(context.Background(), &pb.ChangeBalanceRequest{
			Id:     requestBody.ID,
			Change: requestBody.Amount,
		})

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Balance":    resp.Balance,
		})
	}
}

func (s *RESTServer) handleWithdrawl(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody)

		resp, _ := g.Client.Withdrawal(context.Background(), &pb.ChangeBalanceRequest{
			Id:     requestBody.ID,
			Change: requestBody.Amount,
		})

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Balance":    resp.Balance,
		})
	}
}
