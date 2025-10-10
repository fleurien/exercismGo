package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return float32(3.213)
	case balance >= 0 && balance < 1000:
		return float32(0.5)
	case balance >= 1000 && balance < 5000:
		return float32(1.621)
	default:
		return float32(2.475)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	switch {
	case balance < 0:
		return balance * 3.213 / 100
	case balance >= 0 && balance < 1000:
		return balance * 0.5 / 100
	case balance >= 1000 && balance < 5000:
		return balance * 1.621 / 100
	default:
		return balance * 2.475 / 100
	}
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	switch {
	case balance < 0:
		return (balance * 3.213 / 100) + balance
	case balance >= 0 && balance < 1000:
		return (balance * 0.5 / 100) + balance
	case balance >= 1000 && balance < 5000:
		return (balance * 1.621 / 100) + balance
	default:
		return (balance * 2.475 / 100) + balance
	}
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	i := 0
		for i = 0; balance < targetBalance; i++ {
			interest := Interest(balance)
			balance = balance + interest
		}
	return i
}
