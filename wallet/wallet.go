package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(d int) {
  fmt.Printf("address in Deposit call %v\n", &w.balance)
	w.balance += d
}

func (w *Wallet) Balance() int {
	return w.balance
}
