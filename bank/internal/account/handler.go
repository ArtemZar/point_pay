package account

import (
	"bank/internal/config"
	"bank/internal/ifaces"
	"bank/internal/transport/grpc"
	pb "bank/internal/transport/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net/http"
)

type handler struct {
	GRPCClient *grpc.GRPCClient
}

func NewHandler(cfg *config.Config) ifaces.Handler {
	return &handler{
		GRPCClient: grpc.NewGRPCClient(cfg),
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/create_account", h.CreateAccount(h.GRPCClient))
	router.GET("/get_accounts", h.GetAccounts(h.GRPCClient))
	router.PATCH("/generate_address", h.GenerateAddress(h.GRPCClient))
	router.PATCH("/deposit", h.Deposit(h.GRPCClient))
	router.PATCH("/withdrawal", h.Withdrawl(h.GRPCClient))
}

func (h *handler) CreateAccount(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := CreateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok
		resp, err := g.Client.CreateAccount(context.Background(), &pb.NewUserRequest{
			Email: requestBody.Email,
		})
		if err != nil {
			fmt.Errorf("request CreateAccount error, %v", err)
			c.Status(501)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Account_ID": resp.Id,
		})
	}
}

func (h *handler) GetAccounts(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, _ := g.Client.GetAccounts(context.Background(), &emptypb.Empty{})

		for {
			res, err := resp.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Something happened: %v", err)
			}

			c.JSON(http.StatusOK, gin.H{
				"Account_ID": res.Id,
				"Wallet_ID":  res.WalletId,
				"Balance":    res.Balance,
			})

		}

	}
}

func (h *handler) GenerateAddress(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok
		resp, err := g.Client.GenerateAddress(context.Background(), &pb.NewWalletRequest{
			Id: requestBody.ID,
		})
		if err != nil {
			c.JSON(501, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Wallet_ID":  resp.WalletId,
		})
	}
}

func (h *handler) Deposit(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok

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

func (h *handler) Withdrawl(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok

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
