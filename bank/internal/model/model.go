package model

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	WalletID uint64 `json:"wallet_id" bson:"wallet_id,omitempty"`
	Balance  string `json:"balance" bson:"balance"`
}

type CreateAccountDTO struct {
	Email string `json:"email"`
}

type UpdateAccountDTO struct {
	ID     string `json:"id"`
	Amount string `json:"amount,omitempty"`
}
