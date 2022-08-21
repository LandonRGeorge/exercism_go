package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	r := InterestRate(balance)
	return balance * (float64(r) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	i := Interest(balance)
	return balance + i
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	b := balance
	n := 0
	for {
		b = AnnualBalanceUpdate(b)
		if b >= targetBalance {
			if n > 0 {
				return n + 1
			}
			return n
		} else {
			n++
		}
	}
}
