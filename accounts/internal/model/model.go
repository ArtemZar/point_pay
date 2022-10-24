package model

type MyNumber uint64

type Account struct {
	ID       string   `bson:"_id,omitempty"`
	Email    string   `bson:"email,omitempty"`
	WalletID MyNumber `bson:"wallet_id,omitempty"`
	Balance  string   `bson:"balance,omitempty"`
}
