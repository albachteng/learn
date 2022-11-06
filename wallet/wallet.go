package wallet

import "fmt"


type Bitcoin int

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(d Bitcoin) {
	w.balance += d
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
