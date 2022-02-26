package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64

func (btc Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", btc)
}

type Wallet struct {
	bal Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.bal += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.bal
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.bal {
		return ErrInsufficientFunds
	}
	w.bal -= amount
	return nil
}
