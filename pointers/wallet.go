package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Zenny int

func (z Zenny) String() string {
	return fmt.Sprintf("%d Z", z)
}

type Wallet struct {
	balance Zenny
}

func (w *Wallet) Deposit(amount Zenny) {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Zenny) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Zenny {
	return w.balance
}
