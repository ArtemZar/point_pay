package rest

import (
	"bank/internal/model"
	pb "bank/internal/transport/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type useCaseService interface {
	CreateAccount(requestBody model.CreateAccountDTO) (*pb.AccountResponse, error)
	GenerateAddress(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error)
	Deposit(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error)
	Withdrawal(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error)
	GetAccounts() (pb.Accounts_GetAccountsClient, error)
}

type Handler struct {
	Service useCaseService
}

func NewHandlers(uc useCaseService) *Handler {
	return &Handler{
		Service: uc,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	router.POST("/create_account", h.CreateAccount())
	router.PATCH("/generate_address", h.GenerateAddress())
	router.PATCH("/deposit", h.Deposit())
	router.PATCH("/withdrawal", h.Withdrawal())
	router.GET("/get_accounts", h.GetAccounts())
}

func (h Handler) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.CreateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok
		resp, err := h.Service.CreateAccount(requestBody)
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

func (h Handler) GenerateAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok
		resp, err := h.Service.GenerateAddress(requestBody)
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

func (h Handler) Deposit() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok

		resp, _ := h.Service.Deposit(requestBody)

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Balance":    resp.Balance,
		})
	}
}

func (h Handler) Withdrawal() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.UpdateAccountDTO{}
		c.Bind(&requestBody) //nolint:errcheck // ok

		resp, _ := h.Service.Withdrawal(requestBody)

		c.JSON(http.StatusOK, gin.H{
			"Account_ID": requestBody.ID,
			"Balance":    resp.Balance,
		})
	}
}

func (h Handler) GetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, _ := h.Service.GetAccounts()

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
