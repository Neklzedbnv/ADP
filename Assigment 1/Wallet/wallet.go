package Wallet

type Wallet struct {
	Balance      float64
	Transactions []float64
}

func NewWallet() Wallet {
	return Wallet{Balance: 0}
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, amount)
}

func (w *Wallet) SpendMoney(amount float64) {
	w.Balance -= amount
	w.Transactions = append(w.Transactions, -amount)
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}
