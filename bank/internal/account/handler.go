package account

import (
	"bank/internal/api/grpc"
	pb "bank/internal/api/proto"
	"bank/internal/ifaces"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net/http"
)

type handler struct {
	GRPCClient *grpc.GRPCClient
}

func NewHandler() ifaces.Handler {
	return &handler{
		GRPCClient: grpc.NewGRPCClient(),
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/create_account", h.CreateAccount(h.GRPCClient))
	router.GET("/get_accounts", h.GetAccounts(h.GRPCClient))
	router.PATCH("/generate_address", h.GenerateAddress(h.GRPCClient))
	router.PATCH("/deposit", h.Deposit(h.GRPCClient))
	router.PATCH("/withdrawl", h.Withdrawl(h.GRPCClient))
}

func (h *handler) CreateAccount(g *grpc.GRPCClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := CreateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok
		resp, _ := g.Client.CreateAccount(context.Background(), &pb.NewUserRequest{
			Email: requestBody.Email,
		})

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
		resp, _ := g.Client.GenerateAddress(context.Background(), &pb.NewWalletRequest{
			Id: requestBody.ID,
		})

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