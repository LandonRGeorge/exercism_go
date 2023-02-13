package account

// Define the Account type here.
type Account struct {
	amount int64
	open   bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount, open: true}
}

func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.open {
		return a.amount, false
	}
	if amount < 0 && a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	amount := a.amount
	a.open = false
	a.amount = 0
	return amount, true

}
