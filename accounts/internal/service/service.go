package service

import (
	"accounts/internal/model"
	"accounts/internal/utils"
	"context"
	"encoding/binary"
	"github.com/google/uuid"
)

type accountRepo interface {
	Create(ctx context.Context, account model.Account) (accID string, err error)
	Update(ctx context.Context, account model.Account) error
	Find(ctx context.Context) (acc []model.Account, err error)
	GetOne(ctx context.Context, account model.Account) (acc model.Account, err error)
	UpdateWithTrx(ctx context.Context, accountID, amountOfChange, operationType string) (model.Account, error)
}

type Service struct {
	Storege accountRepo
}

func New(repo accountRepo) *Service {
	return &Service{Storege: repo}
}

func (s *Service) CreateAccount(ctx context.Context, email string) (accID string, err error) {
	return s.Storege.Create(ctx, model.Account{
		Email: email,
	})
}

func (s *Service) GenerateAddress(ctx context.Context, accID string) (walletID uint64, err error) {
	uuid := uuid.New()
	walletAddress := binary.BigEndian.Uint64(uuid[:16])

	updateWallet := model.Account{
		ID:       accID,
		WalletID: model.MyNumber(walletAddress),
		Balance:  "0",
	}

	err = s.Storege.Update(ctx, updateWallet)
	if err != nil {
		utils.Logger.Info("update error: ", err)
		return 0, err
	}

	return walletAddress, nil
}

func (s *Service) Deposit(ctx context.Context, accID, amountOfChange string) (string, error) {
	updatedAccount, err := s.Storege.UpdateWithTrx(ctx, accID, amountOfChange, "deposit")
	if err != nil {
		utils.Logger.Info("update error ", err)
		return "", err
	}
	return updatedAccount.Balance, nil
}

func (s *Service) Withdrawal(ctx context.Context, accID, amountOfChange string) (newBalance string, err error) {
	updatedAccount, err := s.Storege.UpdateWithTrx(ctx, accID, amountOfChange, "withdrawal")
	if err != nil {
		utils.Logger.Info("update error ", err)
		return "", err
	}
	return updatedAccount.Balance, nil
}

func (s *Service) GetAccounts() (acc []model.Account, err error) {
	return s.Storege.Find(context.Background())
}
