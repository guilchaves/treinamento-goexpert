// INFO: 16-Ponteiros e structs
package main

import "fmt"

type Account struct {
	balance float64
}

func (a Account) SimulateDeposit(value float64) float64 {
	a.balance += value
	fmt.Println(a.balance)
	return a.balance
}

func (a *Account) Deposit(value float64) {
	a.balance += value
}

func (a *Account) Withdraw(value float64) {
	a.balance -= value
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func main() {
	account := Account{balance: 100}
	account.SimulateDeposit(200)
	fmt.Println(account.balance)

	account.Deposit(300)
	fmt.Println(account.balance)
}
