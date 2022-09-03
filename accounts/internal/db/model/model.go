package model

type Account struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	WalletID uint64 `json:"wallet_id" bson:"wallet_id,omitempty"`
	Balance  string `json:"balance" bson:"balance,omitempty"`
}
