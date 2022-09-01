package service

import (
	"fmt"
	"strconv"
)

type BalanceChanger interface {
	Operation(oldBalance, amountOfChange float64) float64
}

type Deposite struct{}

func (d Deposite) Operation(oldBalance, amountOfChange float64) float64 {
	return oldBalance + amountOfChange
}

type Withdrawal struct{}

func (w Withdrawal) Operation(oldBalance, amountOfChange float64) float64 {
	return oldBalance - amountOfChange
}

func ChangeBalance(oldBalance, amountOfChange string, operation BalanceChanger) (string, error) {

	obToFloat, _ := strconv.ParseFloat(oldBalance, 32)
	chToFloat, _ := strconv.ParseFloat(amountOfChange, 32)
	res := operation.Operation(obToFloat, chToFloat)
	return fmt.Sprint(res), nil
}
