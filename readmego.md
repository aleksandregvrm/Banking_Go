package main

import (
	"fmt"
)

func main() {
	displayMyRevenue := displayRevenue()
	fmt.Printf("Revenue to display is %.f\n", displayMyRevenue)
}

func displayRevenue() float64 {
	var revenueToCalculate float64
	fmt.Print("Revenue:")
	fmt.Scan(&revenueToCalculate)
	return revenueToCalculate
}

// func calculatePow(val float64, powerOf float64) (fv float64, rfv float64) {
// 	fv = val
// 	rfv = powerOf
// 	return
// }
