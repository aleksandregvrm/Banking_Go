package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	const inflation = 6.6
	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("For how many years?")
	fmt.Scan(&years)
	fmt.Print("Expected return rate:")
	fmt.Scan(&expectedReturnRate)
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflation/100, years)
	fmt.Printf("Your real value after %.0f year would be: %.2f\n", years, futureRealValue)
}
