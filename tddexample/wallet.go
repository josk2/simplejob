package tddexample

import "errors"

var ErrNotEnoughFund = errors.New("You not enough fund")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrNotEnoughFund
	}

	w.balance -= amount
	return nil
}
