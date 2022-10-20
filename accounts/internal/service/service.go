package service

import (
	"accounts/internal/model"
	"accounts/internal/utils"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type accountRepo interface {
	Create(ctx context.Context, account model.Account) (accID string, err error)
	Update(ctx context.Context, account model.Account) error
	Find(ctx context.Context) (acc []model.Account, err error)
	GetOne(ctx context.Context, account model.Account) (acc model.Account, err error)
}

type Service struct {
	Storege accountRepo
}

func New(repo accountRepo) *Service {
	return &Service{Storege: repo}
}

func (s *Service) CreateAccount(ctx context.Context, email string) (accID string, err error) {
	fmt.Printf(email)
	return s.Storege.Create(ctx, model.Account{
		ID:      "",
		Email:   email,
		Balance: "0",
	})
}

func (s *Service) GenerateAddress(ctx context.Context, accID string) (walletID uint64, err error) {
	uuid := uuid.New()
	walletAddress := binary.BigEndian.Uint64(uuid[:16])

	updateWallet := model.Account{
		ID:       accID,
		WalletID: walletAddress,
	}

	err = s.Storege.Update(ctx, updateWallet)
	if err != nil {
		//TODO return error
		utils.Logger.Info("update error ", err)
	}

	return walletAddress, nil
}

func (s *Service) Deposit(ctx context.Context, accID, amountOfChange string) (newBalance string, err error) {
	sourceAcc, _ := s.Storege.GetOne(ctx, model.Account{ID: accID})

	newBalance = ChangeBalance(sourceAcc.Balance, amountOfChange, "deposit")

	updateBalance := model.Account{
		ID:      accID,
		Balance: newBalance,
	}

	err = s.Storege.Update(ctx, updateBalance)
	if err != nil {
		//TODO return error
		utils.Logger.Info("update error ", err)
	}

	return newBalance, nil
}

func (s *Service) Withdrawal(ctx context.Context, accID, amountOfChange string) (newBalance string, err error) {
	sourceAcc, _ := s.Storege.GetOne(ctx, model.Account{ID: accID})

	a, _ := strconv.ParseFloat(sourceAcc.Balance, 32)
	b, _ := strconv.ParseFloat(amountOfChange, 32)

	if a < b {
		return sourceAcc.Balance, fmt.Errorf("not enough balance")
	}

	newBalance = ChangeBalance(sourceAcc.Balance, amountOfChange, "withdrawal")

	updateBalance := model.Account{
		ID:      accID,
		Balance: newBalance,
	}

	err = s.Storege.Update(ctx, updateBalance)
	if err != nil {

		utils.Logger.Info("update error ", err)
	}

	return newBalance, nil
}

func (s *Service) GetAccounts() (acc []model.Account, err error) {
	return s.Storege.Find(context.Background())
}

func ChangeBalance(oldBalance, amountOfChange string, operation string) string {
	obToFloat, _ := strconv.ParseFloat(oldBalance, 32)
	chToFloat, _ := strconv.ParseFloat(amountOfChange, 32)
	var res float64
	switch operation {
	case "withdrawal":
		res = obToFloat - chToFloat

	case "deposit":
		res = obToFloat + chToFloat

	}

	return fmt.Sprint(res)
}
