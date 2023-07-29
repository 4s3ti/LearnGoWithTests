package main

import "fmt"
import "errors"


var ErrInssufficientFunds = errors.New("cannot withdraw. Insufficient funds")

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Printf("address of balance in Depoisit is %v \n", &w.balance)
}

func (w *Wallet) Balance()  Bitcoin {
	// This is valid	
	return w.balance
	// this is also valid
	// return (*w).balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInssufficientFunds
	}

	w.balance -= amount
	return nil
}
