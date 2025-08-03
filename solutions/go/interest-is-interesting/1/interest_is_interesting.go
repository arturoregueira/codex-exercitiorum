package interest

import (
	"math"
)

var interest_rates = []float32{3.213000, 0.500000, 1.621000, 2.475000}
var amounts = []float64{0.00, 1000.00, 5000.00, 99999999999.00}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var outRate float32
	for i := 0; i < len(interest_rates); i++ {

		if balance < amounts[i] {
			outRate = interest_rates[i]

			//fmt.Println(outRate)
			break
		}

		//fmt.Println(outRate)

	}

	return outRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {

	rate := InterestRate(balance)

	return balance * float64(rate) / 100.00
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {

	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	interestRate := float64(InterestRate(balance)) / 100.00
	//fmt.Println(interestRate)
	//targetToEarn := targetBalance - balance

	output := math.Log(targetBalance/balance) / math.Log(1.00+interestRate)
	//fmt.Println(math.Log(targetBalance / balance))
	//fmt.Println(math.Log(1.00 + interestRate))

	return int(math.Ceil(output))
}
