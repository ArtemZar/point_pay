package storage

import (
	"accounts/internal/db/model"
	"context"
)

type Storage interface {
	Create(ctx context.Context, account model.Account) (string, error)
	Update(ctx context.Context, account model.Account) error
	Find(ctx context.Context) ([]model.Account, error)
	GetOne(ctx context.Context, account model.Account) (model.Account, error)
}
